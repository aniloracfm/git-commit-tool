package git

import (
	"fmt"
	"os/exec"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

// GetStageDiff executa o comando 'git diff --cached'
// para obter as diferenças entre os arquivos que foram adicionados
// ao índice (staged) e a última versão confirmada (commit).
// Ele retorna a saída do comando como uma string e um erro,
// caso ocorra algum problema durante a execução do comando.
func GetStageDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf(messages.GitErrorGitDiff, err)
	}
	return string(output), nil
}
