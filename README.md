# GO RELOADED

## Description

Modifying a pre-formated txt file.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)

## Installation

To get started with this project, clone the repository and install dependencies:

```bash
git clone https://github.com/xySaad/go-reloaded
cd go-reloaded
```

## Usage

```bash
go run . sample.txt result.txt
```

## Features

- **Number Using Flags**:

  - `(hex)`: Converts the preceding hexadecimal number to its decimal form.
  - `(bin)`: Converts the preceding binary number to its decimal form.

- **Text Case Modifications Using Flags**:

  - `(up)`: Converts the preceding word to uppercase.
  - `(low)`: Converts the preceding word to lowercase.
  - `(cap)`: Capitalizes the preceding word.
  - These modifiers can also specify a number of preceding words to modify, e.g., `(up, 2)`.

- **Punctuation Corrections**:

  - Standardize space around `.`, `,`, `!`, `?`, `:` and `;`.
  - Correct placement for `'`, including support for phrases within quotation marks.

- **Grammar Corrections**:
  - Automatic correction from `a` to `an` before words starting with a vowel or an `h`.
