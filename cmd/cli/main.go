package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aniloracfm/git-commit-tool/internal/ai"
	"github.com/aniloracfm/git-commit-tool/internal/git"
	"github.com/aniloracfm/git-commit-tool/internal/messages"
	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
)

// readInput exibe o prompt e lê a entrada do usuário.
func readInput(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// readOption lê a entrada do usuário e a converte para minúsculas.
func readOption(reader *bufio.Reader, prompt string) string {
	return strings.ToLower(readInput(reader, prompt))
}

// printbox encapsula o conteúdo entre duas linhas divisórias
func printbox(content string) {
	fmt.Println(messages.Divider)
	fmt.Println(content)
	fmt.Println(messages.Divider)
}

func main() {

	// 1. Carrega as variáveis de ambiente do arquivo .env
	_ = godotenv.Load()

	printbox(messages.GitSearchingChanges)

	// 2. Obtém as alterações staged do Git
	content, hasChanges, err := git.GetStageDiff()
	if err != nil {
		log.Fatalf(messages.GitErrorGitDiff, err)
	}

	// 3. Se não tiver alterações, exibe a mensagem do pack Git e encerra
	if !hasChanges {
		printbox(content)
		return
	}

	// 4. 'Content' guarda o diff  - Escolha do usuário em ver ou não suas
	// alterações
	reader := bufio.NewReader((os.Stdin))

	if readOption(reader, messages.MainPromptShowDiff) == "y" {
		printbox(content)
	}

	// 5. Decisão do usuário: aceitar, editar ou cancelar
	fmt.Println("\n" + messages.AIGenerating)

	commitMessage, err := ai.GenerateCommitSuggestion(content)
	if err != nil {
		log.Fatalf(messages.AIErrorCallFailed, err)
	}

	fmt.Println(messages.AISuggestionHeader)
	printbox(commitMessage)

	//6. Decisão do usuário - aceitar, editar ou negar
	option := readOption(reader, messages.MainPromptAcceptCommit)

	finalCommitMessage := commitMessage

	switch option {
	case "y":
		// Aceitar a sugestão
	case "e":
		// Editar a sugestão
		prompt := promptui.Prompt{
			Label:   "Edit commit message",
			Default: commitMessage,
		}

		result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Edit canceled: %v\n", err)
			return
		}
		finalCommitMessage = result
	default:
		fmt.Println(messages.MainCommitCanceled)
		return
	}

	// 7. Executa o commit com a mensagem final
	if err := git.Commit(finalCommitMessage); err != nil {
		log.Fatalf(messages.GitErrorGitDiff, err)
	}

	fmt.Println(messages.MainCommitSuccess)
}
