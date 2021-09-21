# midget

Midget is a package manager for Friday Night Funkin' mods (basically a mod manager but without a run feature)

Packages repo: [here](github.com/MatusOllah/midget-pkgs)

## Build instructions

1. Clone repo

```
git clone https://github.com/MatusOllah/midget
```

If you are on windows then do this in the git bash

2. Change options in the Makefile

On windows x64:

```
ARCH = amd64
OS = windows
```

On mac:

```
ARCH = arm64
OS = darwin
```

On linux x64:

```
ARCH = amd64
OS = linux
```

3. Build

```
make
```

4. Install

On linux:

```
sudo make install
```