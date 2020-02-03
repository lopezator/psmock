package psmock

import (
	"context"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

// PubSubMock is a mock PubSub Client/Server
type PubSubMock struct {
	Client *pubsub.Client
	Server *pstest.Server
}

// NewPubSubMock creates a pubsub mock a single topic and subscription, both
// called psmock.
func NewPubSubMock(ctx context.Context) (*PubSubMock, error) {
	// Start a fake server running locally.
	server := pstest.NewServer()
	// Use the connection when creating a pubsub client.
	client, err := pubsub.NewClient(
		ctx,
		"psmock",
		option.WithEndpoint(server.Addr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithInsecure()),
	)
	if err != nil {
		return nil, err
	}
	topic, err := client.CreateTopic(ctx, "psmock")
	if err != nil {
		return nil, err
	}
	_, err = client.CreateSubscription(ctx, "psmock", pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		return nil, err
	}
	return &PubSubMock{
		Server: server,
		Client: client,
	}, nil
}

// Close closes the psmock client and server
func (p *PubSubMock) Close() {
	_ = p.Client.Close()
	_ = p.Server.Close()
}

