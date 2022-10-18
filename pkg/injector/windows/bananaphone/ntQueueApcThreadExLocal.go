package bananaphone

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type NtQueueApcThreadExLocal struct {
	Name        string
	Description string
	Debug       bool
}

func NewNtQueueApcThreadExLocal() models.ObjectModel {
	return &NtQueueApcThreadExLocal{
		Name:        "windows/bananaphone/local/NtQueueApcThreadEx-Local",
		Description: "Use native windows api call NtQueueApcThreadExt to inject in the current process. Call is performed using bananaphone from @C-Sto.",
		Debug:       false,
	}
}

func (i *NtQueueApcThreadExLocal) GetImports() []string {

	return []string{
		`"syscall"`,
		`"unsafe"`,
		`bananaphone "github.com/C-Sto/BananaPhone/pkg/BananaPhone"`,
		`"golang.org/x/sys/windows"`,
	}
}

func (e *NtQueueApcThreadExLocal) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/NtQueueApcThreadEx-Local/instanciation.go.tmpl", e)
}

func (e *NtQueueApcThreadExLocal) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/bananaphone/local/NtQueueApcThreadEx-Local/functions.go.tmpl", e)
}
