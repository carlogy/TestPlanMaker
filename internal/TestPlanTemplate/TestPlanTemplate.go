package TestPlanTemplate

import (
	"fmt"
)

type TestPlanTemplate struct {
	Name string
	TicketDetails
	RequiredValidation string
	Description        string
	SavePath           string
	WriteToFile        bool
}

type TicketDetails struct {
	UserStory           string
	UserStoryLink       string
	UserStoryComponents string
	AutomationNeeded    bool
	PM                  string
	Dev                 string
	QA                  string
}

// Factory to test stringer function. For actually functionality will utilize the NewTestPlanBuild factory function to instanstiate a struct and return it's pointer to update with form inputs.

// func NewTestPlan(name, ticket, userStoryLink, validation, description, savePath string, automation bool) *TestPlanTemplate {
// 	tp := &TestPlanTemplate{
// 		Name:               name,
// 		TicketDetails:          ticket,
// 		UserStoryLink:      userStoryLink,
// 		RequiredValidation: validation,
// 		Description:        description,
// 		AutomationNeeded:   automation,
// 		SavePath:           savePath,
// 	}
// 	return tp
// }

func NewTestPlanBuild() *TestPlanTemplate {
	return &TestPlanTemplate{}
}

func (tp TestPlanTemplate) String() string {
	return fmt.Sprintf("TestPlan:\nName:\t%s\nUser Story:\t%s\nJira Link:\t%s\nDescription:\t%s\nAutomation Needed:\t%t\tRequired Validation:\t%s", tp.Name, tp.UserStory, tp.UserStoryLink, tp.Description, tp.AutomationNeeded, tp.RequiredValidation)
}

func (tp TestPlanTemplate) MDString() string {

	title := "#  " + tp.Name + "\n\n"
	link := fmt.Sprintf("User Story Link:\t[%s](%s)\n\n", tp.UserStory, tp.UserStoryLink)
	table, _ := tableBuilder(tp)
	components := fmt.Sprintf("User Story Components:\t%s\n\n", tp.UserStoryComponents)
	auto := fmt.Sprintf("Requires Automation:\t%t\n\n", tp.AutomationNeeded)
	description := fmt.Sprintf("## Description:\n%s\n\n", tp.Description)
	testplan := "## Test Validation"

	return title + link + table + components + auto + "***\n" + description + "***\n" + testplan
}

func tableBuilder(tp TestPlanTemplate) (string, error) {

	table := fmt.Sprintf("| PM  | DEV | QA |\n| :----: | :----: | :----: |\n| %s | %s | %s |\n\n", tp.PM, tp.Dev, tp.QA)

	return table, nil

}
