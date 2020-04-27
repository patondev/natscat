package nats

import (
	"log"
	"github.com/nats-io/nats.go"
)

func (n NatsClass) Publish() {
	if len(n.DefaultURL) == 0 {
		log.Fatalf("[ERROR] Please set the NATS URL\n")
	} else {
		opts := []nats.Option{nats.Name("NATS Sample Publisher")}
		if (len(n.UserCreds) > 0) {
			opts = append(opts, nats.UserCredentials(n.UserCreds))
		}
		// Connect to NATS
		nc, err := nats.Connect(n.DefaultURL, opts...)
		if err != nil {
			log.Fatal(err)
		}
		defer nc.Close()

		if (len(n.PubSubject) == 0) || (len(n.Message) == 0) {
			log.Fatalf("[ERROR] Please set the Subject and Message to Publish\n")
		}

		subj, msg := n.PubSubject, []byte(n.Message)

		if len(n.ReplySubject) > 0 {
			nc.PublishRequest(subj, n.ReplySubject, msg)
		} else {
			nc.Publish(subj, msg)
		}

		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subj, msg)
		}
	}
}