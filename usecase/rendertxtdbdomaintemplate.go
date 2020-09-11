package usecase

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andersonlira/goutils/recode"
	"github.com/andersonlira/goutils/io"
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
	err = renameFiles(domainEntity)

	if err != nil {
		return err
	}

	return addRoutes(domainEntity)
}

func renameFiles(domainEntity string) error {
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}

	sourcePath := fmt.Sprintf("%s%sdomain%smodel_domain.go",currentPath , toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory := filepath.Dir(sourcePath)
	destinationPath := fmt.Sprintf("%s%smodel_%s.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))


	sourcePath = fmt.Sprintf("%s%scontroller%smodel_controller.go",currentPath , toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_controller.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		return err
	}

	sourcePath = fmt.Sprintf("%s%sgateway/txtdb/%smodel_repository.go",currentPath , toolconfig.PlatformSeparator, toolconfig.PlatformSeparator)
	directory = filepath.Dir(sourcePath)
	destinationPath = fmt.Sprintf("%s%s%s_repository.go", directory, toolconfig.PlatformSeparator, strcase.ToSnake(domainEntity))

	return os.Rename(sourcePath, destinationPath)
}

func addRoutes(domainEntity string ) error{
	coder, err := recode.MakeCoder("./controller/router.go")
	if err != nil {
		return err
	}
	routeGetAll := fmt.Sprintf("\tg.GET(\"/%s\", Get%sList)",strcase.ToLowerCamel(domainEntity),domainEntity)
	routeGetByID := fmt.Sprintf("\tg.GET(\"/%s/:id\", Get%sByID)",strcase.ToLowerCamel(domainEntity),domainEntity)
	coder.AddAfterLine("func MapRoutes(e *echo.Echo)","g.GET(",routeGetAll, routeGetByID)
	io.WriteFile("./controller/router.go",coder.NewCodeContent())
	return nil
}

