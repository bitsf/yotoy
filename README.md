# Project Title

YYToy is a CLI based Swiss Army Knife for Developer implemented in the golang language. Inspired by [DevToys](https://github.com/DevToys-app/DevToys).

## Description

YYToy is a command-line tool written in Go that is similar to DevToys, which can process text data in various ways and process text in pipeline.

## Features

- Process and transform JSON data
- Encode and decode hashes
- Encrypt and Decrypt data
- Format json, yaml, html text
- Process data in pipeline

## Getting Started

### Dependencies

- Go (version 1.20 or higher)

### Installing

- Clone the repository to your local machine
- Navigate to the project directory
- Use `go build` to compile the project

### Executing program

- Run the compiled binary with the desired flags and inputs
  - To read json from stdin, and pretty format: `./yytoy -json`
  - To read from a file and encode to base64: `./yytoy -file <path> -base64`
  - To encode raw string to base64: `./yytoy -str "abcde" -base64`

## Help

Any advice for common problems or issues.

## Authors

Contributors names and contact info

- [bitsf](https://github.com/bitsf)

## Version History

- 0.1
    - Initial Release

## License

This project is licensed under the [MIT License](LICENSE) - see the LICENSE file for details

## Acknowledgments

Inspired by [DevToys](https://github.com/DevToys-app/DevToys).
