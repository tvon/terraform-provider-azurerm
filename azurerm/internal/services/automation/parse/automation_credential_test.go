package parse

import (
	"testing"
)

func TestAutomationCredentialId(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *AutomationCredentialId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000",
			Expected: nil,
		},
		{
			Name:     "No Resource Groups Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/",
			Expected: nil,
		},
		{
			Name:     "Resource Group ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/",
			Expected: nil,
		},
		{
			Name:     "Missing Automation Accounts Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/",
			Expected: nil,
		},
		{
			Name:     "Automation Account ID",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1",
			Expected: nil,
		},
		{
			Name:     "Missing Credential Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/credentials/",
			Expected: nil,
		},
		{
			Name:  "Automation Credential ID",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/credentials/credential1",
			Expected: &AutomationCredentialId{
				ResourceGroup: "resGroup1",
				AccountName:   "account1",
				Name:          "credential1",
			},
		},
		{
			Name:     "Wrong Casing",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Automation/automationAccounts/account1/Credentials/credential1",
			Expected: nil,
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := AutomationCredentialID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}
		if actual.AccountName != v.Expected.AccountName {
			t.Fatalf("Expected %q but got %q for Account Name", v.Expected.AccountName, actual.AccountName)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for Resource Group", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
