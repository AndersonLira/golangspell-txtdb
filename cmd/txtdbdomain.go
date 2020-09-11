package cmd

import (
	"fmt"

	"github.com/andersonlira/golangspell-txtdb/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["txtdbdomain"] = runtxtdbdomain
}

func runtxtdbdomain(cmd *cobra.Command, args []string) {
	// Example on how to deal when the expected arguments were not provided
	if len(args) != 1 {
		fmt.Println(`The command txtdbdomain requires 1 argument
		Args:
		DomainEntity: the domain that will be created
		
		Syntax: 
		golangspell txtdbdomain DomainEntity
		
		Examples:
		golagnspell txtdbdomain Group`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered 
	err := usecase.RendertxtdbdomainTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	//TODO: Create your additional logic here
	fmt.Println("txtdbdomain executed!")
}
