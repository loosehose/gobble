package bananaphone

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type Unhooker struct {
	Name        string
	Description string
	Debug       bool
}

func NewUnhooker() models.ObjectModel {
	return &Unhooker{
		Name:        "windows/bananaphone/local/unhooker",
		Description: "Unhook from NTDLL.dll using Hell's Gate method. Call is performed using bananaphone from @C-Sto.",
		Debug:       false,
	}
}

func (i *Unhooker) GetImports() []string {

	return []string{
        `"fmt"`,
		`"syscall"`,
		`"unsafe"`,
		`bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"`,
		`"golang.org/x/sys/windows"`,
	}
}

func (e *Unhooker) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/unhooker/instanciation.go.tmpl", e)
}

func (e *Unhooker) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/unhooker/functions.go.tmpl", e)
}
