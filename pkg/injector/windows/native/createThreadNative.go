package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type CreateThreadNative struct {
	Name        string
	Description string
	Debug       bool
}

func NewCreateThreadNative() models.ObjectModel {
	return &CreateThreadNative{
		Name:        "windows/native/local/CreateThreadNative",
		Description: "Use native windows api call CreateThread to inject into the current process.",
		Debug:       false,
	}
}

func (i *CreateThreadNative) GetImports() []string {

	return []string{
		`"syscall"`,
		`"unsafe"`,
		`"golang.org/x/sys/windows"`,
	}
}

func (e *CreateThreadNative) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/CreateThreadNative/instanciation.go.tmpl", e)
}

func (e *CreateThreadNative) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/CreateThreadNative/functions.go.tmpl", e)
}
