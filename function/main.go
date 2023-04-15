/*
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/containers/podman/v4/pkg/bindings"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/bindings/images"
	"github.com/containers/podman/v4/pkg/specgen"
)

func main() {
	conn, err := bindings.NewConnection(context.Background(), "unix:///opt/homebrew/bin/podman")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = images.Pull(conn, "quay.io/libpod/alpine_nginx", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := specgen.NewSpecGenerator("quay.io/libpod/alpine_nginx", false)
	s.Name = "foobar"
	createResponse, err := containers.CreateWithSpec(conn, s, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Container created.")
	if err := containers.Start(conn, createResponse.ID, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Container started.")
}
/*