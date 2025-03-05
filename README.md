# Password Generator

Welcome to my **Password Generator** project! This handy command-line tool, I built in Go, which allows you to generate secure, customizable passwords in a flash. I built this as a pet project to practice go and provide myself a tool a may use from time to time.

## Overview

**Password Generator** is a lightweight utility that uses Go’s cryptographic packages to produce strong, unpredictable passwords. With built-in options to include digits and special characters, you can tailor your passwords exactly the way you like.

## Features

- **Customizable Length:** Easily set your desired password length.
- **Include Digits & Special Characters:** Use simple flags to ensure your password includes at least one digit and one special character.
- **Cryptographically Secure:** Powered by Go’s `crypto/rand`, ensuring your password is generated with true randomness.
- **Straightforward & Fast:** Get your secure password in just one command.

## Getting Started

### Prerequisites

- **Go:** Ensure you have Go 1.13 or higher installed.  
  [Download Go](https://golang.org/dl/)

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/mechezarreta3/password-generator.git
cd password-generator
```

### Building the Project
Build the executable from the root directory (where your go.mod file is located):

```bash
go build -o password-generator ./cmd/password-generator
```
This command compiles the code and creates an executable named password-generator.

## Usage
Run the executable from your terminal. The tool comes with default settings, but you can customize it with flags.

### Default Settings
Without any flags, the tool generates a 20-character password based on alphabetic characters 

```bash
./password-generator
```

### Custom Options
Customize your password using flags:

`-l`: Password length (**Default: 20**)

`-d`: Include at least one digit (**Default: false**)

`-s`: Include at least one special character (**Default: false**)

For example, to generate a 16-character password that includes both digits and special characters:

```bash
./password-generator -l 16 -d -s
```
