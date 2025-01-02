package logger

import (
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger()
	if Log == nil {
		t.Errorf("Logger não foi inicializado corretamente")
	}
}
