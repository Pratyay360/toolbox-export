package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"path/filepath"
	"strings"
)

func ExportBinary(app string) error {
	if app == "" {
		return errors.New("No application name provided. Usage: export-bin <app_name>")
	}
	if err := checkContainer("Not running inside a toolbox container"); err != nil {
		return err
	}
	containerName, ok := GetContainerName()
	if !ok {
		return errors.New("Could not find container name in /run/.containerenv")
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	binDir := filepath.Join(home, ".local", "toolbox")
	_, statErr := os.Stat(binDir)
	dirExisted := statErr == nil
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return err
	}

	containerBinPath, err := findContainerBinary(app, binDir)
	if err != nil {
		return fmt.Errorf("command %q not found inside this container outside %s: %w", app, binDir, err)
	}

	binName := filepath.Base(containerBinPath)
	aliasPath := filepath.Join(binDir, binName)
	wrapper := fmt.Sprintf("#!/bin/sh\n\tif [ -f /run/.containerenv ]; then\n  exec flatpak-spawn --host toolbox run -c \"%s\" %s \"$@\"\n                elif [ -z \"$container\" ] ; then\n                    exec /usr/bin/toolbox run -c %s %s \"$@\"\n                else\n                    exec toolbox run -c %s %s \"$@\"\n                fi\n            ", containerName, containerBinPath, containerName, containerBinPath, containerName, containerBinPath)

	if err := os.WriteFile(aliasPath, []byte(wrapper), 0o755); err != nil {
		return err
	}
	if err := os.Chmod(aliasPath, 0o755); err != nil {
		return err
	}

	Info(fmt.Sprintf("Exported binary wrapper for '%s' to: %s", binName, aliasPath))
	if !dirExisted {
		Warn(fmt.Sprintf("Add %s to PATH by appending 'export PATH=\"%s:$PATH\"' to your shell rc file (.bashrc / .zshrc)", binDir, binDir))
	}

	return nil
}

func findContainerBinary(app, binDir string) (string, error) {
	excludedDir, err := filepath.Abs(binDir)
	if err != nil {
		return "", err
	}
	if resolved, err := filepath.EvalSymlinks(excludedDir); err == nil {
		excludedDir = resolved
	}

	candidates := []string{app}
	if !strings.ContainsRune(app, filepath.Separator) {
		candidates = nil
		for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
			candidates = append(candidates, filepath.Join(dir, app))
		}
	}
	for _, candidate := range candidates {
		absolute, err := filepath.Abs(candidate)
		if err != nil {
			continue
		}
		if _, err := exec.LookPath(absolute); err != nil {
			continue
		}
		resolved, err := filepath.EvalSymlinks(absolute)
		if err != nil {
			continue
		}
		excluded := false
		for _, binaryPath := range []string{absolute, resolved} {
			relative, err := filepath.Rel(excludedDir, binaryPath)
			if err != nil || (relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))) {
				excluded = true
				break
			}
		}
		if !excluded {
			return absolute, nil
		}
	}
	return "", exec.ErrNotFound
}
