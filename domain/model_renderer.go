package domain

import (
	"strings"
	"os"

	"github.com/golangspell/golangspell/domain"
	"github.com/andersonlira/golangspell-txtdb/appcontext"
	toolconfig "github.com/golangspell/golangspell/config"
	tooldomain "github.com/golangspell/golangspell/domain"
)

//Renderer defines the features delivered by the Code Template Renderer
type Renderer interface {
	//RenderFile renders a template file
	RenderFile(sourcePath string, info os.FileInfo) error

	//RenderPath renders an object (file os directory) in the templates directory
	RenderPath(sourcePath string, info os.FileInfo, err error) error

	//BackupExistingCode make a copy of the changed file
	BackupExistingCode(sourcePath string) error

	//RenderString processing the provided template source file, using the provided variables
	RenderString(spell domain.Spell, commandName string, stringTemplateFileName string, variables map[string]interface{}) (string, error)

	//RenderTemplate renders all templates in the template directory providing the respective variables
	//commandName: specifies the name of the command for which the template will be rendered
	//globalVariables: defines the list of variables (value) which should be provided for rendering all files
	//specificVariables: defines the list of variables (value) which should be provided for rendering specific file names (key)
	RenderTemplate(spell tooldomain.Spell, commandName string, globalVariables map[string]interface{}, specificVariables map[string]map[string]interface{}) error
}

type GenericRenderer struct {
	origin Renderer	
}

//RenderFile renders a template file
func (r GenericRenderer) RenderFile(sourcePath string, info os.FileInfo) error{
	return r.origin.RenderFile(sourcePath,info)
}

//RenderPath renders an object (file os directory) in the templates directory
func (r GenericRenderer) RenderPath(sourcePath string, info os.FileInfo, err error) error {
	return r.origin.RenderPath(sourcePath,info,err)
}

//BackupExistingCode make a copy of the changed file
func (r GenericRenderer) BackupExistingCode(sourcePath string) error {
	return r.origin.BackupExistingCode(sourcePath)
}

//RenderString processing the provided template source file, using the provided variables
func (r GenericRenderer) RenderString(spell domain.Spell, commandName string, stringTemplateFileName string, variables map[string]interface{}) (string, error) {
	return r.origin.RenderString(spell,commandName,stringTemplateFileName,variables)
}

//RenderTemplate renders all templates in the template directory providing the respective variables
func (r GenericRenderer) RenderTemplate(spell tooldomain.Spell, commandName string, globalVariables map[string]interface{}, specificVariables map[string]map[string]interface{}) error {
	currentPath, err := os.Getwd()
	if err == nil && globalVariables != nil {
		globalVariables["ModuleName"]  = strings.ReplaceAll(toolconfig.GetModuleName(currentPath),"\r","") //prevent Windows break line bug
	}
	return r.origin.RenderTemplate(spell,commandName,globalVariables,specificVariables)
}



//GetRenderer returns the current component registered to provide the code rendering features
func GetRenderer() Renderer {
	origin := appcontext.Current.Get(appcontext.Renderer)

	renderer := GenericRenderer{}
	renderer.origin = origin.(Renderer)
	return renderer
}
