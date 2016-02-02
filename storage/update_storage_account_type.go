package storage

import "github.com/jen20/riviera/azure"

type UpdateStorageAccountTypeResponse struct {
	AccountType *string `mapstructure:"accountType"`
}

type UpdateStorageAccountType struct {
	Name              string  `json:"-"`
	ResourceGroupName string  `json:"-"`
	AccountType       *string `json:"accountType,omitempty"`
}

func (command UpdateStorageAccountType) ApiInfo() azure.ApiInfo {
	return azure.ApiInfo{
		ApiVersion:  apiVersion,
		Method:      "PATCH",
		URLPathFunc: storageDefaultURLPathFunc(command.ResourceGroupName, command.Name),
		//SkipArmBoilerplate: true,
		ResponseTypeFunc: func() interface{} {
			return &UpdateStorageAccountTypeResponse{}
		},
	}
}
