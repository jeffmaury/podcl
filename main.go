package main

import (
	"context"
	"fmt"
	"github.com/containers/common/pkg/config"
	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/images"
)

func main() {
	conf, err := config.ReadCustomConfig()
	if err != nil {
		panic(err)
	}
	uri := ""
	identity := ""
	if conf.Engine.ActiveService != "" {
		uri = conf.Engine.ServiceDestinations[conf.Engine.ActiveService].URI
		identity = conf.Engine.ServiceDestinations[conf.Engine.ActiveService].Identity
	}
	fmt.Println("URI:", uri, " identity:", identity)
	ctx, err := bindings.NewConnectionWithIdentity(context.Background(), uri, identity, true)
	if err != nil {
		panic(err)
	}
	_, err = images.GetImage(ctx, "ng", nil)
	if err != nil {
		panic(err)
	}
}
