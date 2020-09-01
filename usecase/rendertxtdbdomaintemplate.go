package usecase

import (
	"github.com/andersonlira/golangspell-txtdb/appcontext"
	"github.com/andersonlira/golangspell-txtdb/domain"
	tooldomain "github.com/golangspell/golangspell/domain"
)

//RendertxtdbdomainTemplate renders the templates defined to the txtdbdomain command with the proper variables
func RendertxtdbdomainTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	globalVariables := map[string]interface{}{
		"DomainEntity": args[0],
		// "[YOUR_COMMAND_ARG_1]": args[1],
	}

	return renderer.RenderTemplate(spell, "txtdbdomain", globalVariables, nil)
}
