package config

import "html/template"

//AppConfig holds configuration file
//that could be shared between packages in the app
type AppConfig struct {
	TmplCache map[string]*template.Template
	UseCache  bool //development mode
}
