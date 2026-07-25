# Git Commit Tool

The **Git Commit Tool** is a command-line interface (CLI) application developed in Go to automate and standardize the commit process. Based on staged changes, the tool suggests standardized commit messages adhering to conventional commit guidelines.

## Project Structure

```text
git-commit-tool/
├── cmd/
│   └── cli/
│       └── main.go         # Application entry point for the CLI
├── internal/
│   ├── git/
│   │   └── git.go          # Git command execution layer (captures staged diff)
│   └── messages/
│       └── messages.go     # Application-wide constants, labels, and feedback messages
├── bin/                    # Compiled binary outputs (ignored by git)
├── Makefile                # Build and execution automation commands
├── go.mod                  # Go module definition
└── README.md               # Project documentation
```

## Execution

Before running the tool, make sure you have staged changes in Git:

```bash
git add <your altered file>
# or stage all changes:
git add .
```

### 1. setup
```bash
make setup
```
### 2. run
```bash
make run
```
