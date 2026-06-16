# ASCII-Art-FS

## Description

ASCII-Art-FS is a Go program that takes a string and an optional banner name as input and prints it in a graphical representation using ASCII characters.

Each character is displayed using a predefined banner (standard, shadow, or thinkertoy), where every character is represented over 8 lines.

The program reads the input string from the command line and renders the corresponding ASCII art in the terminal using the Go file system API.

## Features

- Supports letters, numbers, spaces, and special characters
- Handles multi-line input using `\n`
- Uses different banner styles (standard, shadow, thinkertoy)
- Reads banner files using the Go `io/fs` package
- Maps characters using ASCII values

## Usage

```bash
go run . "Hello"
go run . "Hello" standard
go run . "Hello There!" shadow
go run . "Hello There!" thinkertoy
```
