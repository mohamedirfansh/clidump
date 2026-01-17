# clilog

CLI history logger and markdown converter.

## Installation

```bash
go install github.com/yourusername/clilog/cmd/clilog@latest
```

## Usage

```bash
clilog [command]
```

## Features

- Track CLI command history
- Sanitize sensitive data from commands
- Convert history to markdown format

## Development

```bash
# Build
go build -o clilog ./cmd/clilog

# Run tests
go test ./...

# Run locally
go run ./cmd/clilog
```

## License

See LICENSE file for details.
