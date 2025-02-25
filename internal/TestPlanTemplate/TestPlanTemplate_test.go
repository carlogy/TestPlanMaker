package TestPlanTemplate

import (
	"testing"
)

func TestMDStringerPlanMethod(t *testing.T) {
	tests := []struct {
		tp       TestPlanTemplate
		expected string
	}{
		{
			tp: TestPlanTemplate{
				Name: "Test Plan Test",
				TicketDetails: TicketDetails{
					UserStory:           "TST-100",
					UserStoryLink:       "https://jira.com/TST-100",
					UserStoryComponents: "Test_PlanBuilder",
					AutomationNeeded:    true,
					PM:                  "Foo",
					Dev:                 "Bar",
					QA:                  "Bazz",
				},
				RequiredValidation: "UI",
				Description:        "This is a test plan template.",
				SavePath:           "",
				WriteToFile:        false,
			},
			expected: "#  Test Plan Test\n\n**User Story Link:**\t[TST-100](https://jira.com/TST-100)\n\n| PM  | SE | QAE |\n| :----: | :----: | :----: |\n| Foo | Bar | Bazz |\n\n\n\n**User Story Components:**\tTest_PlanBuilder\n\n**Validation Type:**\tUI\n\n**Requires Automation:**\ttrue\n\n\n\n***\n\n##  Description:\nThis is a test plan template.\n\n\n\n***\n\n##  Test Validation\n\n",
		},
	}

	for _, test := range tests {

		got := test.tp.MDString()

		if got != test.expected {
			t.Errorf("Got:\t%s\nExpecting:\t%s\n", got, test.expected)
		}

	}

}
