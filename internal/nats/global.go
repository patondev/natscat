package nats

import (
	"log"
	"math/rand"
	"github.com/nats-io/nats.go"
	"time"
)

type NatsClass struct {
    DefaultURL string
    UserCreds  string
    ReplySubject string
    ReplyMessage string
    RequestSubject string
    RequestMessage string
    QueueGroupName string
    PubSubject string
    Message string
    SubsSubject string
}

var natsURLForTest = "nats://demo.nats.io" // you can change this
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		if (err != nil){
			log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
		}
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		if (nc.LastError() != nil){
			log.Printf("Exiting: %v", nc.LastError())
		}
	}))
	return opts
}

func randString(length int) string {
	const charset string = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}