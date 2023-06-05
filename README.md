# Repcap
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/) ![GitHub all releases](https://img.shields.io/github/downloads/lemorz56/repcap/total) ![GitHub issues](https://img.shields.io/github/issues-raw/lemorz56/repcap)

A tcpreplay like tool with GUI for pcap files and media controls.
Built with [Go](https://go.dev/) and [Fyne](https://github.com/fyne-io/fyne).

It can be used as a CLI tool but also supports usage through it's GUI.

## Why?
As I was working on another project where I was constantly re-playing a wireshark recording (pcap file), I felt the need for a tool. Sure a python script worked, but it's hard to have fine grained control, especially if there is different parts of the file I want to replay on different occasions. Upon searching GitHub for similar tools, I found a few outdated ones which prompted me to make one myself.

## Screenshots

![App Screenshot](https://via.placeholder.com/468x300?text=App+Screenshot+Here)

## Features
*TBD*

<!-- - Light/dark mode toggle
- Cross platform -->

## Installation

You can install Repcap using Go:
```bash
go install github.com/lemorz56/repcap@latest
```

Using [scoop](https://scoop.sh):
```bash
scoop install my-project
```

or from the [GitHub releases section]()

then run it with to see instructions.
```bash
repcap -h
```

### Usage
*TBD*

---

## Contributing

Contributions are always welcome!

You'll need to install [pre-commit]() and follow [Conventional Commits]().
<!-- pre-commit install && pre-commit install --hook-type commit-msg -->
See `CONTRIBUTING.md` for ways to get started.
- Clone the repository
- Checkout branch named `feat/feature-name` or `bug/bug-name`
- Commit
- Create PR
- Wait for approval!

## Acknowledgements

 - [Fyne](https://github.com/fyne-io/fyne)
 - [PcapReplay - other](https://www.google.com) <!-- todo: link -->

## Authors

- [@lemorz56](https://www.github.com/lemorz56)
