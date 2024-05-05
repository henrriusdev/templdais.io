package helpers

import "strings"

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	// Convierte el primer carácter a mayúscula y concatena el resto de la cadena
	return strings.ToUpper(string(s[0])) + s[1:]
}
