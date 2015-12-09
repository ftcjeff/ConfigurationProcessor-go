package generators

import (
	"strconv"
	"strings"

	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func ServerGenerator(model types.ModelType, channel chan []types.ServerType) {
	defer logger.Trace(logger.Enter())

	var servers []types.ServerType

	definition := model.Definition
	for _, element := range definition.Elements {
		data := element.Data

		for _, tier := range data.Tiers {
			tierName := genTiername(data.Name, data.Id, tier.Id)

			for i := 0; i < tier.MemberCount; i++ {
				var server types.ServerType

				server.Tier = tierName
				server.Hostname = genHostname(tierName, i)

				servers = append(servers, server)
			}
		}
	}

	channel <- servers
}

func genTiername(elementName string, dataId, tierId int) string {
	defer logger.Trace(logger.Enter())

	return normalize(elementName) + "-D" + strconv.Itoa(dataId) + "-T" + strconv.Itoa(tierId)
}

func genHostname(tierName string, memberId int) string {
	defer logger.Trace(logger.Enter())

	return tierName + "-M" + strconv.Itoa(memberId)
}

func normalize(value string) string {
	value = strings.Replace(value, " ", "", -1)
	value = strings.Replace(value, "-", "_", -1)
	return value
}
