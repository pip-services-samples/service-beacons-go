package test_persistence

import (
	"os"
	"testing"

	persist "github.com/pip-services-samples/service-beacons-go/persistence"
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
)

type BeaconsMongoDbPersistenceTest struct {
	persistence *persist.BeaconsMongoDbPersistence
	fixture     *BeaconsPersistenceFixture
}

func newBeaconsMongoDbPersistenceTest() *BeaconsMongoDbPersistenceTest {
	mongoUri := os.Getenv("MONGO_SERVICE_URI")
	mongoHost := os.Getenv("MONGO_SERVICE_HOST")

	if mongoHost == "" {
		mongoHost = "localhost"
	}
	mongoPort := os.Getenv("MONGO_SERVICE_PORT")
	if mongoPort == "" {
		mongoPort = "27017"
	}

	mongoDatabase := os.Getenv("MONGO_SERVICE_DB")
	if mongoDatabase == "" {
		mongoDatabase = "test"
	}

	// Exit if mongo connection is not set
	if mongoUri == "" && mongoHost == "" {
		return nil
	}

	persistence := persist.NewBeaconsMongoDbPersistence()
	persistence.Configure(cconf.NewConfigParamsFromTuples(
		"connection.uri", mongoUri,
		"connection.host", mongoHost,
		"connection.port", mongoPort,
		"connection.database", mongoDatabase,
	))

	fixture := NewBeaconsPersistenceFixture(persistence)

	return &BeaconsMongoDbPersistenceTest{
		persistence: persistence,
		fixture:     fixture,
	}
}

func (c *BeaconsMongoDbPersistenceTest) setup(t *testing.T) {
	err := c.persistence.Open("")
	if err != nil {
		t.Error("Failed to open persistence", err)
	}

	err = c.persistence.Clear("")
	if err != nil {
		t.Error("Failed to clear persistence", err)
	}
}

func (c *BeaconsMongoDbPersistenceTest) teardown(t *testing.T) {
	err := c.persistence.Close("")
	if err != nil {
		t.Error("Failed to close persistence", err)
	}
}

func TestBeaconsMongoDbPersistence(t *testing.T) {
	c := newBeaconsMongoDbPersistenceTest()
	if c == nil {
		return
	}

	c.setup(t)
	t.Run("CRUD Operations", c.fixture.TestCrudOperations)
	c.teardown(t)

	c.setup(t)
	t.Run("Get With Filters", c.fixture.TestGetWithFilters)
	c.teardown(t)
}
