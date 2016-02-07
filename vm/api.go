package vm

import "fmt"

const apiVersion = "2015-05-01-preview"
const apiProvider = "Microsoft.Compute"

func virtualMachineDefaultURLPathFunc(resourceGroupName, vmName string) func() string {
	return func() string {
		return fmt.Sprintf("resourceGroups/%s/providers/%s/virtualMachines/%s", resourceGroupName, apiProvider, vmName)
	}
}
