package utils

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

func ExportLauncher(app string) error {
	if err := checkContainer("Not running from toolbox"); err != nil {
		return err
	}
	if app == "" {
		return errors.New("No application name")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	desktopDir := filepath.Join(home, ".local", "share", "applications")
	if err := os.MkdirAll(desktopDir, 0o755); err != nil {
		return err
	}

	matched, err := getDesktopFiles(app)
	if err != nil {
		return err
	}

	var iconNames []string
	for _, file := range matched {
		lines, err := readLines(file)
		if err != nil {
			continue
		}
		text := ""
		for _, line := range lines {
			if strings.HasPrefix(line, "Icon=") {
				parts := strings.Split(line, "=")
				iconName := ""
				if len(parts) > 1 {
					iconName = strings.TrimSpace(parts[1])
				}
				if !slices.Contains(iconNames, iconName) {
					iconNames = append(iconNames, iconName)
				}
			}

			switch {
			case strings.HasPrefix(line, "Exec=") && strings.Contains(line, app):
				name, _ := GetContainerName()
				text += strings.ReplaceAll(line, "Exec=", fmt.Sprintf("Exec=/usr/bin/toolbox run -c %s ", name))
			case strings.HasPrefix(line, "TryExec="):
				value := ""
				if parts := strings.SplitN(line, "=", 2); len(parts) > 1 {
					value = strings.TrimSpace(parts[1])
				}
				text += fmt.Sprintf("TryExec=%s\n", path.Base(value))
			case strings.HasPrefix(line, "Name="):
				text += strings.ReplaceAll(line, "\n", " (toolbox)\n")
			default:
				text += line
			}
		}

		fileName := filepath.Join(desktopDir, path.Base(file))
		if err := os.WriteFile(fileName, []byte(text), 0o644); err != nil {
			return err
		}
		Info(fmt.Sprintf("Exported desktop file: %s", fileName))
	}

	iconDir := filepath.Join(home, ".local", "share", "icons")
	if err := os.MkdirAll(iconDir, 0o755); err != nil {
		return err
	}
	for _, iconName := range iconNames {
		if iconName == "" {
			continue
		}
		var files []string
		files = append(files, findIconFiles("/usr/share/icons", iconName)...)
		files = append(files, findIconFiles("/usr/share/pixmaps", iconName)...)
		for _, file := range files {
			if strings.HasSuffix(file, "png") || strings.HasSuffix(file, "svg") {
				fileName := filepath.Join(iconDir, path.Base(file))
				if err := copyFile(file, fileName); err != nil {
					return err
				}
				Info(fmt.Sprintf("Exported icon: %s", fileName))
			}
		}
	}

	return nil
}

func getDesktopFiles(app string) ([]string, error) {
	var files []string
	for _, pattern := range []string{
		"/usr/share/applications/*.desktop",
		"/usr/local/share/applications/*.desktop",
	} {
		matches, _ := filepath.Glob(pattern)
		files = append(files, matches...)
	}

	var applications []string
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			continue
		}
		if strings.Contains(string(data), "NotShowIn") {
			continue
		}
		applications = append(applications, file)
	}

	var matched []string
	for _, appFile := range applications {
		lines, err := readLines(appFile)
		if err != nil {
			continue
		}
		for _, line := range lines {
			if strings.HasPrefix(line, "Exec=") && strings.Contains(line, app) {
				if execLaunchesApp(line, app) {
					matched = append(matched, appFile)
					break
				}
			}
		}
	}

	if len(matched) == 0 {
		return nil, errors.New("No application found")
	}
	return matched, nil
}

func execLaunchesApp(line, app string) bool {
	value := strings.TrimPrefix(strings.TrimSpace(line), "Exec=")
	for _, word := range strings.Fields(value) {
		word = strings.Trim(word, `"'`)
		if word != "" && path.Base(word) == app {
			return true
		}
	}
	return false
}

func findIconFiles(root, name string) []string {
	var out []string
	_ = filepath.WalkDir(root, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasPrefix(d.Name(), name+".") {
			out = append(out, p)
		}
		return nil
	})
	return out
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
