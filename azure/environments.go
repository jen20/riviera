package azure

type environment int64

const (
	AzureCloud environment = iota
	AzureGermanCloud
	AzureChinaCloud
	AzureUSGovernment
)

type Endpoints struct {
	resourceManagerEndpointUrl string
	activeDirectoryEndpointUrl string
}

var environments = []Endpoints{
	// AzureCloud
	{"https://management.azure.com", "https://login.microsoftonline.com"},
	// AzureGermanCloud
	{"https://management.microsoftazure.de", "https://login.microsoftonline.de"},
	// AzureChinaCloud
	{"https://management.chinacloudapi.cn", "https://login.chinacloudapi.cn"},
	// AzureUSGovernment
	{"https://management.usgovcloudapi.net", "https://login.microsoftonline.com"},
}

func StringToEnvironment(str string) environment {
	var env environment

	switch str {
	case "AzureCloud":
		env = AzureCloud
	case "AzureGermanCloud":
		env = AzureGermanCloud
	case "AzureChinaCloud":
		env = AzureChinaCloud
	case "AzuerUSGovernment":
		env = AzureUSGovernment
	default:
		env = AzureCloud
	}

	return env
}

func GetEndpoints(e environment) Endpoints {
	return environments[e]
}
