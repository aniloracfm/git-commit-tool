# Git Commit Tool

The **Git Commit Tool** is a command-line interface (CLI) application developed in Go to automate and standardize the commit process. Based on staged changes, the tool suggests standardized commit messages adhering to conventional commit guidelines.

## Project Structure

```text
git-commit-tool/
├── cmd/
│   └── cli/
│       ├── main.go         # Application entry point for the CLI
│       └── main_test.go    # Unit tests for CLI interaction logic
├── internal/
│   ├── ai/
│   │   ├── gemini.go       # Gemini API integration logic
│   │   └── gemini_test.go  # Unit tests for AI validations and prompt creation
│   ├── git/
│   │   ├── git.go          # Git command execution layer (captures staged diff)
│   │   └── git_test.go     # Unit tests for Git operations
│   └── messages/
│       └── messages.go     # Application-wide constants, labels, and feedback messages
├── bin/                    # Compiled binary outputs (ignored by git)
├── Makefile                # Build, test, and execution automation commands
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

## Testing 

The project utilizes Go's native testing tool with table-driven tests across package boundaries.

To execute the full unit test suite: 

```bash
make test
```
