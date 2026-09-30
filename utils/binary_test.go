package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestFindContainerBinary(t *testing.T) {
	root := t.TempDir()
	binDir := filepath.Join(root, ".local", "toolbox")
	containerDir := filepath.Join(root, "bin")
	aliasDir := filepath.Join(root, "aliases")
	siblingDir := binDir + "-other"
	for _, dir := range []string{binDir, containerDir, aliasDir, siblingDir} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}
	wrapper := filepath.Join(binDir, "app")
	binary := filepath.Join(containerDir, "app")
	sibling := filepath.Join(siblingDir, "app")
	for _, file := range []string{wrapper, binary, sibling} {
		if err := os.WriteFile(file, []byte("#!/bin/sh\n"), 0o755); err != nil {
			t.Fatal(err)
		}
	}
	alias := filepath.Join(aliasDir, "app")
	if err := os.Symlink(wrapper, alias); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		app  string
		path string
		want string
	}{
		{"skip wrapper first in PATH", "app", binDir + ":" + containerDir, binary},
		{"wrapper only", "app", binDir, ""},
		{"explicit wrapper", wrapper, containerDir, ""},
		{"skip symlink to wrapper", "app", aliasDir + ":" + containerDir, binary},
		{"explicit symlink to wrapper", alias, containerDir, ""},
		{"explicit container binary", binary, binDir, binary},
		{"allow similarly named directory", "app", siblingDir, sibling},
		{"missing command", "missing", containerDir, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PATH", tt.path)
			got, err := findContainerBinary(tt.app, binDir)
			if tt.want == "" {
				if !errors.Is(err, exec.ErrNotFound) {
					t.Fatalf("expected executable not found, got %q, %v", got, err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("got %q, want %q", got, tt.want)
			}
		})
	}
}
