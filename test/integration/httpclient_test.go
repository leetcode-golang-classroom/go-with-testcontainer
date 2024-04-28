package integration_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	http_client "github.com/leetcode-golang-classroom/go-with-testcontainer/internal/http-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestHttpClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Client HTTP Suite")
}

var _ = Describe("HttpClient", Ordered, func() {
	var container testcontainers.Container
	var ctx context.Context
	var mappedPort string

	BeforeAll(func() {
		ctx = context.Background()
		req := testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Dockerfile:    "./test/integration/test-server/Dockerfile",
				Context:       "./../../",
				PrintBuildLog: true,
				KeepImage:     true,
			},
			ExposedPorts: []string{"80/tcp"},
			WaitingFor:   wait.ForHTTP("/").WithStartupTimeout(10 * time.Second),
		}
		c, err := testcontainers.GenericContainer(ctx,
			testcontainers.GenericContainerRequest{
				ContainerRequest: req,
				Started:          true,
			},
		)
		Expect(err).NotTo(HaveOccurred())
		port, err := c.MappedPort(ctx, "80")
		Expect(err).NotTo(HaveOccurred())
		mappedPort = port.Port()
		container = c
	})
	AfterAll(func() {
		err := container.Terminate(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	When("fetching ages from a mock integration API", func() {
		It("Should successfully GET the Age of Sig from API server", func() {
			fmt.Println(mappedPort)
			baseUrl := fmt.Sprintf("http://localhost:%s", mappedPort)
			client, err := http_client.NewHttpClient(baseUrl)
			Expect(err).NotTo(HaveOccurred())
			res, err := client.GetAge("Sig")
			Expect(err).NotTo(HaveOccurred())
			Expect(res.Age).To(Equal(62))
		})
	})
})
