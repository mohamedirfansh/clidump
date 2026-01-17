# clidump

CLI history logger and markdown converter.

## Installation

```bash
go install github.com/yourusername/clidump/cmd/clidump@latest
```

## Usage

```bash
clidump [command]
```

## Features

- Track CLI command history
- Sanitize sensitive data from commands
- Convert history to markdown format

## Development

```bash
# Build
go build -o clidump ./cmd/clidump

# Run tests
go test ./...

# Run locally
go run ./cmd/clidump
```

## License

See LICENSE file for details.
