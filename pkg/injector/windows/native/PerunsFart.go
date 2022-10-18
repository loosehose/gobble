package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type PeRunsFart struct {
	Name        string
	Description string
	Debug       bool
}

func NewPeRunsFart() models.ObjectModel {
	return &PeRunsFart{
		Name:        "windows/native/local/PeRunsFart",
		Description: "This program executes shellcode in a process",
		Debug:       false,
	}
}

func (i *PeRunsFart) GetImports() []string {

	return []string{
	    `"fmt"`,
	    `"github.com/timwhitez/Doge-Gabh/pkg/Gabh"`,
	    `"syscall"`,
	    `"unsafe"`,
        `"golang.org/x/sys/windows"`,
        `"log"`,
	}
}

func (e *PeRunsFart) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/perunsfart/instanciation.go.tmpl", e)
}

func (e *PeRunsFart) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/perunsfart/functions.go.tmpl", e)
}
