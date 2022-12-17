# Minecraft Server Status (Polybar module)

[![build binaries](https://github.com/bbSnavy/bar-craft/actions/workflows/build.yml/badge.svg)](https://github.com/bbSnavy/bar-craft/actions/workflows/build.yml)

Stalking servers has never been so easy on your bar

# Install

## Prebuilt Binaries

* [Download Here](https://github.com/bbSnavy/bar-craft/releases)
  - linux x86 64bit
  - linux arm 64bit

* Install the binary to PATH

## AUR

```sh
TODO
```

# Usage

## Terminal

```sh
bar-craft cubecraft.net "{server.name}: {players.now}/{players.max}"
# cubecraft.net: 854/1000
```

## Polybar config example

```ini
modules-right = barcraft
...
[module/barcraft]
type = custom/script
exec = bar-craft hypixel.net "{server.name} {players.now}/{players.max}"
interval = 15
```

## Script arguments

See `bar-craft -h` for the arguments list

## Formats

[Read here](https://github.com/bbSnavy/bar-craft/blob/main/format.go) until further documentation

# Credits

* https://github.com/Tnze/go-mc Awesome Project
