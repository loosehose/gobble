package native

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type HellsGate struct {
	Name        string
	Description string
	Debug       bool
}

func NewHellsGate() models.ObjectModel {
	return &HellsGate{
		Name:        "windows/native/local/HellsGate",
		Description: "Use native windows api call CreateThread to inject into the current process.",
		Debug:       false,
	}
}

func (i *HellsGate) GetImports() []string {

	return []string{
		`"syscall"`,
		`"unsafe"`,
        `"fmt"`,
        `"crypto/sha1"`,
	    `"crypto/sha256"`,
	    `"encoding/hex"`,
	    `gabh "github.com/timwhitez/Doge-Gabh/pkg/Gabh"`,
	}
}

func (e *HellsGate) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/HellsGate/instanciation.go.tmpl", e)
}

func (e *HellsGate) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/injector/windows/native/local/HellsGate/functions.go.tmpl", e)
}
