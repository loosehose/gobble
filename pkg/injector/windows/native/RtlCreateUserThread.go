package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type RtlCreateUserThread struct {
	Name        string
	Description string
	Debug       bool
}

func NewRtlCreateUserThread() models.ObjectModel {
	return &RtlCreateUserThread{
		Name:        "windows/native/local/RtlCreateUserThread",
		Description: "This program executes shellcode in a process",
		Debug:       false,
	}
}

func (i *RtlCreateUserThread) GetImports() []string {

	return []string{
	    `"fmt"`,
	    `"golang.org/x/sys/windows"`,
	    `"syscall"`,
	    `"unsafe"`,
        `"log"`,
        `"os"`,
	    `ps "github.com/mitchellh/go-ps"`,
	}
}

func (e *RtlCreateUserThread) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/RtlCreateUserThread/instanciation.go.tmpl", e)
}

func (e *RtlCreateUserThread) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/RtlCreateUserThread/functions.go.tmpl", e)
}
