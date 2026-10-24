package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"charm.land/lipgloss/v2"
)

type artifactKind string

const (
	artifactDesktop artifactKind = "desktop"
	artifactIcon    artifactKind = "icon"
	artifactBinary  artifactKind = "binary"
)

type artifact struct {
	kind artifactKind
	path string
}

var (
	badgeStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	appStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#9A75F0"))

	kindDesktopStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#00D7D7"))

	kindIconStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF79C6"))

	kindBinaryStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#50FA7B"))

	filePathStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA"))

	hintStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6C6C6C"))

	skipStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888")).
			Italic(true)

	cancelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFAF00")).
			Italic(true)
)

func promptArtifact(art artifact, reader *bufio.Reader) string {
	var kindTag string
	switch art.kind {
	case artifactDesktop:
		kindTag = kindDesktopStyle.Render("[desktop]")
	case artifactIcon:
		kindTag = kindIconStyle.Render("[icon]")
	case artifactBinary:
		kindTag = kindBinaryStyle.Render("[binary]")
	default:
		kindTag = lipgloss.NewStyle().Foreground(lipgloss.Color("#888888")).Render(fmt.Sprintf("[%s]", art.kind))
	}

	badge := badgeStyle.Render("?")
	path := filePathStyle.Render(art.path)
	hint := hintStyle.Render("[y/N/a/q]: ")

	fmt.Printf("%s Remove %s %s? %s", badge, kindTag, path, hint)

	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println()
		return "quit"
	}

	response = strings.TrimSpace(strings.ToLower(response))
	switch response {
	case "y", "yes":
		return "yes"
	case "a", "all":
		return "all"
	case "q", "quit":
		return "quit"
	default:
		return "no"
	}
}

func findArtifacts(home, app string) ([]artifact, error) {
	var artifacts []artifact
	seen := make(map[string]bool)

	add := func(kind artifactKind, p string) {
		if !seen[p] {
			seen[p] = true
			artifacts = append(artifacts, artifact{kind: kind, path: p})
		}
	}

	desktopDir := filepath.Join(home, ".local", "share", "applications")
	pattern := filepath.Join(desktopDir, "*.desktop")
	matches, _ := filepath.Glob(pattern)

	var iconNames []string
	for _, file := range matches {
		lines, err := readLines(file)
		if err != nil {
			continue
		}
		isToolboxExported := false
		hasExecMatch := false
		var fileIconName string
		for _, line := range lines {
			if strings.HasPrefix(line, "Exec=") {
				if strings.Contains(line, "toolbox run") {
					isToolboxExported = true
				}
				if execLaunchesApp(line, app) {
					hasExecMatch = true
				}
			}
			if strings.HasPrefix(line, "Icon=") {
				if parts := strings.SplitN(line, "=", 2); len(parts) > 1 {
					fileIconName = strings.TrimSpace(parts[1])
				}
			}
		}

		if isToolboxExported && hasExecMatch {
			add(artifactDesktop, file)
			if fileIconName != "" {
				iconNames = append(iconNames, fileIconName)
			}
		}
	}

	iconDir := filepath.Join(home, ".local", "share", "icons")
	if entries, err := os.ReadDir(iconDir); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue // Do not touch subdirectories (e.g. icon themes)
			}
			name := entry.Name()
			baseNoExt := strings.TrimSuffix(name, filepath.Ext(name))

			isMatch := (baseNoExt == app)
			if !isMatch {
				if slices.Contains(iconNames, baseNoExt) {
					isMatch = true
				}
			}

			if isMatch {
				add(artifactIcon, filepath.Join(iconDir, name))
			}
		}
	}

	binPath := filepath.Join(home, ".local", "toolbox", app)
	if data, err := os.ReadFile(binPath); err == nil {
		if strings.Contains(string(data), "toolbox") {
			add(artifactBinary, binPath)
		}
	}

	return artifacts, nil
}

func RemoveExported(app string) error {
	if app == "" {
		return errors.New("no application name provided; usage: toolbox-export --rm <application>")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("resolve home directory: %w", err)
	}

	artifacts, err := findArtifacts(home, app)
	if err != nil {
		return fmt.Errorf("find artifacts: %w", err)
	}

	if len(artifacts) == 0 {
		Warn(fmt.Sprintf("No exported artifacts found for '%s'", app))
		return nil
	}

	fmt.Printf("\nFound %d artifact(s) for %s:\n\n", len(artifacts), appStyle.Render(app))

	reader := bufio.NewReader(os.Stdin)
	var removeErrors []error
	removeAll := false

	for _, art := range artifacts {
		shouldRemove := false

		if removeAll {
			shouldRemove = true
		} else {
			action := promptArtifact(art, reader)
			switch action {
			case "yes":
				shouldRemove = true
			case "all":
				shouldRemove = true
				removeAll = true
			case "quit":
				fmt.Println(cancelStyle.Render("Aborted remaining removals."))
				return errors.Join(removeErrors...)
			case "no":
				fmt.Println(skipStyle.Render(fmt.Sprintf("Skipped %s", art.path)))
				continue
			}
		}

		if shouldRemove {
			if err := os.Remove(art.path); err != nil {
				removeErrors = append(removeErrors, fmt.Errorf("remove %s (%s): %w", art.kind, art.path, err))
			} else {
				Info(fmt.Sprintf("Removed %s: %s", art.kind, art.path))
			}
		}
	}

	return errors.Join(removeErrors...)
}
