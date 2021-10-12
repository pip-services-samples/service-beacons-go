package main

import (
	"os"

	cont "github.com/pip-services-samples/service-beacons-go/containers"
)

func main() {
	proc := cont.NewBeaconsProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
