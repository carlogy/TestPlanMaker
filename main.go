package main

import (
	"fmt"
	"log"

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

	BuildTemplate(tp, "internal/Templates/testPlan.tmpl")

}
