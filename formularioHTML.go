/*
*  Author MALDRU
*  Email andres.latorre@ing-developers.com
*  Copyright (c) 2019. All rights reserved.
*/
package tools

import "strconv"

//ValidarLongitud valida la longitud de un string en un rango especificado, si max es 0 no se toma en cuenta el maximo
func ValidarLongitud(valor string, min, max int) bool {
	numCaracteres := len(valor)
	if max == 0 {
		return numCaracteres >= min
	}
	return numCaracteres >= min && numCaracteres <= max
}

//ValidarValoresPosibles valida que el valor este entre algun posible valor
func ValidarValoresPosibles(valor string, posibles ...string) bool {
	for _, posible := range posibles {
		if valor == posible {
			return true
		}
	}
	return false
}

//EsNumerico valida si es un numero valido
func EsNumerico(valor string) bool {
	_, err := strconv.ParseInt(valor, 10, 64)
	return err == nil
}
