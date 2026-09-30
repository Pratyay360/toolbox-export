# toolbox-export

Script for exporting applications from toolbox or any other containers.

add this to your rc(.bashrc .zshrc / .profile ) files ...

don't just run any Script available on the internet. read it first unless you are a risk taker.

```bash
toolbox-export init
```
this sets everything up for you ->

```bash
export PATH="$HOME/.local/bin:$HOME/.local/toolbox:$PATH"
```

## Build

```bash
go build -o toolbox-export .
install -D toolbox-export "$HOME/.local/bin/toolbox-export"
```

## Usage
Enter the container
```bash
toolbox enter <container>
```
Export both the desktop launcher and the binary wrapper

```bash
toolbox-export firefox
```

Or export them individually

```bash
toolbox-export export firefox
toolbox-export binary firefox
```

download binaries
```bash
curl -fsSL https://github.com/pratyay360/toolbox-export/blob/main/install.sh | sh
```


## don't know I am also testing 
```bash
curl -sf http://goblin.run/github.com/pratyay360/toolbox-export | PREFIX=~/.local/bin sh
```


just some small scripts ment for my personal use.
built around [container-toolbx](https://github.com/containers/toolbox)
