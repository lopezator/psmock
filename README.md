# psmock

psmock is a Pub/Sub client/server mock implementation, with a pre-created topic and subscription pair, both called `psmock`.

Is possible to create new ones using `psmock.Client.CreateTopic` and `psmock.Server.CreateSubscription` functions. 

## Usage

Just call `NewPubSubMock`, and then use the `psmock.Client` and `psmock.Server` as you would do with a regular `PubSub` client and server.
 
No configuration or any credentials required.

```go
psmock, err := NewPubSubMock(context.Background())
if err != nil {
    // do something with the err
}
defer pubsubmock.Close()
```

For more details, see [psmock.go](psmock.go), open an issue or refer to the [pstest godoc](https://godoc.org/cloud.google.com/go/pubsub/pstest). 



