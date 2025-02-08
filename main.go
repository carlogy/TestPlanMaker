package main

import (
	"fmt"
	"log"
	"os"

	TPF "github.com/carlogy/TestPlanMaker/internal/TestPlanForm"
	TP "github.com/carlogy/TestPlanMaker/internal/TestPlanTemplate"
)

func main() {
	fmt.Println("Welcome to test plan maker!\nLet's build out a test plan template together.")

	tp := TP.NewTestPlanBuild()
	tpf := TPF.BuilForm(tp)

	err := tpf.Run()
	if err != nil {
		log.Fatal(err)
	}

	// add spinner func here ...

	// BuildTemplate(tp, "internal/Templates/testPlan.tmpl")
	//
	err = BuildTemplateFromString(tp)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if tp.SavePath == "" {
		fmt.Println("Succesfully saved file to Desktop dir")
	}
	fmt.Printf("Succesfully saved file to path:\t%s\n\n", tp.SavePath)
}
