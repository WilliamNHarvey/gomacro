package williamnharvey

import (
	"context"
	"fmt"

	"github.com/WilliamNHarvey/gomacro/_example/williamnharvey/williamnharveybackend"
	"github.com/WilliamNHarvey/gomacro/_example/williamnharvey/williamnharveyfrontend"
)

func main() {
	williamnharveyfrontend.InteractableBackend = williamnharveybackend.InteractableBackend{}

	newStruct := williamnharveyfrontend.InterfaceStruct{}
	williamnharveyfrontend.DoSomethingWithInterfaceStruct(context.Background(), &newStruct)

	fmt.Println(newStruct.Name)
}
