package main

import (
	"html/template"
	"io"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	var allFiles []string

	err := filepath.WalkDir("web/views", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			allFiles = append(allFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Error walking views:", err)
	}

	log.Println("Parsing templates:", allFiles)

	return &Templates{
		templates: template.Must(template.ParseFiles(allFiles...)),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	e.Static("/static", "web/static")

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index.html", nil)
	})
	
	e.Logger.Fatal(e.Start(":3000"))
}