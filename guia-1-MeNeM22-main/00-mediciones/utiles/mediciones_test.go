package utiles

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExistencia(t *testing.T) {
	entries, err := os.ReadDir("../resultados")
	if err != nil {
		assert.Error(t, err)
	}
	if len(entries) == 0 {
		assert.Fail(t, "Archivo no encontrado")
	}

	var maxSize int64 = 0
	for _, e := range entries {
		info, inferr := e.Info()
		if inferr != nil {
			assert.Error(t, inferr)
		}
		if info.Size() > maxSize {
			maxSize = info.Size()
		}
	}
	fmt.Print(maxSize)

	assert.NotEqual(t, int64(0), maxSize)
}
