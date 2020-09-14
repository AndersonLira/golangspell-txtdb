package usecase

import (
	"fmt"
	"strconv"
	"strings"
	
	"github.com/andersonlira/goutils/io"
	"github.com/andersonlira/goutils/recode"
	"github.com/iancoleman/strcase"
)

//RendertxtdbsortTemplate renders the templates defined to the txtdbsort command with the proper variables
func RendertxtdbsortTemplate(args []string) error {
	domainEntity := args[0]
	field := strcase.ToCamel(args[1])
	domainSnake := strcase.ToSnake(domainEntity)

	operator := "<"

	if len(args) > 2 {
		if b, _ := strconv.ParseBool(args[2]); b {
			operator = ">"
		}
	}
	
	coder, err := recode.MakeCoder(fmt.Sprintf("./gateway/txtdb/%s_repository.go", domainSnake))
	if err != nil {
		return err
	}
	

	line1 := "\tsort.Slice(list, func(i, j int) bool {"
	line2 := fmt.Sprintf("\t\treturn list[i].%s %s list[j].%s",field,operator,field)
	line3 := "\t})"
	coder.AddAfterLine(fmt.Sprintf("func Get%sList() []domain.%s",domainEntity,domainEntity),"return list",line1,line2,line3)

	if !strings.Contains(coder.NewCodeContent(),`"sort"`) {
		coder.AddAfterLine("import (",")","\t\"sort\"")
	}

	io.WriteFile(fmt.Sprintf("./gateway/txtdb/%s_repository.go", domainSnake),coder.NewCodeContent())
	
	return nil
}
