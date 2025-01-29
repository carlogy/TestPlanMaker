package main

import (
	"fmt"
	"os"
	"os/user"
	"text/template"

	TP "github.com/carlogy/TestPlanMaker/internal/TestPlanTemplate"
)

func BuildTemplate(tp *TP.TestPlanTemplate, templatePath string) {

	fmt.Println("Starting to template buildout...")

	tmpl, err := template.ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	saveFilePath, err := getSavePath(tp)
	if err != nil {
		fmt.Println("Error: unable to resolve save file path.")
	}

	f, err := os.Create(saveFilePath)
	if err != nil {
		fmt.Println(err)
	}

	err = tmpl.Execute(f, tp)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Succesfully saved file to path:\t", saveFilePath)
}

func getSavePath(tp *TP.TestPlanTemplate) (string, error) {

	if tp.SavePath == "" {
		user, err := user.Current()
		if err != nil {
			return "", err
		}

		savePath := fmt.Sprintf("%s/Desktop/%s.md", user.HomeDir, tp.UserStory)

		return savePath, nil
	}

	savePath := fmt.Sprintf("%s/%s.md", tp.SavePath, tp.UserStory)

	return savePath, nil

}
