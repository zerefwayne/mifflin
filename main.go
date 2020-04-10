package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func CommandGenerator(cmd string) (string, []string) {

	command := strings.Fields(cmd)

	var mainCmd string
	var args []string

	mainCmd = command[0]

	if len(command) > 1 {
		args = command[1:]
	}

	return mainCmd, args

}

func CommandExec(cmd string) ([]byte, error){

	var response []byte

	cmd, args := CommandGenerator(cmd)

	if out, err := exec.Command(cmd, args...).Output(); err != nil {
		return response, err
	} else {
		return out, nil
	}
}

func CommandRunner(cmd string) {

	if resp, err := CommandExec(cmd); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(resp[:]))
	}

}

func main() {

	CommandRunner("pwd")
	CommandRunner("ls -a")
	CommandRunner("ls")
	CommandRunner("git status")
}
