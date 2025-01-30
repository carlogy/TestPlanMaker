package main

import (
	"fmt"
	"os"
	"os/user"
	"text/template"

	TP "github.com/carlogy/TestPlanMaker/internal/TestPlanTemplate"
)

func BuildTemplateFromString(tp *TP.TestPlanTemplate) {

	if tp == nil {

		fmt.Println("Empty instance, please ensure you are entering valid tokens.")
		os.Exit(1)
	}

	fmt.Println("Starting to write template ")

	saveFilePath, err := getSavePath(tp)
	if err != nil {
		fmt.Println("Error: unable to resolve save file path.")
	}

	f, err := os.Create(saveFilePath)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	num, err := fmt.Fprintf(f, tp.MDString())
	if err != nil {
		fmt.Errorf("Error: %w\n experienced while writing file\n", err)
	}

	fmt.Printf("Succesfully saved file to path:\t%s\nTotal bytes written:\t%d\n", saveFilePath, num)

}

func BuildTemplate(tp *TP.TestPlanTemplate, templatePath string) {

	if tp == nil {

		fmt.Println("Empty instance, please ensure you are entering valid tokens.")
		os.Exit(1)
	}

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
