package build

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	cqueues "github.com/pip-services3-go/pip-services3-messaging-go/queues"
	queues "github.com/pip-services3-go/pip-services3-rabbitmq-go/queues"
)

type RabbitMQMessageQueueFactory struct {
	cbuild.Factory
	config     *cconf.ConfigParams
	references cref.IReferences
}

func NewRabbitMQMessageQueueFactory() *RabbitMQMessageQueueFactory {
	c := RabbitMQMessageQueueFactory{}
	c.Factory = *cbuild.NewFactory()

	memoryQueueDescriptor := cref.NewDescriptor("pip-services3-rabbitmq", "message-queue", "rabbitmq", "*", "*")

	c.Register(memoryQueueDescriptor, func(locator interface{}) interface{} {
		name := ""
		descriptor, ok := locator.(*cref.Descriptor)
		if ok {
			name = descriptor.Name()
		}

		return queues.NewEmptyRabbitMQMessageQueue(name)
	})
	return &c
}

func (c *RabbitMQMessageQueueFactory) Configure(config *cconf.ConfigParams) {
	c.config = config
}

func (c *RabbitMQMessageQueueFactory) SetReferences(references cref.IReferences) {
	c.references = references
}

// Creates a message queue component and assigns its name.
//
// Parameters:
//   - name: a name of the created message queue.
func (c *RabbitMQMessageQueueFactory) CreateQueue(name string) cqueues.IMessageQueue {
	queue := queues.NewEmptyRabbitMQMessageQueue(name)

	if c.config != nil {
		queue.Configure(c.config)
	}
	if c.references != nil {
		queue.SetReferences(c.references)
	}

	return queue
}
