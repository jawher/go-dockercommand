package dockercommand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDockerInspect(t *testing.T) {
	docker, err := NewDocker("")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	containerID, err := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"ls", "/"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	inspect, err := docker.Inspect(containerID)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	assert.NotEmpty(t, inspect)
}
