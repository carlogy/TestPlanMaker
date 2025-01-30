package TestPlanTemplate

import (
	"testing"
)

func TestStringerTestPlanMethod(t *testing.T) {

	cases := []struct {
		testPlan TestPlanTemplate
		expected string
	}{
		{
			testPlan: *NewTestPlan("name", "Ticket#", "https://jiralink.com", "UI Validation", "This is a test ticket", "", false),
			expected: "TestPlan:\nName:\tname\nUser Story:\tTicket#\nUser Story Link:\thttps://jiralink.com\nDescription:\tThis is a test ticket\n,Automation Needed:\tfalse\tRequired Validation:\tUI Validation",
		},
	}

	for i := 0; i < len(cases); i++ {
		actual := cases[i].testPlan.String()

		if actual != cases[i].expected {
			t.Errorf("Actual:\t%s does not match\nExpected:\t %s", actual, cases[i].expected)
		}
	}

}
