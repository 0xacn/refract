package function

import ( 
  "fmt"
  "context"

  "github.com/containers/podman/v4/pkg/bindings"
  "github.com/containers/podman/v4/pkg/bindings/containers"
  "github.com/containers/podman/v4/pkg/bindings/images"
  "github.com/containers/podman/v4/pkg/specgen"
)

func createFunction(conn *bindings.Connection, imageName string, functionName string) error {
  ctx := context.Background()
  _, err := images.Pull(ctx, imageName, nil)

  if err != nil {
    return err
  }

  s := specgen.NewSpecGenerator(imageName, false)
  s.Name = functionName

  createOptions := &containers.CreateOptions{}
  createResponse, err := containers.CreateWithSpec(ctx, s, createOptions)
  if err != nil {
    return err
  }

  startOptions := &containers.StartOptions{}
  if err := containers.Start(ctx, createResponse.ID, startOptions); err != nil {
    return err
  } 

  fmt.Printf("Container %s created and started.\n", functionName)
  return nil
}

