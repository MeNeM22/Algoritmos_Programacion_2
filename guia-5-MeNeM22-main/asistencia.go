package guia5

import (
	"github.com/untref-ayp2/data-structures/bitmap"
	"github.com/untref-ayp2/data-structures/set"
)

// Asistencias es una estructura que representa la asistencia de los alumnos a las clases
type Asistencias struct {
	alumnos      []*bitmap.BitMap //[1a][a2]...
	totalClases  uint8
	totalAlumnos uint
}

// NewAsistencias crea un nuevo registro de asistencias
func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias {
	var asistencia Asistencias
	asistencia.alumnos = make([]*bitmap.BitMap, cantAlumnos)
	asistencia.totalClases = cantClases
	asistencia.totalAlumnos = cantAlumnos
	for i := 0; i < len(asistencia.alumnos); i++ {
		asistencia.alumnos[i] = bitmap.NewBitMap()
	}
	return &asistencia
}

// RegistrarAsistencia registra la asistencia de un alumno a una clase
// Si el alumno o la clase no existen, no hace nada
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {
	if alumno <= a.totalAlumnos && clase <= a.totalClases {
		a.alumnos[alumno].On(clase)
	}
}

// Asistio devuelve true si el alumno asistió a la clase indicada
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	aux := false
	aux, _ = a.alumnos[alumno].IsOn(clase)
	return aux
}

// ListarClase devuelve un set con los alumnos que asistieron a la clase indicada
func (a *Asistencias) ListarClase(clase uint8) set.Set[uint] {
	lista := set.NewListSet[uint]()
	if clase <= a.totalClases {
		for i := 0; i < int(a.totalAlumnos); i++ {
			if a.Asistio(uint(i), clase) {
				lista.Add(uint(i))
			}
		}
	}
	return lista
}

// ListarAlumno devuelve un set con las clases a las que asistió el alumno indicado
func (a *Asistencias) ListarAlumno(alumno uint) set.Set[uint8] {
	lista := set.NewListSet[uint8]()
	if alumno <= a.totalAlumnos {
		for i := 0; i < int(a.totalClases); i++ {
			if a.Asistio(alumno, uint8(i)) {
				lista.Add(uint8(i))
			}
		}
	}
	return lista
}

func (a *Asistencias) Intersecti(s1, s2 *set.ListSet[uint8]) *set.ListSet[uint8] {
	lista := set.NewListSet[uint8]()
	for _, aux := range s2.Values() {
		if s1.Contains(aux) {
			lista.Add(aux)
		}
	}
	return lista
}

// ListarAsistencias devuelve un set con las clases a las que asistieron todos los alumnos
func (a *Asistencias) ListarAsistencias() set.Set[uint8] {
	lista := set.NewListSet[uint8]()
	for i := 0; i <= len(a.alumnos); i++ {
		perfect := true
		alumnosPresentes := a.ListarClase(uint8(i))
		for j := 0; j <= int(a.totalAlumnos); j++ {
			if !alumnosPresentes.Contains(uint(i)) {
				perfect = false
			}
		}
		if perfect {
			lista.Add(uint8(i))
		}
	}
	return lista
}

// ListarAsistenciaPerfecta devuelve un set con los alumnos que asistieron a todas las clases
func (a *Asistencias) ListarAsistenciaPerfecta() set.Set[uint] {
	return
}

/*// Asistencias es una estructura que representa la asistencia de los alumnos a las clases
type Asistencias struct {
	alumnos      []*bitmap.BitMap //[1a][a2]...
	totalClases  uint8
	totalAlumnos uint
}

// NewAsistencias crea un nuevo registro de asistencias
func NewAsistencias(cantAlumnos uint, cantClases uint8) *Asistencias {
	var asistencia Asistencias
	asistencia.alumnos = make([]*bitmap.BitMap, cantAlumnos)
	asistencia.totalClases = cantClases
	asistencia.totalAlumnos = cantAlumnos
	for i := 0; i < len(asistencia.alumnos); i++ {
		asistencia.alumnos[i] = bitmap.NewBitMap()
	}
	return &asistencia
}

// RegistrarAsistencia registra la asistencia de un alumno a una clase
// Si el alumno o la clase no existen, no hace nada
func (a *Asistencias) RegistrarAsistencia(alumno uint, clase uint8) {
	if alumno <= a.totalAlumnos && clase <= a.totalClases {
		a.alumnos[alumno].On(clase)
	}
}

// Asistio devuelve true si el alumno asistió a la clase indicada
func (a *Asistencias) Asistio(alumno uint, clase uint8) bool {
	asistio := false

	if alumno <= a.totalAlumnos && clase <= a.totalClases {
		aux, _ := a.alumnos[alumno].IsOn(clase)
		asistio = aux
	}
	return asistio
}

// ListarClase devuelve un set con los alumnos que asistieron a la clase indicada
func (a *Asistencias) ListarClase(clase uint8) set.Set[uint] {
	lista := set.NewListSet[uint]()
	for i := 0; i < len(a.alumnos); i++ {
		if a.Asistio(uint(i), clase) {
			lista.Add(uint(i))
		}
	}
	return lista
}

// ListarAlumno devuelve un set con las clases a las que asistió el alumno indicado
func (a *Asistencias) ListarAlumno(alumno uint) set.Set[uint8] {
	lista := set.NewListSet[uint8]()
	for i := 0; i < 32; i++ {
		if a.Asistio(alumno, uint8(i)) {
			lista.Add(uint8(i))
		}
	}
	return lista
}

func (a *Asistencias) Intersecti(s1, s2 *set.ListSet[uint8]) *set.ListSet[uint8] {
	result := set.NewListSet[uint8]()

	for _, elemento := range s1.Values() {
		if s2.Contains(elemento) { // si el elemento esta tambien en s2
			result.Add(elemento) // se agrega a result
		}
	}

	return result
}

// ListarAsistencias devuelve un set con las clases a las que asistieron todos los alumnos
func (a *Asistencias) ListarAsistencias() set.Set[uint8] {
	perfect := set.NewListSet[uint8]()

	for i := 0; i < len(a.alumnos); i++ {
		alumnos := a.totalAlumnos
		perfecto := true
		alumnosPresentes := a.ListarClase(uint8(i))
		for j := 0; j < int(alumnos); j++ {
			if !alumnosPresentes.Contains(uint(j)) {
				perfecto = false

			}
		}
		if perfecto {
			perfect.Add(uint8(i))
		}
	}

	return perfect
}

// ListarAsistenciaPerfecta devuelve un set con los alumnos que asistieron a todas las clases
func (a *Asistencias) ListarAsistenciaPerfecta() set.Set[uint] {
	perfect := set.NewListSet[uint]()

	for i := 0; i < len(a.alumnos); i++ {
		clases := a.totalClases
		perfecto := true
		clasesAsistidas := a.ListarAlumno(uint(i))
		for j := 0; j < int(clases); j++ {
			if !clasesAsistidas.Contains(uint8(j)) {
				perfecto = false

			}
		}
		if perfecto {
			perfect.Add(uint(i))
		}
	}

	return perfect
}*/
