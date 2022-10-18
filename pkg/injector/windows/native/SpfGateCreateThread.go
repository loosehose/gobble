package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type SpfGateCreateThread struct {
	Name        string
	Description string
	Debug       bool
}

func NewSpfGateCreateThread() models.ObjectModel {
	return &SpfGateCreateThread{
		Name:        "windows/native/local/SpfGateCreateThread",
		Description: "This program executes shellcode in a process",
		Debug:       false,
	}
}

func (i *SpfGateCreateThread) GetImports() []string {

	return []string{
	    `"crypto/sha1"`,
	    `"crypto/sha256"`,
	    `"encoding/hex"`,
	    `"fmt"`,
	    `"syscall"`,
	    `"unsafe"`,
	    `"github.com/timwhitez/Doge-Gabh/pkg/Gabh"`,
	    `"golang.org/x/sys/windows"`,
	}
}

func (e *SpfGateCreateThread) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/SpfGateCreateThread/instanciation.go.tmpl", e)
}

func (e *SpfGateCreateThread) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/SpfGateCreateThread/functions.go.tmpl", e)
}
