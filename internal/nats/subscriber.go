package nats

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to:%s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Printf("Exiting: %v", nc.LastError())
	}))
	return opts
}

func (n NatsClass) SubscribeListen() {
	if len(n.DefaultURL) == 0 {
		log.Printf("[ERROR] Please set the NATS URL\n")
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