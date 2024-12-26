package middleware

type ContextKey string

const (
	ClientIDKey ContextKey = "clientID"
	PurposeKey  ContextKey = "purpose"
)
