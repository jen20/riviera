package search

import "github.com/jen20/riviera/azure"

type CreateOrUpdateSearchServiceResponse struct {
	ID       *string             `mapstructure:"id"`
	Name     *string             `mapstructure:"name"`
	Location *string             `mapstructure:"location"`
	Tags     *map[string]*string `mapstructure:"tags"`
}

type CreateOrUpdateSearchService struct {
	Name               string             `json:"-"`
	ResourceGroupName  string             `json:"-"`
	Location           string             `json:"-" riviera:"location"`
	Tags               map[string]*string `json:"-" riviera:"tags"`
	Sku                *string            `json:"sku,omitempty"`
	ReplicaCount       *string            `json:"replicaCount,omitempty"`
	PartitionCount     *string            `json:"partitionCount,omitempty"`
	Status             *string            `mapstructure:"status"`
	StatusDetails      *string            `mapstructure:"statusDetails"`
	ProvisioningStatus *string            `mapstructure:"provisioningStatus"`
}

func (s CreateOrUpdateSearchService) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "PUT",
		URLPathFunc: searchServiceDefaultURLPath(s.ResourceGroupName, s.Name),
		ResponseTypeFunc: func() interface{} {
			return &CreateOrUpdateSearchServiceResponse{}
		},
	}
}
