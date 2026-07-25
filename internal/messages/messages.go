package messages

// ============================================================================
// Geral
// ============================================================================
const (
	Divider = "--------------------------------------------------"
)

// ============================================================================
// main.go
// ============================================================================
const (
	MainPromptShowDiff = "Deseja exibir as alterações? Digite 'y' para sim ou 'enter' para continuar. "
)

// ============================================================================
// git.go
// ============================================================================
const (
	GitSearchingChanges = "Buscando alterações..."
	GitErrorGitDiff     = "Erro ao obter diferenças do Git: %w"
	GitNoChangesFound   = "Nenhuma alteração encontrada."
	GitTipAdd           = "Dica: Execute 'git add <seu-arquivo-alterado> antes do run."
)
