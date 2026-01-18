# clidump

Intelligent AI powered CLI assistant helping you convert natural language into CLI commands and helps you easily archive unfamiliar CLI commands automatically for future reference.

## Features

- It converts natural language input into appropriate CLI commands and pastes it directly in the next line ready to be executed.
```bash
clidump % ct "List all the files in the current directory"
clidump % ls -la
clidump % 
clidump % ct "Delete all Evicted Pods across all namespaces"
clidump % kubectl delete pod --all --field-selector=status.phase=Failed -A
```

- Archive previous N commands into a neat archive-version-number.md file with the explanation of what each command does for future reference into current directory.
```bash
clidump % ct dump
clidump % ls
clidump-1.md
```

## Installation

1. Download the latest release of the clidump binary from: https://github.com/mohamedirfansh/clidump/releases
2. Add the location of the binary to your PATH (add it in .bashrc or .zshrc for it to be permenant)
```bash
export PATH=/location/of/binary:$PATH
```
3. Run these 2 commands:
```bash
clidump --install
source ~/.bashrc # or ~/.zshrc (for mac)
```

## Usage

```bash
# Command from english
ct "<what you want to execute in natural english>"

# Dump your old commands with explanations
ct dump
```

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

[![GitHub](https://img.shields.io/github/license/mohamedirfansh/clidump)](https://github.com/mohamedirfansh/clidump/blob/master/LICENSE)

This project is licensed under the **[MIT License](http://opensource.org/licenses/mit-license.php)** - see the [LICENSE](https://github.com/mohamedirfansh/clidump/blob/master/LICENSE) file for more details.

---

> This project was built during the Hack&Roll 2026 Hackathon. View the Devpost submission [here](https://devpost.com/software/clidump)