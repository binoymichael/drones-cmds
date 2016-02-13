package integrations_test

import (
	"fmt"
	"os"
	"testing"

	. "github.com/cloudnativego/drones-cmds"
	"github.com/streadway/amqp"
)

type fakeMessage struct {
	a string
	b string
}

func TestAmqpDispatcherSubmitsToQueue(t *testing.T) {

	rabbitPw := os.Getenv("WERCKER_RABBITMQ_PASSWORD")
	rabbitUserName := os.Getenv("WERCKER_RABBITMQ_USERNAME")
	rabbitHost := os.Getenv("WERCKER_RABBITMQ_HOST")
	rabbitPort := os.Getenv("WERCKER_RABBITMQ_PORT")
	rabbitURL := "amqp://guest:guest@192.168.99.100:5672/"
	if len(rabbitHost) > 0 {
		rabbitURL = fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUserName, rabbitPw, rabbitHost, rabbitPort)
	}

	conn, err := amqp.Dial(rabbitURL)
	failOnError(t, err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(t, err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(t, err, "Failed to declare a queue")
	dispatcher := NewAMQPDispatcher(ch, q.Name)
	err = dispatcher.DispatchMessage("unused", fakeMessage{a: "hello", b: "world"})
	failOnError(t, err, "Failed to dispatch message on channel/queue")
}

func failOnError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Errorf("%s: %s", msg, err)
	}
}