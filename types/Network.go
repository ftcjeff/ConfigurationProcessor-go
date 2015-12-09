package types

import (
	"net"
)

type NetworkChannel struct {
	Id          int
	NetworkInfo NetworkInfoType
}

type NetworkType struct {
	Name         string `json:"name:`
	Cidr         string `json:"cidr"`
	Netmask      string `json:"netmask,omitempty"`
	Gateway      string `json:"gateway,omitempty"`
	StartAddress string `json:"start-address,omitempty"`

	NetworkInfo NetworkInfoType `json:"network-info,omitempty"`
}

type NetworkInfoType struct {
	NextIndex int `json:"next-index"`
}

type IpAddressDetailsType struct {
	IpAddress net.IP     `json:"ip-address,omitempty"`
	Netmask   net.IPMask `json:"netmask,omitempty"`
	Gateway   net.IP     `json:"gateway,omitempty"`
}
