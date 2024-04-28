package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/hbourgeot/templdais.io/views"
	"github.com/hbourgeot/templdais.io/views/layout"
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
