package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"port-service/internal/core/repository"
	"port-service/internal/core/service"
	"port-service/internal/data-source/storage/inmem"
	"port-service/internal/transport/handler"
	"testing"
)

type HttpTestSuite struct {
	suite.Suite
	storage     *inmem.InmemStore
	portService handler.PortService
	httpServer  handler.HttpServer
}

// structure to compare Port data from request and in Storage
// no need to compare CreatedAt and UpdatedAt fields
type ComparePort struct {
	Id          string
	Name        string
	Code        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
}

func NewHttpTestSuite() *HttpTestSuite {
	suite := &HttpTestSuite{}

	// create inmem storage
	storage := inmem.New()

	suite.storage = storage

	// create port repository
	portStoreRepo := repository.NewPortRepository(storage)

	// create port service
	suite.portService = service.NewPortService(portStoreRepo)

	// create handler server with application injected
	suite.httpServer = handler.NewHttpServer(suite.portService)

	return suite
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, NewHttpTestSuite())
}

func (suite *HttpTestSuite) TestFetchingPorts() {
	portsRequest, err := os.ReadFile("fixtures/ports_request.json")
	require.NoError(suite.T(), err)

	requestPorts := getRequestPorts(suite.T(), portsRequest)

	// create POST /ports request
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer(portsRequest))

	w := httptest.NewRecorder()

	// run request
	suite.httpServer.FetchPorts(w, req)

	res := w.Result()
	defer res.Body.Close()

	// read response body
	_, err = io.ReadAll(res.Body)
	require.NoError(suite.T(), err)

	comparePorts := convertStoreToComparePorts(suite.storage.GetMap())

	require.Equal(suite.T(), http.StatusOK, res.StatusCode)
	require.Equal(suite.T(), comparePorts, requestPorts)
}

func (suite *HttpTestSuite) TestFetchingPorts_badRequest() {
	// create POST /ports request
	req := httptest.NewRequest(http.MethodPost, "/ports", bytes.NewBuffer([]byte("bad_request")))

	w := httptest.NewRecorder()

	// run request
	suite.httpServer.FetchPorts(w, req)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(suite.T(), http.StatusBadRequest, res.StatusCode)
}

// get Ports from fixture to compare
func getRequestPorts(t *testing.T, portsJSON []byte) map[string]ComparePort {
	t.Helper()
	var ports map[string]ComparePort
	err := json.Unmarshal(portsJSON, &ports)
	for portId, port := range ports {
		port.Id = portId
		ports[portId] = port
	}
	fmt.Println(ports)
	require.NoError(t, err)
	return ports
}

func convertStoreToComparePorts(storePorts map[string]repository.Port) map[string]ComparePort {
	result := make(map[string]ComparePort)

	for portId, storePort := range storePorts {
		result[portId] = ComparePort{
			storePort.Id,
			storePort.Name,
			storePort.Code,
			storePort.City,
			storePort.Country,
			storePort.Alias,
			storePort.Regions,
			storePort.Coordinates,
			storePort.Province,
			storePort.Timezone,
			storePort.Unlocs,
		}
	}
	return result
}
