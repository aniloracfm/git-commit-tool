package ai

import "time"

const (
	// Limites
	MaxDiffSize = 20000 // 20 KB

	// Endpoints e Configurações
	GeminiBaseURL    = "https://generativelanguage.googleapis.com"
	GeminiApiVersion = "v1beta"
	GeminiModel      = "gemini-flash-latest"
	GeminiEndpoint   = GeminiBaseURL + "/" + GeminiApiVersion + "/models/" + GeminiModel + ":generateContent?key=%s"
	ContentTypeJSON  = "application/json"
	HTTPTimeout      = 10 * time.Second
)
