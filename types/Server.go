package types

type ServerType struct {
	Tier        string               `json:"tier"`
	Hostname    string               `json:"hostname"`
	NetworkInfo IpAddressDetailsType `json:"network-info"`
}
