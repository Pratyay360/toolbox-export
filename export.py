#!/bin/python3

import glob
import os
import shutil
import sys


def check_args() -> None:
    args = sys.argv
    if len(args) == 1:
        error("No application name")


def check_container() -> None:
    if not os.path.exists("/run/.containerenv"):
        error("Not running from toolbox")


def error(msg: str) -> None:
    print(f"\033[31;1m[ERROR]\033[0m {msg}")
    exit(1)


def info(msg: str) -> None:
    print(f"\033[32;1m[INFO]\033[0m {msg}")


def get_desktop_files() -> list[str]:
    files: list[str] = []
    files.extend(glob.glob("/usr/share/applications/*.desktop"))
    files.extend(glob.glob("/usr/local/share/applications/*.desktop"))

    applications: list[str] = []
    for file in files:
        with open(file) as f:
            text: str = f.read()
            if "NotShowIn" in text:
                continue
            applications.append(file)

    matched_files: list[str] = []
    for app in applications:
        with open(app) as f:
            lines: list[str] = f.readlines()
            for line in lines:
                if line.startswith("Exec=") and sys.argv[1] in line:
                    if (
                        sys.argv[1]
                        == line.strip().strip("Exec=").split(" ")[0].split("/")[-1]
                    ):
                        matched_files.append(app)
                        break
    if matched_files == []:
        error("No application found")
    else:
        return matched_files


def export() -> None:
    desktop_dir: str = os.path.expanduser("~") + "/.local/share/applications"
    os.makedirs(desktop_dir, exist_ok=True)
    name = ""
    icon_names: list[str] = []
    for file in get_desktop_files():
        text: str = ""
        with open(file) as f:
            for line in f.readlines():
                if line.startswith("Icon="):
                    icon_name = line.split("=")[1].strip()
                    if icon_name not in icon_names:
                        icon_names.append(icon_name)
                if line.startswith("Exec=") and sys.argv[1] in line:
                    with open("/run/.containerenv") as c:
                        for l in c.readlines():
                            if l.startswith("name"):
                                name = l.split("=")[1].strip().replace('"', "")
                                break
                    text += line.replace(
                        "Exec=", f"Exec=/usr/bin/toolbox run -c {name} "
                    )
                elif line.startswith("TryExec="):
                    value = line.split("=", 1)[1].strip()
                    binary = os.path.basename(value)
                    text += f"TryExec={binary}\n"
                elif line.startswith("Name="):
                    text += line.replace("\n", " (toolbox)\n")
                else:
                    text += line

        file_name: str = desktop_dir + "/" + file.split("/")[-1]
        with open(file_name, "w") as f:
            f.write(text)
            info(f"Exported desktop file: {file_name}")

    for icon_name in icon_names:
        if icon_name != "":
            files: list[str] = []
            files.extend(
                glob.glob(f"/usr/share/icons/**/{icon_name}.*", recursive=True)
            )
            files.extend(
                glob.glob(f"/usr/share/pixmaps/**/{icon_name}.*", recursive=True)
            )
            icon_dir: str = f"{os.path.expanduser('~')}/.local/share/icons"
            for file in files:
                if file.endswith("png") or file.endswith("svg"):
                    file_name = f"{icon_dir}/{file.split('/')[-1]}"
                    shutil.copy(file, file_name)
                    info(f"Exported icon: {file_name}")


if __name__ == "__main__":
    check_container()
    check_args()
    export()
