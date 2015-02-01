package command

import (
	"fmt"
	"io"
	"os/exec"
)

func ExecuteCommand(output io.Writer, cmd string) {
	command, _ := composedExecuteCommand(output, cmd)
	command.Run()
}

func basicExecuteCommand(output io.Writer, cmd string) (*exec.Cmd, error) {
	command := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s", cmd))
	command.Stdout = output
	command.Stderr = output

	return command, nil
}

func composedExecuteCommand(requestOutput io.Writer, cmd string) (*exec.Cmd, error) {
	command := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s", cmd))
	commandOutput, err := command.StdoutPipe()
	go supervisor(command, requestOutput, commandOutput)

	return command, err
}

func supervisor(command *exec.Cmd, output io.Writer, input io.Reader) {
	if _, err := io.Copy(output, input); err != nil {
		command.Process.Kill()
	}
}
