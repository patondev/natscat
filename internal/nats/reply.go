package nats

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func (n NatsClass) Reply() {
	if len(n.DefaultURL) == 0 {
		log.Fatalf("[ERROR] Please set the NATS URL\n")
	} else {
		opts := []nats.Option{nats.Name("NATS Sample Responder")}
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

		if (len(n.ReplySubject) == 0) || (len(n.ReplyMessage) == 0) {
			log.Fatalf("[ERROR] Please set Reply subject and message\n")
		}

		subj, reply, i := n.ReplySubject, n.ReplyMessage, 0

		if (len(n.QueueGroupName) == 0){
			log.Fatalf("[ERROR] Please set QueueGroupName\n")
		}

		nc.QueueSubscribe(subj, n.QueueGroupName, func(msg *nats.Msg) {
			i++
			printMsg(msg, i)
			msg.Respond([]byte(reply))
		})
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening on [%s]", subj)

		// Kill signal handler
		sigs := make(chan os.Signal, 1)
	    signal.Notify(sigs,syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT)
	    //kill -SIGHUP XXXX

	    <-sigs
	    log.Printf("Draining...")
		nc.Drain()
	}
}