package main

import (
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
)

func createRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob("./templates/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	views, err := filepath.Glob("./templates/views/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	components, err := filepath.Glob("./templates/components/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	for _, view := range views {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		for _, component := range components {
			layoutCopy = append(layoutCopy, component)
		}
		files := append(layoutCopy, view)
		r.AddFromFiles(filepath.Base(view), files...)
	}

	return r
}
