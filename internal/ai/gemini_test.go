package ai

import (
	"os"
	"strings"
	"testing"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

// TestGenerateCommitSuggestion_Validation tests local validations before making API calls
func TestGenerateCommitSuggestion_Validation(t *testing.T) { // t (*testing.T): gerenciador de testes do Go
	// Garante que a API Key não fique suja de testes anteriores no ambiente
	os.Unsetenv("GEMINI_API_KEY")

	// Tabela de cenários de teste para validações locais
	tests := []struct {
		name        string
		diff        string
		setupEnv    func() // Função opcional para preparar o ambiente antes de rodar o cenário
		expectedErr string
	}{
		{
			name:        "deve retornar erro quando o diff estiver vazio",
			diff:        "",
			setupEnv:    nil,
			expectedErr: messages.AIErrorEmptyDiff,
		},
		{
			name:        "deve retornar erro quando o diff contiver apenas espaços em branco",
			diff:        "   \n\t  ",
			setupEnv:    nil,
			expectedErr: messages.AIErrorEmptyDiff,
		},
		{
			name:        "deve retornar erro quando o diff for maior que o limite permitido",
			diff:        strings.Repeat("a", MaxDiffSize+1), // Cria uma string maior que o limite
			setupEnv:    nil,
			expectedErr: "git diff muito grande",
		},
		{
			name: "deve retornar erro quando a GEMINI_API_KEY nao estiver configurada",
			diff: "diff --git a/main.go b/main.go",
			setupEnv: func() {
				os.Unsetenv("GEMINI_API_KEY") // Garante que a chave esteja ausente
			},
			expectedErr: messages.AIErrorMissingAPIKey,
		},
	}

	// Iteração sobre cada cenário na tabela
	for _, tt := range tests { // tt (table test): representa o item/cenário atual
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupEnv != nil {
				tt.setupEnv()
			}

			// Chama a função real que estamos testando
			gotSuggestion, err := GenerateCommitSuggestion(tt.diff)

			// Validação 1: O teste esperava um erro e ele não veio
			if err == nil {
				t.Fatalf("esperava um erro contendo %q, mas obteve nil (sucesso). Resultado: %q", tt.expectedErr, gotSuggestion)
			}

			// Validação 2: O erro veio, mas precisamos verificar se a mensagem contém o texto esperado
			if !strings.Contains(err.Error(), tt.expectedErr) {
				t.Errorf("mensagem de erro incorreta.\nObtido: %q\nEsperado conter: %q", err.Error(), tt.expectedErr)
			}
		})
	}
}
