package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type InjectProcess struct {
	Name        string
	Description string
	Debug       bool
}

func NewInjectProcess() models.ObjectModel {
	return &InjectProcess{
		Name:        "windows/native/local/InjectProcess",
		Description: "This program executes shellcode in a process",
		Debug:       false,
	}
}

func (i *InjectProcess) GetImports() []string {

	return []string{
	    `"fmt"`,
	    `"golang.org/x/sys/windows"`,
	    `"syscall"`,
	    `"unsafe"`,
	    `ps "github.com/mitchellh/go-ps"`,
	}
}

func (e *InjectProcess) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/InjectProcess/instanciation.go.tmpl", e)
}

func (e *InjectProcess) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/InjectProcess/functions.go.tmpl", e)
}
