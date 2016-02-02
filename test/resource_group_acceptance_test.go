package test

import (
	"testing"

	"github.com/jen20/riviera/azure"
)

func TestAccCreateResourceGroup(t *testing.T) {
	rgName := RandPrefixString("testrg_", 20)

	Test(t, TestCase{
		Steps: []Step{
			&StepCreateResourceGroup{
				Name:     rgName,
				Location: azure.WestUS,
			},
			&StepAssert{
				Checks: []AssertFunc{
					CheckStringProperty("resourcegroup", "Name", rgName),
					CheckStringProperty("resourcegroup", "Location", azure.WestUS),
				},
			},
		},
	})
}
