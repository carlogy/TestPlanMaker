package TestPlanTemplate

import "fmt"

type TestPlanTemplate struct {
	Name               string
	UserStory          string
	JiraLink           string
	AutomationNeeded   bool
	RequiredValidation string
	Description        string
	SavePath           string
}

// Factory to test stringer function. For actually functionality will utilize the NewTestPlanBuild factory function to instanstiate a struct and return it's pointer to update with form inputs.

func NewTestPlan(name, ticket, jiraLink, validation, description, savePath string, automation bool) *TestPlanTemplate {
	tp := &TestPlanTemplate{
		Name:               name,
		UserStory:          ticket,
		JiraLink:           jiraLink,
		RequiredValidation: validation,
		Description:        description,
		AutomationNeeded:   automation,
		SavePath:           savePath,
	}
	return tp
}

func NewTestPlanBuild() *TestPlanTemplate {
	return &TestPlanTemplate{}
}

func (tp TestPlanTemplate) String() string {
	return fmt.Sprintf("TestPlan:\nName:\t%s\nUser Story:\t%s\nJira Link:\t%s\nDescription:\t%s\nAutomation Needed:\t%t\tRequired Validation:\t%s", tp.Name, tp.UserStory, tp.JiraLink, tp.Description, tp.AutomationNeeded, tp.RequiredValidation)
}
