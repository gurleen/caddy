package executor

import (
	"os/exec"
	"text/template"

	"github.com/caddyserver/caddy/v2"
)

func init() {
	caddy.RegisterModule(Executor{})
}

type Executor struct {
}

// CaddyModule returns the Caddy module information.
func (Executor) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.templates.functions.executor",
		New: func() caddy.Module { return new(Executor) },
	}
}

func (Executor) CustomTemplateFunctions() template.FuncMap {
	return template.FuncMap{
		"exec": func(cmd string) string {
			out, _ := exec.Command("bash", "-c", cmd).Output()
			return string(out)
		},
	}
}
