package ai

import (
	"os"
	"strings"
	"testing"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

// TestGenerateCommitSuggestion_Validation valida antes da chamada de API
// t (*testing.T): gerenciador de testes do Go
// Garante que a API Key não fique suja de testes anteriores no ambiente
// setupEnv -> // Função opcional para preparar o ambiente antes de rodar o cenário
// MaxDiffSize+1 -> Cria uma string maior que o limite
func TestGenerateCommitSuggestion_Validation(t *testing.T) {
	os.Unsetenv("GEMINI_API_KEY")

	// Tabela de cenários de teste para validações locais
	tests := []struct {
		name        string
		diff        string
		setupEnv    func()
		expectedErr string
	}{
		{
			name:        "Erro quando o diff estiver vazio",
			diff:        "",
			setupEnv:    nil,
			expectedErr: messages.AIErrorEmptyDiff,
		},
		{
			name:        "Erro quando o diff contiver apenas espaços em branco",
			diff:        "   \n\t  ",
			setupEnv:    nil,
			expectedErr: messages.AIErrorEmptyDiff,
		},
		{
			name:        "Erro quando o diff for maior que o limite permitido",
			diff:        strings.Repeat("a", MaxDiffSize+1),
			setupEnv:    nil,
			expectedErr: "git diff muito grande",
		},
		{
			name: "deve retornar erro quando a GEMINI_API_KEY nao estiver configurada",
			diff: "diff --git a/main.go b/main.go",
			setupEnv: func() {
				os.Unsetenv("GEMINI_API_KEY")
			},
			expectedErr: messages.AIErrorMissingAPIKey,
		},
	}

	// Iteração sobre cada cenário na tabela
	// tt (table test): representa o item/cenário atual
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnv != nil {
				tt.setupEnv()
			}

			// Chama a função real que estamos testando
			gotSuggestion, err := GenerateCommitSuggestion(tt.diff)

			// 1. O teste esperava um erro e ele não veio
			if err == nil {
				t.Fatalf("esperava um erro contendo %q, mas obteve nil (sucesso). Resultado: %q", tt.expectedErr, gotSuggestion)
			}

			// 2. Teve o erro - verifica se a mensagem contém o texto esperado
			if !strings.Contains(err.Error(), tt.expectedErr) {
				t.Errorf("mensagem de erro incorreta.\nObtido: %q\nEsperado conter: %q", err.Error(), tt.expectedErr)
			}
		})
	}
}
