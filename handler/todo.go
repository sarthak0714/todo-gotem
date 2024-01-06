package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/todo-gotem/model"
)

var Todos []model.Todo



func GetTodos(c echo.Context, comp templ.Component) error {
	
	return comp.Render(c.Request().Context(), c.Response())
}
