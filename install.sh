#!/usr/bin/env bash

set -euo pipefail

os=$(uname -s)

if [[ "$os" != "Linux" ]]; then
    echo "Unsupported OS: $os"
    exit 1
fi

echo "Installing toolbox-export..."

mkdir -p "$HOME/.local/bin"

tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"' EXIT

curl -fSsL \
    "https://github.com/Pratyay360/toolbox-export/releases/latest/download/toolbox-export_Linux_x86_64.tar.gz" \
    -o "$tmpdir/toolbox-export.tar.gz"

tar -xzf "$tmpdir/toolbox-export.tar.gz" -C "$HOME/.local/bin"

chmod +x "$HOME/.local/bin/toolbox-export"

echo "Installation complete."
echo "Installed to: $HOME/.local/bin/toolbox-export"
