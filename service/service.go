package service

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/zerefwayne/mifflin/utils"
	"log"
)

func StatusUtil(subService string) {

	if out, err := utils.CommandExec(fmt.Sprintf("systemctl show -p ActiveState %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s | %s", subService, string(out[:]))
	}

}

func StartUtil(subService string) {

	if _, err := utils.CommandExec(fmt.Sprintf("sudo systemctl start %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s | started successfully.\n", subService)
	}

}

func StopUtil(subService string) {

	if _, err := utils.CommandExec(fmt.Sprintf("sudo systemctl stop %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s | stopped successfully.\n", subService)
	}

}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func ManageServices(cmd string, service string) {

	subServices := utils.FetchSubServices(service)
	options := append([]string{"all"}, subServices...)

	var selectedSubServices []string

	switch cmd {
	case "status":
		for _, subService := range subServices {
			StatusUtil(subService)
		}
		fmt.Println()
		break
	case "start":

		prompt := &survey.MultiSelect{
			Message: "Select the services to start",
			Options: options,
		}

		_ = survey.AskOne(prompt, &selectedSubServices)

		_, selectedAll := Find(selectedSubServices, "all")

		if selectedAll {

			for _, subService := range subServices {
				StartUtil(subService)
			}

		} else {

			for _, subService := range selectedSubServices {
				StartUtil(subService)
			}

		}

		break

	case "stop":
		prompt := &survey.MultiSelect{
			Message: "Select the services to stop",
			Options: options,
		}

		_ = survey.AskOne(prompt, &selectedSubServices)

		_, selectedAll := Find(selectedSubServices, "all")

		if selectedAll {

			for _, subService := range subServices {
				StopUtil(subService)
			}

		} else {

			for _, subService := range selectedSubServices {
				StopUtil(subService)
			}

		}

		break
	}

}
