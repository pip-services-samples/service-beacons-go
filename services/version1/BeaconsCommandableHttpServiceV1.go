package services1

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cservices "github.com/pip-services3-go/pip-services3-rpc-go/services"
)

type BeaconsCommandableHttpServiceV1 struct {
	*cservices.CommandableHttpService
}

func NewBeaconsCommandableHttpServiceV1() *BeaconsCommandableHttpServiceV1 {
	c := &BeaconsCommandableHttpServiceV1{}
	c.CommandableHttpService = cservices.InheritCommandableHttpService(c, "v1/beacons")

	c.DependencyResolver.Put("controller", cref.NewDescriptor("pip-services-beacons", "controller", "*", "*", "1.0"))
	return c
}

func (c *BeaconsCommandableHttpServiceV1) Register() {
	if !c.SwaggerAuto && c.SwaggerEnabled {
		c.RegisterOpenApiSpecFromFile("./swagger/beacons_v1.yaml")
	}
	c.CommandableHttpService.Register()
}
