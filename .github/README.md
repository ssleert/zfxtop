<div align="center">
<img src="https://github.com/ssleert/zfxtop/blob/master/assets/images/logo.png" width="50%">


### *`fetch`  top written by `bubbletea` enjoyer* :dolphin:

<br>
</div>

<img src="https://github.com/ssleert/zfxtop/blob/master/assets/images/ui.png" width="50%" align="right">

# Description 📖
Historically I don't use bars like `polybar`/`waybar` and so I have nowhere to put information about **time**, **CPU** load, **RAM** usage, **disk** usage, and other cool stuff so I decided to write a little ***tui*** utility for this purpose. Also I tried to make it as nice and nice to use as i can.

At this point `zfxtop` is in a very early stage of development so feel free to open an **issue** with your problems and suggestions.

The code is written in `Go` because it is blue :cup_with_straw:.

> **Note** that this is only **compatible** with the `linux` kernel and `x86` cpus at this time.

<br>

# Installation ☁️
with `curl`
```fish
curl -sSL raw.githubusercontent.com/ssleert/zfxtop/master/install.sh | sh
```
with `wget`
```fish
wget -qO- raw.githubusercontent.com/ssleert/zfxtop/master/install.sh | sh
```
Or if on Arch you can install the aur package zfxtop

# Configuration ⚙️
I decided that using `toml` or `yaml` makes no sense for such small configuration files, but it **increases code size**, so I decided to use `ini`
```fish
~/.config/zfxtop/conf.ini
```

<details>
<summary>config file example</summary>

```ini
[tui]

# time between info update in millisecond
update = 100

# requires nerd font
icons = true

# can be rounded, sharp, double, ascii, dot
borders = rounded

# enable or disable colors
colors = true


# colors are set in the 256-color palette
[colors]
faint  = 238
mid    = 245

load0  = 27
load1  = 63
load2  = 99
load3  = 135
load4  = 171
load5  = 207

tempr0 = 49
tempr1 = 79
tempr2 = 109
tempr3 = 139
tempr4 = 169
tempr5 = 199

list0  = 109
list1  = 79
list2  = 169
```

</details>

# Building 📦
install `Go` before it
```fish
git clone https://github.com/ssleert/zfxtop.git
cd zfxtop/
./scripts/build.sh
```
If you have a `CPU` with `amd64` architecture you can try to **build** with **optimizations**
```fish
# v2, v3, v4 supported
GOAMD64=v3 ./scripts/build.sh
```
with `podman/docker`
```fish
podman build . -t zfxtop
podman run -it --rm --name zfxtop zfxtop
```
> btw you can set alias with podman and use it as regular command
```bash
alias zfxtop="podman run -it --rm --name zfxtop zfxtop"
```

# Contribute
Before contributing, please run `contribute.sh` script
```fish
./scripts/contribute.sh
```

<div align="center">
<hr>

### made with 🫀 by `sfome`

</div>












