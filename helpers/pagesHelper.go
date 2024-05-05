package helpers

import (
	"fmt"
	"os"
	"strings"

	"github.com/a-h/templ"
	"github.com/hbourgeot/templdais"
)

func GetComponentPages() []templdais.MenuItem {
	items := []templdais.MenuItem{}
	// Directorio que quieres explorar
	directory := "usr/local/bin/views/pages"

	// Extensión de los archivos que deseas filtrar
	fileExtension := ".templ"

	// Lista de archivos a excluir
	excludeList := map[string]bool{
		"docs.templ": true,
	}

	// Leer los archivos en el directorio
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error al leer el directorio:", err)
		return nil
	}

	// Filtrar archivos por extensión y excluyendo los específicos
	for _, file := range files {
		if file.IsDir() {
			continue // Saltar directorios
		}

		filename := file.Name()
		if strings.HasSuffix(filename, fileExtension) && !excludeList[filename] {
			var text string
			file := strings.TrimSuffix(filename, fileExtension)
			if ContainsCapital(file) {
				println(file)
				text = CapitalizeAndSplit(file)
				println(text)
			} else {
				text = CapitalizeFirst(file)
			}

			link := templ.SafeURL("/components/" + file)

			item := templdais.MenuItem{Text: text, Link: link}
			items = append(items, item)
		}
	}

	return items
}
