package storage

import "github.com/jen20/riviera/azure"

type DeleteStorageAccount struct {
	Name              string `json:"-"`
	ResourceGroupName string `json:"-"`
}

func (command DeleteStorageAccount) ApiInfo() azure.ApiInfo {
	return azure.ApiInfo{
		ApiVersion:  apiVersion,
		Method:      "DELETE",
		URLPathFunc: storageDefaultURLPathFunc(command.ResourceGroupName, command.Name),
		ResponseTypeFunc: func() interface{} {
			return nil
		},
	}
}
