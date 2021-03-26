package test_queues

import (
	"os"
	"testing"

	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	rabbitqueue "github.com/pip-services3-go/pip-services3-rabbitmq-go/queues"
	"github.com/stretchr/testify/assert"
)

func TestRabbitMQMessageQueue(t *testing.T) {

	var queue *rabbitqueue.RabbitMQMessageQueue
	var fixture *MessageQueueFixture

	rabbitmqHost := os.Getenv("RABBITMQ_SERVICE_HOST")
	if rabbitmqHost == "" {
		rabbitmqHost = "localhost"
	}
	rabbitmqPort := os.Getenv("RABBITMQ_SERVICE_PORT")
	if rabbitmqPort == "" {
		rabbitmqPort = "5672"
	}

	rabbitmqExchange := os.Getenv("RABBITMQ_EXCHANGE")
	if rabbitmqExchange == "" {
		rabbitmqExchange = "test"
	}

	rabbitmqQueue := os.Getenv("RABBITMQ_QUEUE")
	if rabbitmqQueue == "" {
		rabbitmqQueue = "test"
	}

	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	if rabbitmqUser == "" {
		rabbitmqUser = "user"
	}

	rabbitmqPassword := os.Getenv("RABBITMQ_PASS")
	if rabbitmqPassword == "" {
		rabbitmqPassword = "password"
	}

	if rabbitmqHost == "" && rabbitmqPort == "" {
		return
	}

	queueConfig := cconf.NewConfigParamsFromTuples(
		"exchange", rabbitmqExchange,
		"queue", rabbitmqQueue,
		"options.auto_create", true,
		//"connection.protocol", "amqp",
		"connection.host", rabbitmqHost,
		"connection.port", rabbitmqPort,
		"credential.username", rabbitmqUser,
		"credential.password", rabbitmqPassword,
	)

	queue = rabbitqueue.NewEmptyRabbitMQMessageQueue("testQueue")
	queue.Configure(queueConfig)

	fixture = NewMessageQueueFixture(queue)

	qOpnErr := queue.Open("")
	if qOpnErr == nil {
		queue.Clear("")
	}

	defer queue.Close("")

	t.Run("RabbitMQMessageQueue:Send Receive Message", fixture.TestSendReceiveMessage)
	err := queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:Receive Send Message", fixture.TestReceiveSendMessage)
	err = queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:Receive And Complete Message", fixture.TestReceiveCompleteMessage)
	err = queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:Receive And Abandon Message", fixture.TestReceiveAbandonMessage)
	err = queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:Send Peek Message", fixture.TestSendPeekMessage)
	err = queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:Peek No Message", fixture.TestPeekNoMessage)
	err = queue.Clear("")
	assert.Nil(t, err)
	t.Run("RabbitMQMessageQueue:On Message", fixture.TestOnMessage)

}
