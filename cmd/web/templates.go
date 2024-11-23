package main

import (
	"html/template"
	"path/filepath"
)

func setTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := []string{
			"./ui/html/index.html",
			"./ui/html/nav/nav.html",
			page,
		}

		templateSet, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = templateSet
	}

	return cache, nil
}
