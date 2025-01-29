package TestplanForm

import (
	"errors"

	TestPlanTemplate "github.com/carlogy/TestPlanMaker/internal/TestPlanTemplate"
	"github.com/charmbracelet/huh"
)

func BuilForm(tp *TestPlanTemplate.TestPlanTemplate) *huh.Form {

	form := huh.NewForm(

		huh.NewGroup(
			huh.NewInput().
				Title("Test Plan Title:").
				Value(&tp.Name).
				Validate(func(str string) error {
					if len(str) < 1 {
						return errors.New("Test Plan Name is required.")
					}
					return nil
				}),

			huh.NewInput().
				Title("User Story:").
				Value(&tp.UserStory),

			huh.NewInput().
				Title("Jira Link:").
				Value(&tp.JiraLink),

			huh.NewSelect[string]().
				Title("Validation Type:").
				Options(
					huh.NewOption("UI validation", "UI"),
					huh.NewOption("API Validation", "API"),
					huh.NewOption("UI & API Validation", "Both"),
				).Value(&tp.RequiredValidation),

			huh.NewText().
				Title("Description:").
				CharLimit(800).
				Value(&tp.Description),

			huh.NewInput().Title("Save file to path:").
				CharLimit(260).
				Value(&tp.SavePath),

			huh.NewConfirm().
				Title("Is Automation needed?").
				Value(&tp.AutomationNeeded),
		),
	)

	return form
}
