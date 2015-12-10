package segment

import (
	"errors"
	"net"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func GetNetworkInfo(model types.ModelType, segmentType string) (types.IpAddressDetailsType, error) {
	defer logger.Trace(logger.Enter())

	var info types.IpAddressDetailsType

	definition := model.Definition
	topology := definition.Topology

	if len(segmentType) == 0 {
		segmentType = "internal"
	}

	for _, segment := range topology.NetworkSegments {
		if segment.Name != segmentType {
			continue
		}

		ip, _, _ := net.ParseCIDR(segment.Cidr)

		info.IpAddress = ip.String()
		info.Netmask = segment.Netmask
		info.Gateway = segment.Gateway

		return info, nil
	}

	return info, errors.New("Could not find network segment")
}
