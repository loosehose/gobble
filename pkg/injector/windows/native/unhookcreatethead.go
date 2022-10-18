package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type UnhookCreateThread struct {
	Name        string
	Description string
	Debug       bool
}

func NewUnhookCreateThread() models.ObjectModel {
	return &UnhookCreateThread{
		Name:        "windows/native/local/UnhookCreateThread",
		Description: "This program creates a new thread while unhooking from NTDLL using RecycleGate.",
		Debug:       false,
	}
}

func (i *UnhookCreateThread) GetImports() []string {

	return []string{
	    `"crypto/sha1"`,
	    `"crypto/sha256"`,
	    `"encoding/hex"`,
	    `"fmt"`,
	    `"github.com/timwhitez/Doge-Gabh/pkg/Gabh"`,
	    `"syscall"`,
	    `"unsafe"`,
	}
}

func (e *UnhookCreateThread) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/UnhookCreateThread/instanciation.go.tmpl", e)
}

func (e *UnhookCreateThread) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/UnhookCreateThread/functions.go.tmpl", e)
}
