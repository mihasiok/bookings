package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCacbe      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
