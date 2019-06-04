package resolvergen

import (
	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
)

func New() plugin.Plugin {
	return &Plugin{}
}

type Plugin struct{}

var _ plugin.CodeGenerator = &Plugin{}

func (m *Plugin) Name() string {
	return "resovlergen"
}
func (m *Plugin) GenerateCode(data *codegen.Data) error {
	if !data.Config.Resolver.IsDefined() {
		return nil
	}

	resolverBuild := &ResolverBuild{
		Data:         data,
		PackageName:  data.Config.Resolver.Package,
		ResolverType: data.Config.Resolver.Type,
	}
	//filename := data.Config.Resolver.Filename

	options := templates.Options{
		PackageName: data.Config.Resolver.Package,
		Filename:    data.Config.Resolver.Filename,
		Data:        resolverBuild,
	}
	pkg := data.Config.Middleware.ImportPath()
	if len(pkg) > 0 {
		options.ExtImports = []string{pkg}
	}
	//if _, err := os.Stat(filename); os.IsNotExist(errors.Cause(err)) {
	return templates.Render(options)
	//}

	//log.Printf("Skipped resolver: %s already exists\n", filename)
	//return nil
}

type ResolverBuild struct {
	*codegen.Data

	PackageName  string
	ResolverType string
}
