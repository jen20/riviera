package vm

import "github.com/jen20/riviera/azure"

type Plan struct {
	Name      int `json:"name" mapstructure:"name"`
	Publisher int `json:"publisher" mapstructure:"publisher"`
	Product   int `json:"product" mapstructure:"product"`
}

type ImageReference struct {
	Publisher string `json:"publisher" mapstructure:"publisher"`
	Offer     string `json:"offer" mapstructure:"offer"`
	SKU       string `json:"sku" mapstructure:"sku"`
	Version   string `json:"version" mapstructure:"version"`
}

type OSDisk struct {
	Name         string  `json:"name" mapstructure:"name"`
	URI          string  `json:"Uri" mapstructure:"Uri"`
	Caching      *string `json:"caching,omitempty" mapstructure:"caching"`
	CreateOption string  `json:"createOption" mapstructure:"createOption"`
}

type DataDisk struct {
	Name           string `json:"name" mapstructure:"name"`
	DiskSizeGB     string `json:"diskSizeGB" mapstructure:"diskSizeGB"`
	LUN            int    `json:"lun" mapstructure:"lun"`
	VHD            string `json:"vhd" mapstructure:"vhd"`
	CreationOption string `json:"creationOption" mapstructure:"creationOption"`
}

type StorageProfile struct {
	ImageReference *ImageReference `json:"imageReference,omitempty" mapstructure:"imageReference"`
	OSDisk         OSDisk          `json:"oSDisk" mapsstructure:"osDisk"`
	DataDisks      []DataDisk      `json:"dataDisks,omitempty" mapstructure:"dataDisks"`
}

type SourceVault struct {
	ID string `json:"id" mapstructure:"id"`
}

type VaultCertificate struct {
	CertificateURL   string  `json:"certificateUrl" mapstructure:"certificateUrl"`
	CertificateStore *string `json:"certificateStore" mapstructure:"certificateStore"`
}

type Secrets struct {
	SourceVault       SourceVault        `json:"sourceVault" mapstructure:"sourceVault"`
	VaultCertificates []VaultCertificate `json:"vaultCertificates,omitempty" mapstructure:"vaultCertificates"`
}

type WindowsConfiguration struct {
	ProvisionVMAgent          *bool                      `json:"provisionVMAgent,omitempty" mapstructure:"provisionVMAgent"`
	EnableAutomaticUpdates    *bool                      `json:"enableAutomaticUpdates,omitempty" mapstructure:"enableAutomaticUpdates"`
	WinRM                     []WinRM                    `json:"winRM,omitempty" mapstructure:"winRM"`
	AdditionalUnattendContent *AdditionalUnattendContent `json:"additionalUnattendContent" mapstructure:"additionalUnattendContent"`
}

type AdditionalUnattendContent struct {
	Pass        string `json:"pass" mapstructure:"pass"`
	Component   string `json:"component" mapstructure:"component"`
	SettingName string `json:"settingName" mapstructure:"settingName"`
	Content     string `json:"content" mapstructure:"content"`
}

type WinRM struct {
	Protocol       string  `json:"protocol" mapstructure:"protocol"`
	CertificateURL *string `json:"certificateUrl,omitempty" mapstructure:"certificateUrl"`
}

// This is undocumented so \_(ツ)_/¯
type SSH struct {
	PublicKeys []SSHPublicKey `json:"publicKeys,omitempty mapstructure:"publicKeys"`
}

// This is undocumented so \_(ツ)_/¯
type SSHPublicKey struct {
	Path    string `json:"path" mapstructure:"path"`
	KeyData string `json:"keyData" mapstructure:"keyData"`
}

type LinuxConfiguration struct {
	DisablePasswordAuthentication *bool `json:"disablePasswordAuthentication,omitempty" mapstructure:"disablePasswordAuthentication"`
	SSH                           *SSH  `json:"ssh" mapstructure:"ssh"`
}

type OSProfile struct {
	ComputerName         *string               `json:"computerName" mapstructure:"computerName"`
	AdminUsername        string                `json:"adminUsername" mapstructure:"adminUsername`
	AdminPassword        string                `json:"adminPassword" mapstructure:"adminPassword`
	CustomData           *string               `json:"customData" mapstructure:"customData"`
	WindowsConfiguration *WindowsConfiguration `json:"windowsConfiguration" mapstructure:"windowsConfiguration"`
	LinuxConfiguration   *LinuxConfiguration   `json:"linuxConfiguration" mapstructure:"linuxConfiguration"`
	Secrets              *Secrets              `json:"secrets" mapstructure:"secrets"`
}

type AvailabilitySetRef struct {
	ID string `json:"id" mapstructure:"id"`
}

type NetworkInterfaceRef struct {
	ID string `json:"id" mapstructure:"id"`
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterfaceRef `json:"networkInterfaces" mapstructure:"networkInterfaces"`
}

type HardwareProfile struct {
	VMSize string `json:"vmSize" mapstructure:"vmSize"`
}

type CreateOrUpdateVirtualMachineResponse struct {
	ID              string              `mapstructure:"id"`
	Name            string              `mapstructure:"name"`
	Location        string              `mapstructure:"location"`
	Tags            map[string]*string  `mapstructure:"tags"`
	AvailabilitySet *AvailabilitySetRef `mapstructure:"availabilitySet"`
	HardwareProfile HardwareProfile     `mapstructure:"hardwareProfile"`
	StorageProfile  StorageProfile      `mapstructure:"storageProfile"`
	OSProfile       OSProfile           `mapstructure:"osProfile"`
	NetworkProfile  NetworkProfile      `mapstructure:"networkProfile"`
}

type CreateOrUpdateVirtualMachine struct {
	Name              string              `json:"-"`
	ResourceGroupName string              `json:"-"`
	Location          string              `json:"-" riviera:"location"`
	Tags              map[string]*string  `json:"-" riviera:"tags"`
	Plan              *Plan               `json:"-" riviera:"plan"`
	AvailabilitySet   *AvailabilitySetRef `json:"availabilitySet"`
	HardwareProfile   HardwareProfile     `json:"hardwareProfile"`
	StorageProfile    StorageProfile      `json:"storageProfile"`
	OSProfile         OSProfile           `json:"osProfile"`
	NetworkProfile    NetworkProfile      `json:"networkProfile"`
}

func (command CreateOrUpdateVirtualMachine) APIInfo() azure.APIInfo {
	return azure.APIInfo{
		APIVersion:  apiVersion,
		Method:      "PUT",
		URLPathFunc: virtualMachineDefaultURLPathFunc(command.ResourceGroupName, command.Name),
		ResponseTypeFunc: func() interface{} {
			return &CreateOrUpdateVirtualMachineResponse{}
		},
	}
}
