package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/henrriusdev/templdais.io/views"
	"github.com/henrriusdev/templdais.io/views/layout"
	"github.com/henrriusdev/templdais.io/views/pages"
)

type HTTPError struct {
	Status int
	Err    error
}

func NewHTTPError(status int, err error) HTTPError {
	return HTTPError{
		Status: status,
		Err:    err,
	}
}

func (he HTTPError) Error() string {
	return fmt.Sprintf("error %d %s (%s)", he.Status, he.Err.Error(), http.StatusText(he.Status))
}

func ErrorHandler() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := c.Next()
		if err == nil {
			return nil
		}

		status := http.StatusInternalServerError
		if e, ok := err.(HTTPError); ok {
			status = e.Status
		}

		// Aquí puedes renderizar tu vista personalizada en lugar de enviar un mensaje de error.
		if status == fiber.StatusNotFound {
			return Render(c, layout.NotFound(err.Error()))
		}

		return Render(c, layout.Error(status, err.Error()))
	}
}

func Render(c fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(layout.Base(component))
	for _, o := range options {
		o(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}

func home(c fiber.Ctx) error {
	return Render(c, views.Home())
}

func components(c fiber.Ctx) error {
	component := c.Params("component", "accordion")

	var page templ.Component
	switch component {
	case "accordion":
		page = pages.AccordionPage()
	case "alert":
		page = pages.AlertPage()
	case "badge":
		page = pages.BadgePage()
	case "breadcrumbs":
		page = pages.BreadcrumbPage()
	case "bottomNavigation":
		page = pages.BottomNavPage()
	case "button":
		page = pages.ButtonPage()
	case "card":
		page = pages.CardPage()
	case "checkbox":
		page = pages.CheckboxPage()
	case "code":
		page = pages.CodePage()
	case "divider":
		page = pages.DividerPage()
	case "dropdown":
		page = pages.DropdownPage()
	case "menu":
		page = pages.MenuPage()
	case "navbar":
		page = pages.NavbarPage()
	default:
		return NewHTTPError(http.StatusNotFound, fmt.Errorf("%s Page not found, maybe we are working on it", component))
	}

	return Render(c, layout.Components(page))
}
