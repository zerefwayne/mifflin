package warehouse

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/kyokomi/emoji"
	"gopkg.in/yaml.v2"
	"log"
	"os/exec"
)

type boilerplateConfig struct {
	Name                 string
	Description          string
	Repository           string
	Dependencies         []string
	OptionalDependencies []string `yaml:"optional-dependencies"`
	UpdatedOn            string   `yaml:"updated-on"`
}

type yamlConfig struct {
	Boilerplates map[string]boilerplateConfig
}

var config = `
boilerplates:
  go-psql-rest:
      name: "Go PostgreSQL Rest API"
      description: "A Rest API Written in Golang (Mux) and PostgreSQL used as database."
      repository: "https://github.com/mifflin-org/bp-go-psql-rest.git"
  go-psql-rest-jwt:
      name: "Go PostgreSQL Rest API secured via Local Auth and JWT"
      description: "A Rest API Written in Golang (Mux) and PostgreSQL used as database, secured by Local Authentication and JWT"
      repository: "https://github.com/mifflin-org/bp-go-psql-rest-jwt.git"
`

func fetchAvailableBoilerplates() (yamlConfig, error) {

	var data yamlConfig

	if err := yaml.Unmarshal([]byte(config), &data); err != nil {
		return data, err
	}

	return data, nil

}

func inputBoilerplateChoice(data yamlConfig) (boilerplateConfig, error) {

	var boilerplates []string
	var selectedBoilerplate boilerplateConfig

	for k, _ := range data.Boilerplates {
		boilerplates = append(boilerplates, k)
	}

	choice := ""

	prompt := &survey.Select{
		Message: "Choose a boilerplate:",
		Options: boilerplates,
	}

	if err := survey.AskOne(prompt, &choice); err != nil {
		return selectedBoilerplate, err
	} else {
		return data.Boilerplates[choice], nil
	}

}

func outputBoilerplateData(boilerplate boilerplateConfig) {
	fmt.Println()
	fmt.Println("Name:", boilerplate.Name)
	fmt.Println("Description:", boilerplate.Description)
	fmt.Println("Repository:", boilerplate.Repository)
	fmt.Println()
}

func inputProjectName() (string, error) {

	name := ""

	prompt := &survey.Input{
		Message: "Enter folder name",
	}

	if err := survey.AskOne(prompt, &name); err != nil {
		return "", err
	}

	return name, nil

}

func fetchBoilerplate(boilerplate boilerplateConfig) {

	var projectName string

	if data, err := inputProjectName(); err != nil {
		log.Fatal(err)
	} else {
		projectName = data
	}

	fmt.Printf("Installing in ./%s\n", projectName)

	cmd := fmt.Sprintf("git clone %s %s", boilerplate.Repository, projectName)

	if _, err := exec.Command("bash", "-c", cmd).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s Successfully cloned!\n", emoji.Sprint(":dizzy:"))
	}

}

func InstallBoilerplate() {

	var availableBoilerplates yamlConfig
	var selectedBoilerplate boilerplateConfig

	if data, err := fetchAvailableBoilerplates(); err != nil {
		log.Fatal(err)
	} else {
		availableBoilerplates = data
	}

	if data, err := inputBoilerplateChoice(availableBoilerplates); err != nil {
		log.Fatal(err)
	} else {
		selectedBoilerplate = data
	}

	outputBoilerplateData(selectedBoilerplate)

	fetchBoilerplate(selectedBoilerplate)

}
