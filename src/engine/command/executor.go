package command

import (
	"fmt"
	"io"
	"os/exec"
)

func ExecuteCommand(output io.Writer, cmd string) {
	command := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s", cmd))
	command.Stdout = output
	command.Stderr = output
	command.Run()
}
