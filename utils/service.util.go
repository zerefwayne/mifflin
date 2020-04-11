package utils

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

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
