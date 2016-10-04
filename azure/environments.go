package azure

type environment int64

const (
	AzureCloud environment = iota
	AzureGermanCloud
)

type Endpoint struct {
	resourceManagerEndpointUrl string
	activeDirectoryEndpointUrl string
}

var environments = []Endpoint{
	// AzureCloud
	{"https://management.azure.com", "https://login.microsoftonline.com"},
	// GermanCloud
	{"https://management.microsoftazure.de", "https://login.microsoftonline.de"},
}

func GetEndpoints(e environment) Endpoint {
	return environments[e]
}
