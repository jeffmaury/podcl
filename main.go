package main

import (
	"context"
	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/images"
)

func main() {
	ctx, err := bindings.NewConnection(context.Background(), "unix://.//pipe/default-podman-machine")
	if err != nil {
		panic(err)
	}
	_, err = images.GetImage(ctx, "ng", nil)
	if err != nil {
		panic(err)
	}
}
