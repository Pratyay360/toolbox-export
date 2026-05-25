#!/usr/bin/env bash
mkdir -p "$HOME/.local/toolbox"
mkdir -p "$HOME/.local/share/applications"

"$HOME/.config/toolbox/export" "$@" 2>/dev/null || true
"$HOME/.config/toolbox/bexport" "$@" 2>/dev/null || true
