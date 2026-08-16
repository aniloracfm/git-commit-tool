# Git Commit Tool

<div style="text-align: justify;">

Maintaining a clean and standardized commit history ensures code traceability, improves code reviews, and unlocks automated changelog generation and release management.

Keeping this in mind, **git-commit-tool** was developed as a lightweight command-line interface **(CLI)** application in Go to **automate** and **standardize Git commits** using Artificial Intelligence.

The tool automatically analyzes your code changes to generate semantic commit messages adhering strictly to the **Conventional Commits v1.0.0** specification, eliminating the need to write commit messages manually.

---

## Features

- **Conventional Commits Compliance:** Enforces semantic commit types (`feat`, `fix`, `docs`, `refactor`, `test`).
- **Fully Automated Workflow:** Automatically triggers on `git add` to process repository changes and present ready-to-use semantic suggestions instantly.
- **Dynamic Scope Ingestion:** Automatically includes parenthetical scope for single-module changes and omits it on multi-file diffs to prevent log clutter.
- **Interactive Terminal Workflow:** Review generated messages, edit them inline, inspect diffs on demand, or cancel cleanly before committing.
- **Fail-Fast Local Guardrails:** Blocks oversized diffs locally to ensure cost control and resource protection.
- **Shift-Left Security:** Built-in prompt guardrails prevent API keys, tokens, or credentials from leaking into commit history.
- **Binary Distribution:** Static compilation in Go bundles dependencies into the final executable, enabling direct execution without local interpreters.

</div>

---

## Prerequisites

- **Go:** Version 1.22 or higher.
- **Git:** Installed and configured.
- **Gemini API Key:** Key obtained from [Google AI Studio](https://aistudio.google.com/api-keys).

---

## Configuration

**1. Create a `.env` file at the root of the project based on `.env.example`:**

```bash
cp .env.example .env
```

**2. Add your Gemini API key inside `.env`:**

```bash
GEMINI_API_KEY=your_actual_api_key_here
```

**3. Install the shell integration:**

```bash
make setup
```

---

## Execution

**1. Stage your changes in Git based on your workflow:**

- To stage all modified files for a single commit:

   ```bash
   git add .
   ```

- To stage a specific file for a granular commit:

   ```bash
   git add <file_path>
   ```

**1.1. Interactive Workflow:**

- **Diff Inspection:**

  - Press `y` to view the staged changes directly in the terminal.
  - Press `Enter` (default) to skip and proceed to message generation.

- **Commit Action Selection:**

  - **Accept**(`y`): Executes the commit with the suggested conventional commit message.
  - **Edit**(`e`): Opens an inline prompt pre-filled with the suggestion for manual tweaks before committing.
  - **Reject**(`n`): Cancels the operation safely without modifying repository state.

---

## Testing

This project uses Go's standard `testing` package along with unit and integration tests to ensure CLI reliability, command execution, and prompt structure integrity.

To run the full test suite via `Makefile`:

```bash
make test
```

---

## System Architecture

For in-depth architectural decisions, Clean Architecture boundaries, security hardening, prompt engineering contracts, and scalability roadmap, please refer to [here](../git-commit-tool/docs/architecture.md).

## License

Distributed under the MIT License.
