package types

type NetworkChannel struct {
	Id          int
	NetworkInfo NetworkInfoType
}

type TopologyType struct {
	Product         ProductType  `json:"product"`
	NetworkSegments SegmentsType `json:"segments"`
}

type SegmentType struct {
	Name         string `json:"name:`
	Cidr         string `json:"cidr"`
	Netmask      string `json:"netmask,omitempty"`
	Gateway      string `json:"gateway,omitempty"`
	StartAddress string `json:"start-address,omitempty"`

	NetworkInfo NetworkInfoType `json:"network-info,omitempty"`
}

type SegmentsType []SegmentType

type NetworkInfoType struct {
	NextIndex int `json:"next-index"`
}

type IpAddressDetailsType struct {
	IpAddress string `json:"ip-address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
}
