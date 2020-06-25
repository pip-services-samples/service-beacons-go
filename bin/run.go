package main

import (
	"os"

	bproc "github.com/pip-services-samples/pip-services-beacons-go/container"
)

func main() {
	proc := bproc.NewBeaconsProcess()
	proc.SetConfigPath("./config/config.yml")
	proc.Run(os.Args)
}
