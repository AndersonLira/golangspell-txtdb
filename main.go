package main

import (
	"fmt"
	
	"github.com/andersonlira/golangspell-txtdb/cmd"
	_ "github.com/andersonlira/golangspell-txtdb/config"
	_ "github.com/andersonlira/golangspell-txtdb/gateway/template"
	_ "github.com/andersonlira/golangspell-txtdb/gateway/customlog"

)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Printf("An error occurred while executing the command. Message: %s\n", err.Error())
	}
}
