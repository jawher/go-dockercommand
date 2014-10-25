package dockercommand

import "testing"

func TestDockerRm(t *testing.T) {
	docker := &Docker{}
	containerID, err := docker.Run(&RunOptions{
		Image: "ubuntu",
		Cmd:   []string{"ls", "/"},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	err = docker.Rm(&RmOptions{
		Container: []string{containerID},
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
}
