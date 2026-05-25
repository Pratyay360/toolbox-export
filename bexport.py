#!/usr/bin/env python3

import os
import subprocess
import sys
from pathlib import Path


def check_args() -> str:
    args = sys.argv
    if len(args) == 1:
        error("No application name provided. Usage: export-bin <app_name>")
    return args[1]


def check_container() -> None:
    if not os.path.exists("/run/.containerenv"):
        error("Not running inside a toolbox container")


def error(msg: str) -> None:
    print(f"\033[31;1m[ERROR]\033[0m {msg}")
    sys.exit(1)


def info(msg: str) -> None:
    print(f"\033[32;1m[INFO]\033[0m {msg}")


def get_container_name() -> str:
    check_container()
    try:
        with open("/run/.containerenv") as f:
            for line in f.readlines():
                if line.startswith("name="):
                    return line.split("=")[1].strip().replace('"', "")
    except Exception as e:
        error(f"Failed to read container env: {e}")

    error("Could not find container name in /run/.containerenv")
    return ""


def bin_export() -> None:
    target_bin = check_args()
    container_name = get_container_name()

    # Define and ensure the export directory exists BEFORE writing files to it
    bin_dir = Path.home() / ".local" / "toolbox"
    dir_existed = bin_dir.exists()
    bin_dir.mkdir(parents=True, exist_ok=True)

    try:
        # Using 'which' to find the absolute path of the binary inside the container
        result = subprocess.run(
            ["which", target_bin], capture_output=True, text=True, check=True
        )
        container_bin_path = result.stdout.strip()

        if container_bin_path:
            bin_name = os.path.basename(container_bin_path)
            alias_path = bin_dir / bin_name
            wrapper_content = f'''#!/bin/sh\n\tif [ -f /run/.containerenv ]; then
                    exec flatpak-spawn --host toolbox run -c "{container_name}" {container_bin_path} "$@"
                elif [ -z "$container" ] ; then
                    exec /usr/bin/toolbox run -c {container_name} {bin_name} "$@"
                else
                    exec host-spawn toolbox run -c {container_name} {container_bin_path} "$@"
                fi
            '''
            with open(alias_path, "w") as f:
                f.write(wrapper_content)

            os.chmod(alias_path, 0o755)
            info(f"Exported binary wrapper for '{bin_name}' to: {alias_path}")
            if not dir_existed:
                path_line = f'\nexport PATH="{bin_dir}:$PATH"\n'
                print(
                    f"\nAdd this to the end of your shell rc file (.bashrc / .zshrc):\n{path_line}"
                )

    except subprocess.CalledProcessError:
        error(f"Command '{target_bin}' not found inside this container.")


if __name__ == "__main__":
    bin_export()
