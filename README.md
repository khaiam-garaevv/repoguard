
# RepoGuard

**RepoGuard** is an open-source command-line tool written in Go that analyzes repository structure and checks for essential project files.

## Features

- Scan repository directories
- Detect missing required files
- Identify directories incorrectly used in place of required files
- Display a clear repository analysis report
- Return an appropriate exit code for automated workflows

## Requirements

- Go 1.27.4 or later

## Installation

Clone the repository:

```bash
git clone https://github.com/khaiam-garaevv/repoguard.git
cd repoguard
```

## Usage

Analyze the current directory:

```bash
go run .
```

Analyze a specific directory:

```bash
go run . ./path/to/repository
```

## Development

Run tests:

```bash
go test -count=1 ./...
```

Run static analysis:

```bash
go vet ./...
```

Build the project:

```bash
go build ./...
```

## Project Structure

```text
repoguard/
├── internal/
│   ├── checker/
│   ├── reporter/
│   └── scanner/
├── .github/
│   └── workflows/
├── main.go
├── go.mod
├── README.md
└── LICENSE
```

## Contributing

Contributions, bug reports, and suggestions are welcome.

Please open an issue before submitting major changes.

## License

This project is licensed under the MIT License.
See the [LICENSE](LICENSE) file for details.
