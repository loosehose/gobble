package bananaphone

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type NtCreateThreadEx struct {
	Name        string
	Description string
	Debug       bool
}

func NewNtCreateThreadEx() models.ObjectModel {
	return &NtCreateThreadEx{
		Name:        "windows/bananaphone/NtCreateThreadEx",
		Description: "Use native windows api call NtCreateThreadEx to inject in the current process. Call is performed using bananaphone from @C-Sto.",
		Debug:       false,
	}
}

func (i *NtCreateThreadEx) GetImports() []string {

	return []string{
		`"syscall"`,
		`"unsafe"`,
		`bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"`,
		`"golang.org/x/sys/windows"`,
        `ps "go-ps"`,
	}
}

func (e *NtCreateThreadEx) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/NtCreateThreadEx/instanciation.go.tmpl", e)
}

func (e *NtCreateThreadEx) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/NtCreateThreadEx/functions.go.tmpl", e)
}
