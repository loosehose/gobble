package evasion

import (
	"embed"
	"strings"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type HideConsoleEvasion struct {
	Name        string
	Description string
	Debug       bool
	Show        string
	ShowBool 	bool
}

func NewHideConsoleEvasion() models.ObjectModel {
	return &HideConsoleEvasion{
		Name:        "hideconsole",
		Debug: false,
		Description: `Prevent windows console to be displayed
  Argument(s):
    Show: Show the console or not. Default is "false".`,
		Show: "false",
	}
}

func (e *HideConsoleEvasion) GetImports() []string {

	return []string{
		`"syscall"`,
	}
}

func (e *HideConsoleEvasion) RenderInstanciationCode(data embed.FS) (string, error) {
	if strings.ToLower(e.Show) == "true" {
		e.ShowBool = true
	}

	return common.CommonRendering(data, "templates/evasions/hideconsole/instanciation.go.tmpl", e)
}

func (e *HideConsoleEvasion) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/evasions/hideconsole/functions.go.tmpl", e)
}
