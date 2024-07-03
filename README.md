# Project Title

DevToys implemented in the golang language.

## Description

This project is a command-line tool written in Go that is similar to DevToys, which can process text data in various ways.

## Features

- Process and transform JSON data
- Encode and decode hashes
- Encrypt and Decrypt data
- Format json, yaml, html text

## Getting Started

### Dependencies

- Go (version 1.20 or higher)

### Installing

- Clone the repository to your local machine
- Navigate to the project directory
- Use `go build` to compile the project

### Executing program

- Run the compiled binary with the desired flags and inputs
  - To read json from stdin, and pretty format: `./godevtoy -json`
  - To read from a file and encode to base64: `./godevtoy -file <path> -base64`
  - To encode raw string to base64: `./godevtoy -str "abcde" -base64`

## Help

Any advise for common problems or issues.
```
```markdown
command to run if program contains helper info:
```
```markdown
./your_binary_name --help
```

## Authors

Contributors names and contact info

- [bitsf](https://github.com/bitsf)

## Version History

- 0.1
    - Initial Release

## License

This project is licensed under the [MIT License](LICENSE.md) - see the LICENSE file for details

## Acknowledgments

Inspiration, code snippets, etc.
