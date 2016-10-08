package azure

type Endpoints struct {
	resourceManagerEndpointUrl string
	activeDirectoryEndpointUrl string
}

func GetEndpoints(location string) Endpoints {
	var e Endpoints

	switch location {
	case GermanyCentral:
	case GermanyEast:
		e = Endpoints{"https://management.microsoftazure.de", "https://login.microsoftonline.de"}
	case ChinaEast:
	case ChinaNorth:
		e = Endpoints{"https://management.chinacloudapi.cn", "https://login.chinacloudapi.cn"}
	case USGovIowa:
	case USGovVirginia:
		e = Endpoints{"https://management.usgovcloudapi.net", "https://login.microsoftonline.com"}
	default:
		e = Endpoints{"https://management.azure.com", "https://login.microsoftonline.com"}
	}

	return e
}
