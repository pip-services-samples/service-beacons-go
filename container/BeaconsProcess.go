package container

import (
	factory "github.com/pip-services-samples/pip-services-beacons-go/build"
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	gbuild "github.com/pip-services3-go/pip-services3-grpc-go/build"
	rbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
)

type BeaconsProcess struct {
	cproc.ProcessContainer
}

func NewBeaconsProcess() *BeaconsProcess {
	c := &BeaconsProcess{
		ProcessContainer: *cproc.NewProcessContainer("beacons", "Beacons microservice"),
	}

	c.AddFactory(factory.NewBeaconsServiceFactory())
	c.AddFactory(rbuild.NewDefaultRpcFactory())
	c.AddFactory(gbuild.NewDefaultGrpcFactory())

	return c
}
