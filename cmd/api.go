package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/hbourgeot/templdais.io/views"
	"github.com/hbourgeot/templdais.io/views/layout"
	"github.com/hbourgeot/templdais.io/views/pages"
)

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
	case "breadcrumbs":
		page = pages.BreadcrumbPage()
	case "button":
		page = pages.ButtonPage()
	case "card":
		page = pages.CardPage()
	case "menu":
		page = pages.MenuPage()
	case "navbar":
		page = pages.NavbarPage()
	}

	return Render(c, layout.Components(page))
}
