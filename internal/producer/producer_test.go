package producer_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozoncp/ocp-classroom-api/internal/producer"
)

var _ = Describe("Producer", func() {

	// It is supposed that Kafka is running on that endpoint
	const (
		broker      = "127.0.0.1:9094"
		wrongBroker = "127.0.0.1:9093"
	)

	var (
		ctx context.Context
	)

	BeforeEach(func() {

		ctx = context.Background()
	})

	Describe("New call", func() {

		It("returns new LogProducer if Kafka is reachable", func() {

			got, err := producer.New(ctx, broker)

			Expect(got).ShouldNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns err if Kafka is not reachable", func() {

			got, err := producer.New(ctx, wrongBroker)

			Expect(got).Should(BeNil())
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Send call", func() {

		var logProducer producer.LogProducer

		BeforeEach(func() {

			var err error
			logProducer, err = producer.New(ctx, broker)
			if err != nil {
				Fail("could not create LogProducer instance")
			}
		})

		It("returns nil if everything is fine", func() {

			simpleReq := "request"
			simpleRes := "response"
			var simpleErr error = nil

			err := logProducer.Send(producer.Created, simpleReq, simpleRes, simpleErr)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns err if json.Marshal returns err", func() {

			// I don't know hot to invoke this case
		})
	})

	// TODO: add tests
})
