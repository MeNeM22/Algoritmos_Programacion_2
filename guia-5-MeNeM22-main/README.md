# Guía 5
## Mapa de Bits

### Ejercicio 1

Implementar un registro de lluvias anuales, donde por cada mes se registran los días en que hubo precipitaciones. Usar [mapa de bits](https://pkg.go.dev/github.com/untref-ayp2/data-structures@v0.8.0/bitmap) visto en clase como estructura de base. 

Archivo [lluvias.go](./lluvias.go)
  
### Ejercicio 2

Implementar un registro de asistencia, utilizando como contenedor de datos [mapa de bits](https://pkg.go.dev/github.com/untref-ayp2/data-structures@v0.8.0/bitmap) visto en clase, para un curso que tiene n alumnos y m clases. 

Archivo [asistencia.go](./asistencia.go)
  
## Tablas de Hash

### Ejercicio 3

Dada una tabla de hash de tamaño 11 y la función de hashing h(x) = x mod 11, inserte los números 4.371, 1.323, 6.173, 4.199, 4.344, 9.679, 1.989, resolviendo los choques con:
  - el método de hashing cerrado
  - el método de hashing abierto
  
Dibuje las tablas resultantes en cada caso.

### Ejercicio 4

Dada una tabla de hash de tamaño 7 y la función de hashing h(x) = x mod 7, inserte los números 1, 8, 27, 125, 216, 343, resolviendo los choques con:
 - el método de hashing cerrado
  - el método de hashing abierto
  
 Dibuje las tablas resultantes en cada paso.
