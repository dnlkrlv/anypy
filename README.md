# AnyPy

`Pyenv` is required: https://github.com/pyenv/pyenv

A simple tool to execute commands in a specific Python version.

## Installation

```bash
git clone https://github.com/dnlkrlv/anypy.git
cd anypy
go build -o anypy
mv anypy /usr/local/bin/anypy
```

## Usage

### General

```bash
anypy <version> <python-command>
```

### Example

```bash
anypy 3.10.0 python --version
```
