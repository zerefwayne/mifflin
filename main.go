package main

import (
	"github.com/zerefwayne/mifflin/service"
)

func main() {

	service.ManageService("postgresql", "status")
}
