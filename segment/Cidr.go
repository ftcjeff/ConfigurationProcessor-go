package segment

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
)

func GetIpAddressesForCidr(cidr string) ([]string, int, error) {
	defer logger.Trace(logger.Enter())

	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		panic(err)
	}

	ipBytes := IpToBytes(ipNet.IP)
	logger.Log(fmt.Sprintf("ipBytes: %+v", ipBytes))

	ones, bits := ipNet.Mask.Size()
	numAddrs := int(math.Exp2(float64(bits - ones)))

	addresses := GenerateAddresses(ipBytes, numAddrs)

	return addresses, numAddrs, nil
}

func IpToBytes(address net.IP) [4]int {
	tokens := strings.Split(address.String(), ".")

	var bytes [4]int

	bytes[0], _ = strconv.Atoi(tokens[0])
	bytes[1], _ = strconv.Atoi(tokens[1])
	bytes[2], _ = strconv.Atoi(tokens[2])
	bytes[3], _ = strconv.Atoi(tokens[3])

	return bytes
}

func CidrToIp(val string) (string, string) {
	tokens := strings.Split(val, "/")

	address := tokens[0]
	cidr := tokens[1]

	return address, cidr
}

func GenerateAddresses(ipBytes [4]int, numAddrs int) []string {
	var addresses []string

	for i := 0; i < numAddrs; i++ {
		addr := fmt.Sprintf("%d.%d.%d.%d", ipBytes[0], ipBytes[1], ipBytes[2], ipBytes[3])
		addresses = append(addresses, addr)

		if ipBytes[3] >= 255 {
			ipBytes[3] = 0
			ipBytes[2]++

			if ipBytes[2] >= 255 {
				ipBytes[2] = 0
				ipBytes[1]++
			}

			if ipBytes[1] >= 255 {
				ipBytes[1] = 0
				ipBytes[0]++
			}
		} else {
			ipBytes[3]++
		}
	}

	return addresses
}
