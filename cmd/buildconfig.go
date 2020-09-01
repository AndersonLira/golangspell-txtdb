package cmd

import (
	"fmt"

	"encoding/json"

	"github.com/golangspell/golangspell/domain"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["build-config"] = runBuildConfigCommand
}

func runBuildConfigCommand(cmd *cobra.Command, args []string) {
	configBytes, err := json.MarshalIndent(buildSpellConfig(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(configBytes))
}

func buildSpellConfig() domain.Spell {
	return domain.Spell{
		Name: "golangspell-txtdb",
		URL:  "github.com/andersonlira/golangspell-txtdb",
		Commands: map[string]*domain.Command{
			"build-config": &domain.Command{
				Name:             "build-config",
				ShortDescription: "Builds the config necessary for adding this plugin to the Golang Spell tool",
				LongDescription: `Builds the config necessary for adding this plugin to the Golang Spell tool.
This command must be available in all Golang Spell plugins to make it possible the plugin addition to the platform.

Syntax: 
golangspell build-config
`,
			},
			"golangspell-txtdb-hello": &domain.Command{
				Name:             "golangspell-txtdb-hello",
				ShortDescription: "The golangspell-txtdb-hello says Hello! using your new Golangspell base structure",
				LongDescription: `The golangspell-txtdb-hello says Hello! using your new Golangspell base structure
The Architectural Model is based in the Clean Architecture and is the basis to add more resources like domain models and repositories.
You can use this as a template to create your own commands. 
Please notice that ALL your commands must be prefixed with the name of your Spell (golangspell-txtdb). It will avoid name colision with the Spells from other authors 
Args:
name: Your name (required) to be added to the Hello!. Example: Elvis"

Syntax: 
golangspell golangspell-txtdb-hello [name]
`,
				ValidArgs: []string{"name"},
			},
		"txtdbdomain": &domain.Command{
	Name:             "txtdbdomain",
	ShortDescription: "The txtdbdomain [TODO: PUT HERE THE NEW COMMAND FEATURE]",
	LongDescription: `The txtdbdomain [TODO: PUT HERE THE NEW COMMAND FEATURE EXTENDED DESCRIPTION]
Args:
[TODO: PUT HERE THE NEW COMMAND ARGS DESCRIPTION]

Syntax: 
golangspell [TODO: PUT HERE THE NEW COMMAND SYNTAX]

Examples:
[TODO: PUT HERE THE NEW COMMAND EXAMPLES IF NEEDED]`,
ValidArgs: []string{"DomainEntity"},
			},
		},
	}
}