package producer

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/Shopify/sarama"
)

// TODO: change code as Slava V. wishes

type RpcName = int

const (
	Create RpcName = iota
	Update
	Remove
)

type LogProducer interface {
	Send(rpc RpcName, req, res interface{}, err error) error
}

const (
	KafkaBroker = "127.0.0.1:9094"
	KafkaTopic  = "events"
)

type logProducer struct {
	syncProducer sarama.SyncProducer
}

func New() (LogProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{KafkaBroker}, config)
	if err != nil {
		return nil, err
	}

	return &logProducer{syncProducer: producer}, nil
}

func (lp *logProducer) Send(rpc RpcName, req, res interface{}, err error) error {

	if lp == nil {
		return errors.New("LogProducer is nil")
	}

	message := map[string]interface{}{
		"rpc_name":  rpcNameToString(rpc),
		"request":   req,
		"response":  res,
		"timestamp": time.Now(),
	}

	if err != nil {
		message["error"] = err.Error()
	}

	b, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic:     KafkaTopic,
		Partition: -1,
		Value:     sarama.StringEncoder(b),
	}

	_, _, err = lp.syncProducer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func rpcNameToString(et RpcName) string {

	switch et {
	case Create:
		return "Create"
	case Update:
		return "Update"
	case Remove:
		return "Remove"
	}

	return "undefined rpc name"
}
