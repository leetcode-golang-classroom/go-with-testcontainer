package integration_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	db_client "github.com/leetcode-golang-classroom/go-with-testcontainer/internal/db-client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestDBClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Client Suite")
}

var _ = Describe("DB Client", Ordered, func() {
	var container testcontainers.Container
	var ctx context.Context
	var connString string
	BeforeAll(func() {
		ctx = context.Background()
		c, err := mysql.RunContainer(ctx,
			testcontainers.WithImage("mysql:5.7"),
			mysql.WithDatabase("foo"),
			mysql.WithUsername("root"),
			mysql.WithPassword("password"),
			mysql.WithScripts("./scripts/init.sql"),
		)
		Expect(err).NotTo(HaveOccurred())
		connString, err = c.ConnectionString(ctx)
		Expect(err).NotTo(HaveOccurred())
		container = c
	})
	AfterAll(func() {
		err := container.Terminate(ctx)
		Expect(err).NotTo(HaveOccurred())
	})
	When("fetching from the database", func() {
		It("Should successfully connect to the database", func() {
			fmt.Println(connString)
			log.Println(connString)
			client, err := db_client.NewDbClient(connString)
			Expect(err).NotTo(HaveOccurred())
			Expect(client).NotTo(BeNil())
		})
		It("Should successfuly fetch a profile from the database", func() {
			client, err := db_client.NewDbClient(connString)
			Expect(err).NotTo(HaveOccurred())
			Expect(client).NotTo(BeNil())
			res, err := client.Db.Query("SELECT * FROM profile")
			Expect(err).NotTo(HaveOccurred())
			defer res.Close()
			for res.Next() {
				var (
					id   int64
					name string
				)
				err := res.Scan(&id, &name)
				Expect(err).NotTo(HaveOccurred())
				fmt.Println(id, name)
			}
		})
	})
})
