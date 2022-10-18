package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type EarlyBird struct {
	Name        string
	Description string
	Debug       bool
}

func NewEarlyBird() models.ObjectModel {
	return &EarlyBird{
		Name:        "windows/native/local/EarlyBird",
		Description: "Use native windows api call CreateThread to inject into the current process.",
		Debug:       false,
	}
}

func (i *EarlyBird) GetImports() []string {

	return []string{
		`"syscall"`,
		`"unsafe"`,
		`"golang.org/x/sys/windows"`,
        `"fmt"`,
	    `"log"`,
	}
}

func (e *EarlyBird) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/EarlyBird/instanciation.go.tmpl", e)
}

func (e *EarlyBird) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/EarlyBird/functions.go.tmpl", e)
}
