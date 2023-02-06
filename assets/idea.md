# -v --version
```
        ____     __             ╭───────────────────╮
 ____  / __/  __/ /_____  ____  │ author:  ssleert  │
/_  / / /_|'|/_/ __/ __ \/ __ \ │ lang:    go       │
 / /_/ __/>  </ /_/ /_/ / /_/ / │ version: 0.0.1    │
/___/_/ /_/|_|\__/\____/ .___/  │ commit:  0KLfgc93 │
                      /_/       ╰───────────────────╯
```
# -h --help
```
  zfxtop - fetch top for gen Z with X
 ╭────────────────────────────────────╮
 │ arg          │ desc                │
 │──────────────┼─────────────────────│
 │ -c --config  │ set config location │
 │ -v --version │ get version info    │
 │ -h --help    │ get help info       │
 ╰────────────────────────────────────╯
```

# whithout flags
```
 ╭─  CPU ───────────────── 19:47:27 ──────────────── E5-2660 ─╮
 │ ╭─────────────╮                                             │
 │ │  load 99%  │                         ▁▂▃▄▅▆▇████████████ │
 │ │ 龍fr 3.3ghz │                 ▁▂▃▄▅▆▇████████████████████ │
 │ │  temp 63°C │         ▁▂▃▄▅▆▇████████████████████████████ │
 │ ╰─────────────╯ ▁▂▃▄▅▆▇████████████████████████████████████ │
 │         ▁▂▃▄▅▆▇████████████████████████████████████████████ │
 │ ▁▂▃▄▅▆▇████████████████████████████████████████████████████ │
 ╰─────────────────────────────────────────────────────────────╯
 ╭─  MEM ───────── 15.55 GiB ─╮ ╭─  SWAP ────────── 4.0 GiB ─╮
 │ Used               4.46 GiB │ │ Free               1.23 GiB │
 │                             │ │                             │
 │                 ▁▂▃▄▅▆▇████ │ │                 ▁▂▃▄▅▆▇████ │
 │         ▁▂▃▄▅▆▇████████████ │ │         ▁▂▃▄▅▆▇████████████ │
 │ ▁▂▃▄▅▆▇████████████████████ │ │ ▁▂▃▄▅▆▇████████████████████ │
 ├─────────────────────────────┤ ╰─────────────────────────────╯
 │ Availible          11.0 GiB │ ╭─  DISK ────────── 223 GiB ─╮
 │                             │ │  root  66%       140.9 GiB │
 │                 ▁▂▃▄▅▆▇████ │ │  /home 54%       114.4 GiB │
 │         ▁▂▃▄▅▆▇████████████ │ │  /usr  10%       10.2  GiB │
 │ ▁▂▃▄▅▆▇████████████████████ │ ╰─────────────────────────────╯
 ├─────────────────────────────┤ ╭─  BAT ────────────── 3:45 ─╮
 │ Free               4.10 GiB │ │ ■■■■■■■■■■■■■■■■■■■■■■■ 51% │
 │                             │ ╰─────────────────────────────╯
 │                 ▁▂▃▄▅▆▇████ │ ╭─  INFO ────────────────────╮
 │         ▁▂▃▄▅▆▇████████████ │ │  Distro         Arch Linux │
 │ ▁▂▃▄▅▆▇████████████████████ │ │  Hostname            sfome │
 ╰─────────────────────────────╯ ╰─────────────────────────────╯
                       Q - quit | R - refresh
```


# config file example
```
[tui]

# time between info update in millisecond
update = 300

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


# reverence
- disk usage
  - https://stackoverflow.com/questions/20108520/get-amount-of-free-disk-space-using-go
  - https://www.includehelp.com/golang/get-the-disk-usage-information.aspx

```
│─╭╮╰╯
```

