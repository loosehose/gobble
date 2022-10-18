package encoder

import (
	"embed"

	"github.com/loosehose/gobble/pkg/common"
	"github.com/loosehose/gobble/pkg/models"
)

type ReverseOrderEncoder struct {
	Name        string
	Description string
	Debug       bool
}

func NewReverseOrderEncoder() models.ObjectModel {
	return &ReverseOrderEncoder{
		Name:        "reverse-order",
		Description: "Reverse the order of the shellcode byte array.",
		Debug:       false,
	}
}

func (e *ReverseOrderEncoder) Encode(shellcode []byte) ([]byte, error) {
    for i, j := 0, len(shellcode)-1; i < j; i, j = i+1, j-1 {
        shellcode[i], shellcode[j] = shellcode[j], shellcode[i]
    }

	return shellcode, nil
}

func (e *ReverseOrderEncoder) GetImports() []string {

	return []string{}
}

func (e *ReverseOrderEncoder) RenderInstanciationCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/encoders/reverse-order/instanciation.go.tmpl", e)
}

func (e *ReverseOrderEncoder) RenderFunctionCode(data embed.FS) (string, error) {

	return common.CommonRendering(data, "templates/encoders/reverse-order/functions.go.tmpl", e)
}
