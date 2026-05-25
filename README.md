# toolbox-export
Script for exporting applications from toolbox or any other containers.

add this to your rc(.bashrc .zshrc / .profile ) files ...

don't just run any Script available on the internet. read it first unless you are a risk taker.

export PATH="$HOME/.local/bin:$HOME/.local/toolbox:$PATH"

```bash
curl -fsSL https://github.com/Pratyay360/toolbox-export/raw/refs/heads/main/install.sh | bash
```
## Usage
Enter the container
```bash
toolbox enter <container>
```
export the binaries/ launcher
```bash
toolbox-export firefox
```

just some small scripts ment for my personal use.
built around [container-toolbx](https://github.com/containers/toolbox)
