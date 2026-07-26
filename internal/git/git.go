package git

import (
	"bytes"
	"os"
	"os/exec"
	"strings"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

// GetStageDiff executa o comando 'git diff --cached'
// Retorna o conteúdo a ser exibido e se existem alterações.
// Se tiver vazio, uma mensagem é exibida, lembrando ao usuario o uso do
// 'git add .' antes.
func GetStageDiff() (content string, hasChanges bool, err error) {
	cmd := exec.Command("git", "diff", "--cached")

	var output bytes.Buffer
	cmd.Stdout = &output

	if err := cmd.Run(); err != nil {
		return "", false, err
	}

	diff := strings.TrimSpace(output.String())

	if diff == "" {
		emptyMessage := messages.GitNoChangesFound + "\n" + messages.GitTipAdd
		return emptyMessage, false, nil
	}

	return diff, true, nil
}

// Commit executa o comando 'git commit -m <message>' para criar um novo commit
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
