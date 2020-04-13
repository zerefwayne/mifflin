package utils

import (
	"bufio"
	"fmt"
	"log"
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

func CommandExec(cmd string) ([]byte, error) {

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

func FetchSubServices(service string) []string {

	cmdString := fmt.Sprintf("systemctl list-unit-files | grep %s", service)

	cmd := exec.Command("bash", "-c", cmdString)

	var subServices []string

	if stdout, err := cmd.StdoutPipe(); err != nil {
		log.Fatal(err)
	} else {

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		} else {

			defer cmd.Wait()

			buff := bufio.NewScanner(stdout)

			var returnText []string

			for buff.Scan() {
				returnText = append(returnText, buff.Text())
			}

			for _, subService := range returnText {
				subServices = append(subServices, strings.Fields(subService)[0])
			}

		}

	}

	return subServices

}
