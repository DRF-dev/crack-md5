<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>crack-md5
</h1>
<h3>◦ Developed with the software and tools below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
</p>
<img src="https://img.shields.io/github/languages/top/DRF-dev/crack-md5?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/DRF-dev/crack-md5?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/DRF-dev/crack-md5?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/DRF-dev/crack-md5?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📖 Table of Contents
- [📖 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [📦 Features](#-features)
- [📂 Repository Structure](#-repository-structure)
- [🚀 Getting Started](#-getting-started)
    - [🔧 Installation](#-installation)
    - [🤖 Running crack-md5](#-running-crack-md5)

---


## 📍 Overview

Crack hash md5 by dictionary or incremental method

---

## 📦 Features

- Generate a md5 hash specifying a password
- Crack a md5 hash using a dictionary
- Crack a md5 hash specifying the number of letter of initial password

---


## 📂 Repository Structure

```sh
└── crack-md5/
    ├── .gitignore
    ├── cmd/
    │   └── main.go
    ├── go.mod
    ├── go.sum
    ├── internal/
    │   ├── alphabet.go
    │   ├── error.go
    │   ├── input.go
    │   ├── md5.go
    │   └── wordlist.go
    └── pkg/
        ├── hash-crack-dictionary.go
        ├── hash-crack-incremental.go
        └── random-word.go
```


---


## 🚀 Getting Started

***Dependencies***

Please ensure you have the following dependencies installed on your system:
- Golang

### 🔧 Installation

1. Clone the crack-md5 repository:
```sh
git clone https://github.com/DRF-dev/crack-md5
```

2. Change to the project directory:
```sh
cd crack-md5
```

3. Install the dependencies:
```sh
go install
```

4. Build the project:
```sh
go build -o myapp
```

### 🤖 Running crack-md5

Generate a md5 hash
```sh
./myapp -g $PASSWORD
```

Crack md5 hash using dictionary
```sh
./myapp -f /path/to/dictionary -md5 $MD5_HASH
```

Crack using incremental method
```sh
./myapp -l $PASSWORD_LENGTH -md5 $MD5_HASH
```