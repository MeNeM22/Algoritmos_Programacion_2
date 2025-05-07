package funciones

// Post: Se muestra un mensaje indicando la opción seleccionada o "Opción incorrecta" si la opción no es válida.
func elegirOpcion(opcion int) int {
	var elegido = 0
	/*if opcion > 4 || opcion < 1{}*/
	switch opcion {
	case 1:
		elegido = 1
	case 2:
		elegido = 2
	case 3:
		elegido = 3
	case 4:
		elegido = 4
	}

	return elegido
}
