package container

import (
	bfactory "github.com/pip-services-samples/pip-services-beacons-go/build"
	cproc "github.com/pip-services3-go/pip-services3-container-go/container"
	rpcbuild "github.com/pip-services3-go/pip-services3-rpc-go/build"
)

type BeaconsProcess struct {
	cproc.ProcessContainer
}

func NewBeaconsProcess() *BeaconsProcess {

	bp := BeaconsProcess{}
	bp.ProcessContainer = *cproc.NewEmptyProcessContainer()
	bp.AddFactory(bfactory.NewBeaconsServiceFactory())
	bp.AddFactory(rpcbuild.NewDefaultRpcFactory())
	return &bp
}
