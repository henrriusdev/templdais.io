package helpers

import "strings"

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}
	// Convierte el primer carácter a mayúscula y concatena el resto de la cadena
	return strings.ToUpper(string(s[0])) + s[1:]
}

func splitByCapital(s string) string {
	// Divide la cadena en palabras separadas por mayúsculas sin eliminarlas
	var words []string
	var word string
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			if word != "" {
				words = append(words, word)
			}
			word = string(r)
		} else {
			word += string(r)
		}
	}

	if word != "" {
		words = append(words, word)
	}

	return strings.Join(words, " ")
}

func CapitalizeAllFirst(s string) string {
	// Convierte la primera letra de cada palabra en mayúscula
	for i, r := range s {
		if i == 0 || s[i-1] == ' ' {
			s = s[:i] + strings.ToUpper(string(r)) + s[i+1:]
		}
	}

	return s
}

func CapitalizeAndSplit(s string) string {
	return CapitalizeAllFirst(splitByCapital(s))
}

func ContainsCapital(s string) bool {
	// Verifica si la cadena contiene al menos un carácter en mayúscula
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}
