package guia5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistrar(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, enero)
	lluvias.Registrar(2, enero)
	lluvias.Registrar(10, enero)
	lluvias.Registrar(1, enero)
	lluvias.Registrar(87, enero)
	assert.True(t, lluvias.Llovio(1, enero))
	assert.True(t, lluvias.Llovio(2, enero))
	assert.True(t, lluvias.Llovio(10, enero))
	assert.False(t, lluvias.Llovio(87, enero))
}

func TestIsOn(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, enero)
	assert.False(t, lluvias.Llovio(2, enero))
	assert.True(t, lluvias.Llovio(1, enero))
	lluvias.Registrar(1, diciembre)
	assert.True(t, lluvias.Llovio(1, diciembre))
}

func TestListarMes(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, enero)
	lluvias.Registrar(2, enero)
	lluvias.Registrar(10, enero)
	lluvias.Registrar(1, enero)
	lluvias.Registrar(9, enero)
	lluvias.Registrar(25, enero)
	dias := lluvias.ListarMes(enero)
	require.NotNil(t, dias)
	assert.Equal(t, 5, dias.Size())
	assert.True(t, dias.Contains(1))
	assert.True(t, dias.Contains(2))
	assert.True(t, dias.Contains(9))
	assert.True(t, dias.Contains(10))
	assert.True(t, dias.Contains(25))
}

func TestListarDiasEnAmbosMeses(t *testing.T) {
	lluvias := NewLluvias()
	lluvias.Registrar(1, enero)
	lluvias.Registrar(2, enero)
	lluvias.Registrar(10, enero)
	lluvias.Registrar(1, marzo)
	lluvias.Registrar(10, marzo)
	lluvias.Registrar(25, mayo)
	lluvias.Registrar(25, noviembre)
	dias := lluvias.ListarDiasEnAmbosMeses(enero, marzo)
	require.NotNil(t, dias)
	assert.Equal(t, 2, dias.Size())
	assert.True(t, dias.Contains(1))
	assert.True(t, dias.Contains(10))
	dias = lluvias.ListarDiasEnAmbosMeses(noviembre, mayo)
	assert.Equal(t, 1, dias.Size())
	assert.True(t, dias.Contains(25))
}
