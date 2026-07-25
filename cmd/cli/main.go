package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aniloracfm/git-commit-tool/internal/git"
	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

func main() {
	fmt.Println("Buscando alterações...")

	diff, err := git.GetStageDiff()

	if err != nil {
		log.Fatal(err)
	}

	if diff == "" {
		fmt.Println(messages.GitNoChangesFound)
		fmt.Println(messages.GitTipAdd)
		return
	}

	fmt.Print(messages.MainPromptShowDiff)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" {
		fmt.Println(messages.Divider)
		fmt.Println(diff)
		fmt.Println(messages.Divider)
	}
}
