package nats

import (
	"log"
	"time"
	"github.com/nats-io/nats.go"
)

func (n NatsClass) Request() {
	if len(n.DefaultURL) == 0 {
		log.Fatalf("[ERROR] Please set the NATS URL\n")
	} else {
		// Connect Options.
		opts := []nats.Option{nats.Name("NATS Sample Requestor")}
		opts = setupConnOptions(opts)

		// Use UserCredentials
		if (len(n.UserCreds) > 0) {
			opts = append(opts, nats.UserCredentials(n.UserCreds))
		}

		nc, err := nats.Connect(n.DefaultURL, opts...)
		if err != nil {
			log.Fatal(err)
		}
		defer nc.Close()

		if (len(n.RequestSubject) == 0) || (len(n.RequestMessage) == 0) {
			log.Fatalf("[ERROR] Please set Reply subject and message\n")
		}

		subj, payload := n.RequestSubject, []byte(n.RequestMessage)
		msg, err := nc.Request(subj, payload, 2 * time.Second)
		if err != nil {
			if nc.LastError() != nil {
				log.Fatalf("%v for request", nc.LastError())
			}
			log.Fatalf("%v for request", err)
		}

		log.Printf("Published [%s] : '%s'", subj, payload)
		log.Printf("Received  [%v] : '%s'", msg.Subject, string(msg.Data))
	}
}