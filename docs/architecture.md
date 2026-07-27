# Git Commit Tool - Specification and System Design

<div style="text-align: justify;">

## 1. Context and System Purpose

Maintaining a standardized commit history requires strict compliance with guidelines such as **Conventional Commits v1.0.0**. Inconsistent or vague messages impair change audits, issue tracking, and automated changelog or release workflows.

To solve this, the Git Commit Tool is a Go-based CLI application that analyzes staged modifications (`git diff --cached`), enforces local payload limits, applies embedded system prompt rules (`go:embed`), issues requests to the Google Gemini language model (`gemini-flash-latest`), and provides an interactive terminal interface for developer review prior to commit execution
  
</div>

---

## 2. Package Structure and Module Responsibilities

```text
git-commit-tool/
├── cmd/
│   └── cli/
│       ├── main.go                      # Application entry point for the CLI
│       └── main_test.go                 # Unit tests for CLI interaction logic
├── internal/
│   ├── ai/
│   │   ├── constants.go                 # System limits, timeouts, and configuration values
│   │   ├── gemini.go                    # AI API integration layer and HTTP execution
│   │   ├── gemini_test.go               # Unit tests for AI validations and prompt creation
│   │   └── prompts/
│   │       └── conventional_commits.txt # System Prompt loaded via go:embed
│   ├── git/
│   │   ├── git.go                       # Git command execution wrapper (captures staged diff)
│   │   └── git_test.go                  # Unit tests for Git operations
│   └── messages/
│       └── messages.go                  # Application-wide constants, labels, and feedback messages
├── bin/                                 # Compiled binary outputs (git-ignored)
├── Makefile                             # Build, test, and execution automation commands
├── .env.example                         # Environment variables configuration template
├── go.mod                               # Go module definition
├── LICENSE                              # MIT License
└── README.md                            # Project documentation
```

### Package Responsibilities

<div style="text-align: justify;">

- `cmd/cli`: Entry point for the executable. Handles command-line arguments, user input/output streams, and orchestrates calls between internal packages.
- `internal/ai`: Encapsulates API integration logic with Google Gemini, prompt embedding, and Fail-Fast payload validations.
- `internal/git`: Safely executes system Git commands via os/exec separating binaries from arguments to prevent Command Injection vulnerabilities.
- `internal/messages`: Centralized dictionary for all system labels, errors, and user prompts, enabling clear separation of concerns.

</div>

---

## 3. System Prompt and Guardrails Architecture

<div style="text-align: justify;">

The system prompt enforces strict constraints to ensure generated commit messages adhere to the Conventional Commits specification and security standards.

- **SYSTEM ROLE:** Defines the identity and domain scope of the LLM as an expert Software Engineer.
- **GUIDELINES:** Enforces invariants such as character limits (<= 72), lowercase casing, imperative mood, and zero trailing punctuation.
- **SCOPE RULES:**
  - **Single File/Package Changed:** Enforces concise scope inclusion (e.g., `feat(config): ...`).
  - **Multiple Files/Packages Changed:** Enforces scope omission to prevent log noise (e.g., `feat: ...`).
- **SECURITY & QUALITY GUARDRAILS:** Instructs the model to ignore lockfiles (`go.sum`, `package-lock.json`) and prevents API keys, tokens, or credentials from being reflected in commit messages.
- **OUTPUT FORMAT:** Guarantees plain-text output without Markdown code blocks or conversational text for safe CLI parsing.

</div>

---

## 4. Testing Strategy and Isolation

<div style="text-align: justify;">

The testing suite utilizes the standard Go `testing` package and follows the **Table-Driven-Tests** pattern to ensure deterministic, efficient, and isolated execution.

### 4.1. Unit Tests

- **Gemini API Isolation:** All calls to the Gemini model are decoupled using Go interface abstractions. Mock implementations simulate LLM responses without making real network calls or consuming API quota.
- **Git Command Mocks:** Standard Git operations are abstracted to prevent side effects on the working environment during test execution.

### 4.2. Integration Tests

- **Temporary Git Environments:** Tests that require real Git execution use Go's native `t.TempDir()` function to initialize temporary Git repositories on the fly.
- **Automatic Cleanup:** Temporary repositories, files, and staged changes are completely destroyed after each test run, ensuring no leakage into the developer's local environment.

</div>
