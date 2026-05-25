#!/bin/bash

cpu_arch=$(arch)
if [ ! -f "$HOME/.local/bin/host-spawn" ]; then
    curl -fsSL "https://github.com/1player/host-spawn/releases/download/v1.6.2/host-spawn-$cpu_arch" -o "$HOME/.local/bin/host-spawn"
    chmod +x "$HOME/.local/bin/host-spawn"
fi
git clone --depth 1 https://github.com/pratyay360/toolbox-export.git
mkdir -p "$HOME/.local/toolbox"
mkdir -p "$HOME/.config/toolbox"
mkdir -p "$HOME/.local/share/applications"

cd toolbox-export
install -D export.py "$HOME/.config/toolbox/export"
install -D bexport.py "$HOME/.config/toolbox/bexport"
install -D toolbox-export.sh "$HOME/.local/bin/toolbox-export"
echo Installed!
rm -rf $PWD
