package template

import (
	"github.com/andersonlira/golangspell-txtdb/appcontext"
	"github.com/andersonlira/golangspell-txtdb/config"
	"github.com/golangspell/golangspell/gateway/template"
)

//getRenderer lazy loads a Renderer
func getRenderer() appcontext.Component {
	return &template.Renderer{}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Renderer, getRenderer)
}
