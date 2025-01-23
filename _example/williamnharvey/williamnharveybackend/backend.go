package williamnharveybackend

import (
	"context"

	"github.com/WilliamNHarvey/gomacro/_example/williamnharvey/williamnharveyfrontend"
)

type InteractableBackend struct{}

func (b InteractableBackend) DoSomethingWithInterfaceStruct(ctx context.Context, interfaceStruct *williamnharveyfrontend.InterfaceStruct) (*williamnharveyfrontend.InterfaceStruct, error) {
	interfaceStruct.Name = "Hello, World!"

	return nil
}
