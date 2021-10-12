package services1

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cservices "github.com/pip-services3-go/pip-services3-grpc-go/services"
)

type BeaconsCommandableGrpcServiceV1 struct {
	*cservices.CommandableGrpcService
}

func NewBeaconsCommandableGrpcServiceV1() *BeaconsCommandableGrpcServiceV1 {
	c := &BeaconsCommandableGrpcServiceV1{}
	c.CommandableGrpcService = cservices.InheritCommandableGrpcService(c, "v1.beacons")
	c.DependencyResolver.Put("controller", cref.NewDescriptor("pip-services-beacons", "controller", "*", "*", "1.0"))
	return c
}
