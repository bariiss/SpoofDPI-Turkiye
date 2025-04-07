# Installation Guide

## Table of Contents

<!--ts-->
   * [Binary](#binary)
<!--te-->

## Binary
SpoofDPI will be installed in `~/.spoofdpi/bin`.  
To run SpoofDPI in any directory, add the line below to your `~/.bashrc || ~/.zshrc || ...`
```bash
export PATH=$PATH:~/.spoofdpi/bin
```
```bash
# macOS Intel
curl -fsSL https://raw.githubusercontent.com/bariiss/SpoofDPI-Turkiye/main/install.sh | bash -s darwin-amd64

# macOS Apple Silicon
curl -fsSL https://raw.githubusercontent.com/bariiss/SpoofDPI-Turkiye/main/install.sh | bash -s darwin-arm64

# linux-amd64
curl -fsSL https://raw.githubusercontent.com/bariiss/SpoofDPI-Turkiye/main/install.sh | bash -s linux-amd64

# linux-arm
curl -fsSL https://raw.githubusercontent.com/bariiss/SpoofDPI-Turkiye/main/install.sh | bash -s linux-arm

# linux-arm64
curl -fsSL https://raw.githubusercontent.com/bariiss/SpoofDPI-Turkiye/main/install.sh | bash -s linux-arm64
