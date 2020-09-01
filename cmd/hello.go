package cmd

import (
	"fmt"

	"github.com/andersonlira/golangspell-txtdb/usecase"
	"github.com/spf13/cobra"
)

func init() {
	RunCommandFunctions["golangspell-txtdb-hello"] = runHelloCommand
}

//Module name to initialize with 'Go Modules'
var Module string

//AppName used to define the application's directory and the default value to the config variable with the same name
var AppName string

func runHelloCommand(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println(`The command golangspell-txtdb-hello requires exactly one parameter: name
Args:
name: Your name (required) to be added to the Hello!. Example: Elvis"

Syntax: 
golangspell golangspell-txtdb-hello [name]`)
		return
	}

	err := usecase.SayHello(args)
	if err != nil {
		fmt.Printf("An error occurred while trying to execute the command. Message: %s\n", err.Error())
	}
}
