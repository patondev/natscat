package nats

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/nats.go"
)

func (n NatsClass) SubscribeListen() {
	if len(n.DefaultURL) == 0 {
		log.Fatalf("[ERROR] Please set the NATS URL\n")
	} else {
		opts := []nats.Option{nats.Name("NATS Sample Publisher")}
		opts = setupConnOptions(opts)
		if (len(n.UserCreds) > 0) {
			opts = append(opts, nats.UserCredentials(n.UserCreds))
		}
		// Connect to NATS
		nc, err := nats.Connect(n.DefaultURL, opts...)
		if err != nil {
			log.Fatal(err)
		}
		defer nc.Close()

		if (len(n.SubsSubject) == 0){
			log.Fatalf("[ERROR] Please set the Subject to Subscribe\n")
		}

		subj, i := n.SubsSubject, 0

		nc.Subscribe(subj, func(msg *nats.Msg) {
			i += 1
			printMsg(msg, i)
		})
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Println(err)
		}

		log.Printf("Listening on [%s]", subj)

		// Kill signal handler
		sigs := make(chan os.Signal, 1)
	    signal.Notify(sigs,syscall.SIGHUP,syscall.SIGINT,syscall.SIGTERM,syscall.SIGQUIT)
	    //kill -SIGHUP XXXX

	    <-sigs
	}
}