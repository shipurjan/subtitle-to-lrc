# subtitle-to-lrc 0.1.0

[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)
[![Build and release](https://github.com/shipurjan/subtitle-to-lrc/actions/workflows/go.yml/badge.svg)](https://github.com/shipurjan/subtitle-to-lrc/actions/workflows/go.yml) 

A program that converts a subtitle file into an [`.lrc` file](https://en.wikipedia.org/wiki/LRC_(file_format))

Currently supported input formats: 
* `.srt`
* `.vtt`

## Why?

One use case is to convert podcast subtitles to the lyrics format (.lrc), which can then be played on various portable music/media players

## Usage

```text
NAME:
   subtitle-to-lrc - Convert subtitle files to .lrc format

USAGE:
   subtitle-to-lrc [options] <input-file> [output-file]

   <input-file> - the file must have an allowed subtitle extension (srt, vtt)
   [output-file] - if not provided the program will use <input-file> filename with its extension replaced by .lrc

VERSION:
   0.1.0

AUTHOR:
   Cyprian Zdebski <cyprianz5mail@gmail.com>

GLOBAL OPTIONS:
   --separator value, -s value  Separator to use to join lines when input subtitle file has multiple lines;
      .lrc files can only have one subtitle line for each timestamp (default: "  ")
   --no-length-limit, -n  Disables the length limit of a .lrc file;
      by default a .lrc file can only have a maximum length of 59:59.99
      (some players may not support longer durations) (default: false)
   --help, -h     show help
   --version, -v  print the version

```

## Download

Windows 
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_windows-amd64.zip)
<sub><sup>[32-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_windows-386.zip)</sup></sub>
(ARM: 
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_windows-arm64.zip)
<sub><sup>[32-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_windows-arm.zip)</sup></sub>
)

Linux 
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_linux-amd64.zip)
<sub><sup>[32-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_linux-386.zip)</sup></sub>
(ARM: 
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_linux-arm64.zip)
<sub><sup>[32-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_linux-arm.zip)</sup></sub>
)

MacOS
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_darwin-amd64.zip)
(ARM: 
[64-bit](https://github.com/shipurjan/subtitle-to-lrc/releases/latest/download/subtitle-to-lrc_darwin-arm64.zip)
)

## Development

It's recommended that you use [husky](https://github.com/automation-co/husky) script, 
so that the versioning is correctly updated before pushing

```console
go install github.com/automation-co/husky@latest
husky init
```