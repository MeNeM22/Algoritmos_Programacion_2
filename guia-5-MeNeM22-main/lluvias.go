package guia5

import (
	"github.com/untref-ayp2/data-structures/bitmap"
	"github.com/untref-ayp2/data-structures/set"
)

// Definimos el tipo Mes que puede tomar valores enteros entre 0 y 11
// podemos usar directamente el nombre del Mes en vez de usar el
// valor numérico, por ejemplo: enero, febrero, marzo, etc.
type Mes uint8

const (
	enero Mes = iota
	febrero
	marzo
	abril
	mayo
	junio
	julio
	agosto
	septiembre
	octubre
	noviembre
	diciembre
)

// valida si el Mes y el dia son validos
func validaFecha(d uint8, m Mes) bool {
	switch m {
	case abril, junio, septiembre, noviembre:
		return d >= 1 && d <= 30
	case febrero:
		return d >= 1 && d <= 29
	default: // enero, marzo, mayo, julio, agosto, octubre, diciembre
		return d >= 1 && d <= 31
	}
}

// Lluvias es una estructura que representa la cantidad de lluvias caídas en un año
type Lluvias struct {
	anio [12]*bitmap.BitMap
}

// NewLluvias crea un nuevo registro de lluvias
func NewLluvias() *Lluvias {
	var lluvias Lluvias
	for i := enero; i <= diciembre; i++ {
		lluvias.anio[i] = bitmap.NewBitMap()
	}
	return &lluvias
}

// Registrar lluvia registra un dia del año en que llovió
// Si la fecha es inválida, no hace nada
func (l *Lluvias) Registrar(d uint8, m Mes) {
	if validaFecha(d, m) {
		l.anio[m].On(d)
	}

}

// Desregistrar lluvia desregistra un dia del año en que llovió
func (l *Lluvias) Desregistrar(d uint8, m Mes) {
	if validaFecha(d, m) {
		l.anio[m].Off(d)
	}

}

// Llovió devuelve si llovió en un día del año
// Si la fecha es inválida, devuelve false
func (l *Lluvias) Llovio(d uint8, m Mes) bool {
	if validaFecha(d, m) {
		aux, _ := l.anio[m].IsOn(d)
		return aux
	}
	return false
}

// ListarMes lista los días en que llovió en un Mes
func (l *Lluvias) ListarMes(m Mes) *set.ListSet[uint8] {
	dias := set.NewListSet[uint8]()
	if m <= diciembre {
		mes := l.anio[m]
		for dia := 1; dia <= 31; dia++ {
			llovio, _ := mes.IsOn(uint8(dia))
			if llovio {
				dias.Add(uint8(dia))
			}
		}

	}
	return dias
}

func Intersection(s1, s2 *set.ListSet[uint8]) *set.ListSet[uint8] {
	result := set.NewListSet[uint8]()

	for _, elemento := range s1.Values() {
		if s2.Contains(elemento) { // si el elemento esta tambien en s2
			result.Add(elemento) // se agrega a result
		}
	}

	return result
}

// Dado dos meses, listar los días que llovieron en ambos
func (l *Lluvias) ListarDiasEnAmbosMeses(m1 Mes, m2 Mes) *set.ListSet[uint8] {
	d1 := l.ListarMes(m1)
	d2 := l.ListarMes(m2)
	result := Intersection(d1, d2)
	return result

}
