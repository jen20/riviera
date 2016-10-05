package azure

type environment int64

const (
	AzureCloud environment = iota
	AzureGermanCloud
	AzureChinaCloud
	AzureUSGovernment
)

type Endpoint struct {
	resourceManagerEndpointUrl string
	activeDirectoryEndpointUrl string
}

var environments = []Endpoint{
	// AzureCloud
	{"https://management.azure.com", "https://login.microsoftonline.com"},
	// AzureGermanCloud
	{"https://management.microsoftazure.de", "https://login.microsoftonline.de"},
	// AzureChinaCloud
	{"https://management.chinacloudapi.cn", "https://login.chinacloudapi.cn"},
	// AzureUSGovernment
	{"https://management.usgovcloudapi.net", "https://login.microsoftonline.com"},
}

func GetEndpoints(e environment) Endpoint {
	return environments[e]
}
