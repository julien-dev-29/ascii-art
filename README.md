# ASCII-Art

## Description

ASCII-Art is a Go program that takes a string as input and prints it in a graphical representation using ASCII characters.

Each character is displayed using a predefined banner (standard, shadow, or thinkertoy), where every character is represented over 8 lines.

The program reads the input string from the command line and renders the corresponding ASCII art in the terminal.

## Features

- Supports letters, numbers, spaces, and special characters
- Handles multi-line input using `\n`
- Uses different banner styles (standard, shadow, thinkertoy)
- Reads banner files and maps characters using ASCII values

## Usage

```bash
go run . "Hello"
```
