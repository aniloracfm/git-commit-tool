package main

import (
	"bufio"
	"strings"
	"testing"
)

// TestReadOption testa as entradas de txt na CLI
// t (*testing.T) -> add pelo Go p gerenciar toda a execucao
func TestReadOption(t *testing.T) {

	// Tabela com cenários de teste
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Converte entrada maiúscula 'Y' para minúscula",
			input:    "Y\n",
			expected: "yes",
		},
		{
			name:     "Remove espaços e quebras de linha de 'yes'",
			input:    "yes \n",
			expected: "yes",
		},
		{
			name:     "Trata tecla enter (string vazia)",
			input:    "\n",
			expected: "",
		},
	}
	// Iteração sobre cada cenário
	// tt (table test/target test) -> item atual do loop na tab
	// t.Run -> cria um subteste isolado com o nome do cenário
	// strings.NewReader -> Simula a entrada do teclado add uma str na memoria
	// got -> resultado obtido (convenção)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//Simula a entrada do usuário via stdio
			inputReader := strings.NewReader(tt.input)
			reader := bufio.NewReader(inputReader)

			got := readOption(reader, "Entrada Teste")

			if got != tt.expected {
				t.Errorf("readOption() = %q, esperava %q", got, tt.expected)
			}
		})
	}
}
