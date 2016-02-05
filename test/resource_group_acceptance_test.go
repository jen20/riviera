package test

import (
	"testing"

	"github.com/abdullin/seq"
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
				StateBagKey: "resourcegroup",
				Assertions: seq.Map{
					"Name":     rgName,
					"Location": azure.WestUS,
				},
			},
		},
	})
}
