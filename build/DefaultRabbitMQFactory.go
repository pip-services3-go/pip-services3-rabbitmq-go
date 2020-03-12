package build

import (
	cref "github.com/pip-services3-go/pip-services3-commons-go/refer"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
	rabque "github.com/pip-services3-go/pip-services3-rabbitmq-go/queues"
)

// Creates RabbitMQMessageQueue components by their descriptors.
// See RabbitMQMessageQueue
type DefaultRabbitMQFactory struct {
	*cbuild.Factory
	Descriptor                             *cref.Descriptor
	Descriptor3                            *cref.Descriptor
	RabbitMQMessageQueueFactoryDescriptor  *cref.Descriptor
	RabbitMQMessageQueueFactory3Descriptor *cref.Descriptor
	RabbitMQMessageQueueDescriptor         *cref.Descriptor
	RabbitMQMessageQueue3Descriptor        *cref.Descriptor
}

// NewDefaultRabbitMQFactory method are create a new instance of the factory.
func NewDefaultRabbitMQFactory() *DefaultRabbitMQFactory {
	c := DefaultRabbitMQFactory{}
	c.Factory = cbuild.NewFactory()
	c.Descriptor = cref.NewDescriptor("pip-services", "factory", "rabbitmq", "default", "1.0")
	c.Descriptor3 = cref.NewDescriptor("pip-services3", "factory", "rabbitmq", "default", "1.0")
	c.RabbitMQMessageQueueFactoryDescriptor = cref.NewDescriptor("pip-services", "factory", "message-queue", "rabbitmq", "1.0")
	c.RabbitMQMessageQueueFactory3Descriptor = cref.NewDescriptor("pip-services3", "factory", "message-queue", "rabbitmq", "1.0")
	c.RabbitMQMessageQueueDescriptor = cref.NewDescriptor("pip-services", "message-queue", "rabbitmq", "*", "1.0")
	c.RabbitMQMessageQueue3Descriptor = cref.NewDescriptor("pip-services3", "message-queue", "rabbitmq", "*", "1.0")

	c.RegisterType(c.RabbitMQMessageQueueFactoryDescriptor, rabque.NewRabbitMQMessageQueueFactory)
	c.RegisterType(c.RabbitMQMessageQueueFactory3Descriptor, rabque.NewRabbitMQMessageQueueFactory)
	c.RegisterType(c.RabbitMQMessageQueueDescriptor, rabque.NewRabbitMQMessageQueue)
	c.RegisterType(c.RabbitMQMessageQueue3Descriptor, rabque.NewRabbitMQMessageQueue)
	return &c
}
