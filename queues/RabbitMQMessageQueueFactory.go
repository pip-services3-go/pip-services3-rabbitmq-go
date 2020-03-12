package queues

import (
	cconf "github.com/pip-services3-go/pip-services3-commons-go/config"
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type RabbitMQMessageQueueFactory struct {
	*cbuild.Factory
	Descriptor            *cref.Descriptor
	MemoryQueueDescriptor *cref.Descriptor

	config *cconf.ConfigParams
}

func NewRabbitMQMessageQueueFactory() *RabbitMQMessageQueueFactory {
	c := RabbitMQMessageQueueFactory{}
	c.Factory = cbuild.NewFactory()
	c.Descriptor = cref.NewDescriptor("pip-services3-rabbitmq", "factory", "message-queue", "rabbitmq", "1.0")
	c.MemoryQueueDescriptor = cref.NewDescriptor("pip-services3-rabbitmq", "message-queue", "rabbitmq", "*", "*")

	c.Register(MemoryQueueDescriptor, func(locator interface{}) interface{} {
		descriptor := locator.(cref.Descriptor)
		queue := NewRabbitMQMessageQueue(descriptor.Name())
		if c.config != nil {
			queue.Configure(config)
		}
		return queue
	})
}

func (c *RabbitMQMessageQueueFactory) Configure(config *cconf.ConfigParams) {
	c.config = config
}
