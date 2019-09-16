package network

import (
	"testing"

	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
)

func TestParseRouteTable(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *RouteTableResourceID
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Route Table Segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo",
			Expected: nil,
		},
		{
			Name:     "No Route Table Value",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/routeTables/",
			Expected: nil,
		},
		{
			Name:  "Completed",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/foo/routeTables/example",
			Expected: &RouteTableResourceID{
				Name: "example",
				Base: azure.ResourceID{
					ResourceGroup: "foo",
				},
			},
		},
	}

	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := ParseRouteTableResourceID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}

		if actual.Base.ResourceGroup != v.Expected.Base.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.Base.ResourceGroup, actual.Base.ResourceGroup)
		}
	}
}
