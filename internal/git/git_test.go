package git

import (
	"strings"
	"testing"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

// TestGetStageDiff_NoStagedChanges valida o comportamento quando nao ha arquivos na staging area
// A função é chamada sem garantir 'git add' prévio
func TestGetStageDiff_NoStagedChanges(t *testing.T) {

	content, hasChanges, err := GetStageDiff()

	// 1. Verifica se a execução do git não retornou erro de execução do SO
	if err != nil {
		t.Fatalf("esperava nao ter erro de execucao, mas obteve: %v", err)
	}

	// 2. Se não houver nada no staging, valida a mensagem de aviso padrão
	if !hasChanges {
		if !strings.Contains(content, messages.GitNoChangesFound) {
			t.Errorf("conteudo retornado = %q, esperava conter a mensagem de aviso %q", content, messages.GitNoChangesFound)
		}

		if !strings.Contains(content, messages.GitTipAdd) {
			t.Errorf("conteudo retornado = %q, esperava conter a dica %q", content, messages.GitTipAdd)
		}
	}
}

// TestGetStageDiff_Cenarios usa Table-Driven Tests para validar retornos da GetStageDiff
func TestGetStageDiff_Cenarios(t *testing.T) {
	cenarios := []struct {
		nomeCenario     string
		validarMensagem func(t *testing.T, conteudo string, temMudancas bool, err error)
	}{
		{
			nomeCenario: "retorna erro nulo ao executar git diff",
			validarMensagem: func(t *testing.T, conteudo string, temMudancas bool, err error) {
				if err != nil {
					t.Errorf("esperava erro nil, mas obteve %v", err)
				}
			},
		},
		{
			nomeCenario: "garante retorno nao vazio no conteudo da resposta",
			validarMensagem: func(t *testing.T, conteudo string, temMudancas bool, err error) {
				if strings.TrimSpace(conteudo) == "" {
					t.Errorf("esperava que conteudo nao fose vazio")
				}
			},
		},
	}

	for _, cenario := range cenarios {
		t.Run(cenario.nomeCenario, func(t *testing.T) {
			conteudo, temMudancas, err := GetStageDiff()
			cenario.validarMensagem(t, conteudo, temMudancas, err)
		})
	}
}
