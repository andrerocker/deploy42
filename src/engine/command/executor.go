package command

import (
	"io"
	"os/exec"
)

func ExecuteCommand(input io.Reader, output io.Writer, cmd string) {
	command, _ := composedExecuteCommand(input, output, cmd)
	command.Run()
}

func basicExecuteCommand(input io.Reader, output io.Writer, cmd string) (*exec.Cmd, error) {
	command := exec.Command("/bin/bash", "-c", cmd)
	command.Stdout = output
	command.Stderr = output

	return command, nil
}

func composedExecuteCommand(input io.Reader, requestOutput io.Writer, cmd string) (*exec.Cmd, error) {
	command := exec.Command("/bin/bash", "-c", cmd)
	command.Stdin = input

	commandOutput, err := command.StdoutPipe()
	go supervisor(command, requestOutput, commandOutput)

	return command, err
}

func supervisor(command *exec.Cmd, output io.Writer, input io.Reader) {
	if _, err := io.Copy(output, input); err != nil {
		command.Process.Kill()
	}
}
