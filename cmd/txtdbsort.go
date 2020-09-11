package cmd

import (
	"fmt"

	"github.com/andersonlira/golangspell-txtdb/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["txtdbsort"] = runtxtdbsort
}

func runtxtdbsort(cmd *cobra.Command, args []string) {
	// Example on how to deal when the expected arguments were not provided
	if len(args) != 2 {
		fmt.Println(`The command txtdbsort requires 2 arguments
		Args:
		DomainEntity: the domain that will receive sort feature
		Field: the field that will be used to sort. 
		Syntax: 
		golangspell txtdbsort DomainEntity Field
		
		Examples:
		golangspell txtdbsort Group Name`)
		return
	}

	//Here your template, hosted on the folder "templates" is rendered 
	err := usecase.RendertxtdbsortTemplate(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to render the template. Error: %s\n", err.Error())
		return
	}
	//TODO: Create your additional logic here
	fmt.Println("txtdbsort executed!")
}
