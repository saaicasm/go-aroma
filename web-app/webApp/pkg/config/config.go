package config

import "html/template"

// App Config
type AppConfig struct {
	UseCache bool
	TemplateCache map[string]*template.Template
}
