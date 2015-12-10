package segment

import (
	"errors"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

var addressMaps map[string][]string
var addressIndices map[string]int

func GetNetworkInfo(model types.ModelType, segmentType string) (types.IpAddressDetailsType, error) {
	defer logger.Trace(logger.Enter())

	var info types.IpAddressDetailsType

	if addressMaps == nil {
		addressMaps = make(map[string][]string)
	}

	if addressIndices == nil {
		addressIndices = make(map[string]int)
	}

	definition := model.Definition
	topology := definition.Topology

	if len(segmentType) == 0 {
		segmentType = "internal"
	}

	for _, segment := range topology.NetworkSegments {
		if segment.Name != segmentType {
			continue
		}

		_, exists := addressMaps[segmentType]
		if !exists {
			addresses, _, _ := GetIpAddressesForCidr(segment.Cidr)
			addressMaps[segmentType] = addresses
			addressIndices[segmentType] = 0
		}

		index := addressIndices[segmentType]
		ip := addressMaps[segmentType][index]
		addressIndices[segmentType]++

		info.IpAddress = ip
		info.Netmask = segment.Netmask
		info.Gateway = segment.Gateway

		return info, nil
	}

	return info, errors.New("Could not find network segment")
}
