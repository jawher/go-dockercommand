package dockercommand

import (
	"fmt"
	"testing"
)

func TestDockerInspect(t *testing.T) {
	docker := &Docker{}
	err, containerID := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"ls", "/"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err, inspect := docker.Inspect(containerID)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	fmt.Printf("%s\n", inspect)
}