package psmock

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
)

func TestPsMock(t *testing.T) {
	ctx := context.Background()
	psmock, err := NewPubSubMock(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer psmock.Close()
	_ = psmock.Server.Publish("projects/psmock/topics/psmock", []byte("I'm using psmock mommy!"), nil)

	errors := make(chan error)
	messages := make(chan *pubsub.Message)
	go func() {
		errors <- psmock.Client.Subscription("psmock").Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
			messages <- message
			message.Ack()
		})
	}()

	select {
	case err := <-errors:
		t.Fatal(err)
	case message := <-messages:
		want := "I'm using psmock mommy!"
		got := string(message.Data)
		if got != want {
			t.Errorf("\n got: %v \n want: %v", got, want)
		}
	}
}
