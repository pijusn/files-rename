# files-rename

This is a small utility to mass-rename files to have sequential names.

## Usage

Rename all files in current directory to be named `YYYYMMDD_0001.jpeg`, `YYYYMMDD_0002.jpeg`, ...

```sh
files-rename -directory . -name "YYYYMMDD_%04d"
```