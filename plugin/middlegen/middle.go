package middlegen

import (
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/pkg/errors"
	"log"
	"os"
)

func New() plugin.Plugin {
	return &Plugin{}
}

type Plugin struct{}

var _ plugin.CodeGenerator = &Plugin{}

func (m *Plugin) Name() string {
	return "middlegen"
}
func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Middleware.IsDefined() {
		return nil
	}

	middleBuild := &MiddleBuild{
		Data:         data,
		PackageName:  data.Config.Middleware.Package,
	}
	filename := data.Config.Middleware.Filename

	options := templates.Options{
		PackageName: data.Config.Middleware.Package,
		Filename:    filename,
		Data:        middleBuild,
	}
	if _, err := os.Stat(filename); os.IsNotExist(errors.Cause(err)) {
		return templates.Render(options)
	}

	log.Printf("Skipped resolver: %s already exists\n", filename)
	return nil
}

type MiddleBuild struct {
	*codegen.Data

	PackageName  string
}
