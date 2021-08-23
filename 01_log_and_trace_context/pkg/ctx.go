package pkg

import "fmt"

// ContextKey defines a context key type.
type ContextKey string

func (c ContextKey) String() string {
	return fmt.Sprintf("%s%s", contextKeyPrefix, string(c))
}

const (
	// contextKeyPrefix defines context key prefix for the package (avoid keys collision with other packages).
	contextKeyPrefix = "practicumLogging-"
)
