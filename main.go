package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sarthak0714/todo-gotem/model"
	"github.com/sarthak0714/todo-gotem/templates"
)

var todos []model.Todo



func main() {
    e := echo.New()
	todos=append(todos,model.Todo{Id:1,Status:true,Task:"hii"})
	todos=append(todos,model.Todo{Id:2,Status:true,Task:"hello"})
	comp:=templates.Hello(todos)
    e.GET("/", func(c echo.Context) error {
       return comp.Render(c.Request().Context(),c.Response())
    })
    e.Logger.Fatal(e.Start(":6969"))
}


