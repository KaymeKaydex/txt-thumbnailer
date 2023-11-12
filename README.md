# Txt-Thumbnailer

```txt-thumbnailer``` is a service for rendering 2D graphics thumbnail for txt file in pure Go

## Overview

Txt-Thumbnailer is a service providing a simple interface to create thumbnail for txt files. 

## Examples
Try to convert 1 of examples files with example font

`$ go run cmd/txt-thumbnailer/main.go convert examples/txt/long.txt  --font-size=16  --padding-left=50 --padding-top=50 --padding-right=50 --padding-bottom=50 --font=examples/fonts/MailSansRoman-Light.ttf`
## Installation

## Usage

## Features

- [x] padding-left settings
- [x] padding-right settings
- [x] padding-top settings
- [x] height settings
- [x] width settings
- [x] font-size settings
- [x] height settings
- [x] line-spacing
- [x] fonts switch

- [ ] verbose flag
- [ ] server start
- [ ] auto-escape
- [ ] chars escape
- [ ] DPI settings
- [ ] skip useless strings for long files without padding usage
- [ ] padding bottom correct work (now it ignore one source line from txt file)