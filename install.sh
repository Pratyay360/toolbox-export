#!/usr/bin/env bash

set -euo pipefail

os=$(uname -s)

if [[ "$os" != "Linux" ]]; then
    echo "Unsupported OS: $os"
    exit 1
fi


echo "Installing toolbox-export..."

curl -fSsL https://github.com/pratyay360/toolbox-export/releases/download/latest/toolbox-export-$(uname -m) -o $HOME/.local/bin/toolbox-export
chmod +x $HOME/.local/bin/toolbox-export


echo "Installation complete."
