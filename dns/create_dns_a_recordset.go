package dns

import "github.com/jen20/riviera/azure"

type CreateDNSARecordSet struct {
	Name              string             `json:"-"`
	ResourceGroupName string             `json:"-"`
	ZoneName          string             `json:"-"`
	Location          string             `json:"-" riviera:"location"`
	Tags              map[string]*string `json:"-" riviera:"tags"`
	TTL               int                `json:"-"`
	IPv4Addresses     []string           `json:"-"`
}

func (command CreateDNSARecordSet) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "PUT",
		URLPathFunc: dnsRecordSetDefaultURLPathFunc(command.ResourceGroupName, command.ZoneName, "A", command.Name),
		RequestPropertiesFunc: func() interface{} {
			var addresses []interface{}
			for _, v := range command.IPv4Addresses {
				addresses = append(addresses, struct {
					IPv4Address string `json:"ipv4Address`
				}{
					IPv4Address: v,
				})
			}

			return struct {
				TTL      int         `json:"TTL"`
				ARecords interface{} `json:ARecords`
			}{
				TTL:      command.TTL,
				ARecords: addresses,
			}
		},
		ResponseTypeFunc: func() interface{} {
			return nil
		},
	}
}
