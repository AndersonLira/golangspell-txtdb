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
			"txtdbdomain": &domain.Command{
				Name:             "txtdbdomain",
				ShortDescription: "The txtdbdomain create a new domain structure",
				LongDescription: `Args:
DomainEntity: the entity name

Syntax: 
golangspell txtdbdomain DomainEntity

Examples:
golangspell txtdbdomain Group`,
				ValidArgs: []string{"DomainEntity"},
			},
			"txtdbsort": &domain.Command{
				Name:             "txtdbsort",
				ShortDescription: "The txtdbsort create a sort in repository for giving DomainEnitty",
				LongDescription: `Args:
DomainEntity: the entity Name
Field: the field that will receive sort logic
Desc [optional]: sort is desc mode if value = true. Default false
Syntax: 
golangspell txtdbsort DomainEntity Field [Desc]

Examples:
golangspell txtdbsort Group Name <true>``,
				ValidArgs: []string{"DomainEntity","Field"},
			},
		},
	}
}
