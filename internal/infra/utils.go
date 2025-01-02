package utils

import (
	"time"
)

type Utils interface {
	CurrentMillis() int64
}

// CurrentMillis Retorna o datetime UTC atual em millisegundos
func CurrentMillis() int64 {
	// Obtém a hora atual
	now := time.Now().UTC()
	// Converte para timestamp em milissegundos
	millis := now.UnixNano() / int64(time.Millisecond)
	return millis
}
