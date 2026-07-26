package messages

const (
	// ============================================================================
	// Geral
	// ============================================================================
	Divider = "--------------------------------------------------"

	// ============================================================================
	// main.go
	// ============================================================================
	MainPromptShowDiff     = "Deseja exibir as alterações? Digite 'y' para sim ou 'enter' para continuar. "
	MainPromptAcceptCommit = "Deseja realizar o commit? (y = sim / e = editar / N = cancelar): "
	MainPromptEditCommit   = "Digite a nova mensagem de commit: "
	MainCommitSuccess      = "Commit realizado com sucesso!"
	MainCommitCanceled     = "Operação cancelada pelo usuário."

	// ============================================================================
	// git.go
	// ============================================================================
	GitSearchingChanges = "Buscando alterações..."
	GitErrorGitDiff     = "Erro ao obter diferenças do Git: %v"
	GitNoChangesFound   = "Nenhuma alteração encontrada."
	GitTipAdd           = "Dica: Execute 'git add <seu-arquivo-alterado>' antes do run."

	// ============================================================================
	// gemini.go
	// ============================================================================
	AIGenerating         = "Gerando sugestão ..."
	AISuggestionHeader   = "\nSugestão de Commit:"
	AIErrorDiffTooLarge  = "git diff muito grande (%d bytes).\nO limite é %d bytes."
	AIErrorEmptyDiff     = "O git diff fornecido está vazio."
	AIErrorMissingAPIKey = "A variável de ambiente GEMINI_API_KEY não foi configurada."
	AIPromptDiffFormat   = "Analise este git diff e gere a mensagem de commit:\n\n%s"
	AIErrorCallFailed    = "Falha ao se comunicar com o Gemini: %v"
	AIErrorNoContent     = "Sem sugestão."
	AIErrorMarshalJSON   = "erro ao montar JSON para a API: %w"
	AIErrorDecodeJSON    = "erro ao decodificar resposta JSON do Gemini: %w"
	AIErrorHTTPStatus    = "erro na chamada HTTP do Gemini: status %d"
)
