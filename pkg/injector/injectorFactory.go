package injector

import (
	"fmt"

	"github.com/loosehose/gobble/pkg/injector/windows/bananaphone"
	"github.com/loosehose/gobble/pkg/injector/windows/native"
	"github.com/loosehose/gobble/pkg/models"
)

func GetInjector(injectorType string) (models.ObjectModel, error) {
	if injectorType == "windows/native/local/go-shellcode-syscall" {
		return native.NewSyscallGoShellcode(), nil
	}

	if injectorType == "windows/native/local/hellsgate" {
		return native.NewHellsGate(), nil
	}

	if injectorType == "windows/native/local/spfgatecreatethread" {
		return native.NewSpfGateCreateThread(), nil
	}


	if injectorType == "windows/native/local/perunsfart" {
		return native.NewPeRunsFart(), nil
	}

	if injectorType == "windows/bananaphone/local/unhooker" {
		return bananaphone.NewUnhooker(), nil
	}

	if injectorType == "windows/bananaphone/local/ntqueueapcthreadex-local" {
		return bananaphone.NewNtQueueApcThreadExLocal(), nil
	}

	if injectorType == "windows/bananaphone/local/go-shellcode-syscall" {
		return bananaphone.NewSyscallGoShellcode(), nil
	}

	if injectorType == "windows/bananaphone/local/ninjauuid" {
		return bananaphone.NewNinjaUUID(), nil
	}
	if injectorType == "windows/bananaphone/local/ntcreatethreadex" {
		return bananaphone.NewNtCreateThreadEx(), nil
	}

	return nil, fmt.Errorf("Wrong injector type passed: injector %s is unknown", injectorType)
}
