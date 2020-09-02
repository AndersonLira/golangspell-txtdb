package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andersonlira/golangspell-txtdb/appcontext"
	"github.com/andersonlira/golangspell-txtdb/domain"
	tooldomain "github.com/golangspell/golangspell/domain"
	toolconfig "github.com/golangspell/golangspell/config"
	"github.com/iancoleman/strcase"
)

//RendertxtdbdomainTemplate renders the templates defined to the txtdbdomain command with the proper variables
func RendertxtdbdomainTemplate(args []string) error {
	spell := appcontext.Current.Get(appcontext.Spell).(tooldomain.Spell)
	renderer := domain.GetRenderer()
	domainEntity := args[0]
	globalVariables := map[string]interface{}{
		"DomainEntity": domainEntity,
	}

	err := renderer.RenderTemplate(spell, "txtdbdomain", globalVariables, nil)
	if err != nil {
		return err
	}
	return renameFile(domainEntity)
}

func renameFile(domainEntity string) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	sourcePath := fmt.Sprintf("%s%sdomain%smodel_domain.go",currentPath , toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%smodel_%s.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))
	
	return os.Rename(sourcePath, destinationPath)
}

