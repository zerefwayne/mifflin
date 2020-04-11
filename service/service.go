package service

import (
	"fmt"
	"github.com/zerefwayne/mifflin/utils"
	"log"
)

var Services map[string][]string


func init() {

	Services = make(map[string][]string)

	Services["docker"] = append(Services["docker"], "docker")
	Services["docker"] = append(Services["docker"], "docker.service")

	Services["postgresql"] = append(Services["postgresql"], "postgresql")
	Services["postgresql"] = append(Services["postgresql"], "postgresql.service")

}


func StatusUtil(subService string) {

	if out, err := utils.CommandExec(fmt.Sprintf("systemctl show -p ActiveState %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Println(string(out[:]))
	}

}

func StartUtil(subService string) {

	if _, err := utils.CommandExec(fmt.Sprintf("sudo systemctl start %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s | started successfully.\n\n", subService)
	}

}

func StopUtil(subService string) {

	if _, err := utils.CommandExec(fmt.Sprintf("sudo systemctl stop %s", subService)); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%s | stopped successfully.\n\n", subService)
	}


}

func ManageService(service string, cmd string) {

	subServices := Services[service]

	switch cmd{
		case "status":
			for _, subService := range subServices {
				StatusUtil(subService)
			}
			break
		case "start":
			for _, subService := range subServices {
				StartUtil(subService)
			}
			break
		case "stop":
			for _, subService := range subServices {
				StopUtil(subService)
			}
			break
	}

}
