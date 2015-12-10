package builders

import (
	"os"
	"sync"

	"github.com/ftcjeff/ConfigurationProcessor/flags"
	"github.com/ftcjeff/ConfigurationProcessor/logger"
	"github.com/ftcjeff/ConfigurationProcessor/types"
)

func HostsBuilder(model types.ModelType, wg sync.WaitGroup) {
	defer logger.Trace(logger.Enter())

	wg.Add(1)

	servers := model.Models.Servers

	fileName := flags.FlagOutputPath + "/" + "net"
	os.Mkdir(fileName, 0777)
	fileName = fileName + "/" + "hosts"
	if fp, err := os.Create(fileName); err != nil {
		panic(err)
	} else {
		defer fp.Close()

		for _, s := range servers {
			fp.WriteString(s.Hostname + "\t\t" + s.NetworkInfo.IpAddress + "\n")
		}

		logger.Log("Created file: " + fileName)
	}
}
