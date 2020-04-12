package boil

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os/exec"
)

type boilerplateConfig struct {
	Name string
	Description string
	Repository string
	Dependencies []string
	OptionalDependencies []string `yaml:"optional-dependencies"`
	UpdatedOn string `yaml:"updated-on"`
}

type yamlConfig struct {
	Boilerplates map[string]boilerplateConfig
}

func LoadBoilerplates() {

	if yamlFile, err := ioutil.ReadFile("/home/zerefwayne/github/personal/mifflin/boil/boilerplates.yaml"); err != nil {
		log.Fatal(err)
	} else {
		var data yamlConfig

		if err := yaml.Unmarshal(yamlFile, &data); err != nil {
			log.Fatal(err)
		} else {

			var boilerplates []string

			for k, _ := range data.Boilerplates {
				boilerplates = append(boilerplates, k)
			}

			choice := ""

			prompt := &survey.Select{
				Message: "Choose a boilerplate:",
				Options: boilerplates,
			}

			if err := survey.AskOne(prompt, &choice); err != nil {
				log.Fatal(err)
			} else {

				fmt.Println()
				fmt.Println("Name:", data.Boilerplates[choice].Name)
				fmt.Println("Description:", data.Boilerplates[choice].Description)
				fmt.Println("Repository:", data.Boilerplates[choice].Repository)
				fmt.Println()

				currentDirectory := false
				prompt := &survey.Confirm{
					Message: "Would you like to download boilerplate in current directory?",
				}

				if err := survey.AskOne(prompt, &currentDirectory); err != nil {
					log.Fatal(err)
				} else {

					if currentDirectory {
						fmt.Println()
						fmt.Println("Installing in current directory...")
						fmt.Println()

						if out, err := exec.Command("git", "clone", data.Boilerplates[choice].Repository).Output(); err != nil {
							log.Fatal(err)
						} else {

							fmt.Println("\u2714 Successfully cloned!")

							fmt.Printf("\n%s\n", string(out[:]))
						}

					} else {
						fmt.Println("Exiting.")
					}



				}

			}

		}
	}
}