package nats

import (
	"log"
    "testing"
    "math/rand"
    "time"
    //"github.com/nats-io/nats.go"
    "syscall"
)

var natsURL = "nats://127.0.0.1:4222" // you can change this
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func randString(length int) string {
	const charset string = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func TestPubSub(t *testing.T) {
	psubject := "subject_"+randString(5)
	mpub := "message_"+randString(10)

	go func(){
		time.Sleep(2 * time.Second)
		log.Printf("Subject %v with Message %v created!\n",psubject, mpub)
		npub := NatsClass{DefaultURL: natsURL, PubSubject: psubject, Message: mpub}
		npub.Publish()
		time.Sleep(2 * time.Second)
		log.Printf("Subject %v with Message %v created!\n",psubject, mpub)
		npub := NatsClass{DefaultURL: natsURL, PubSubject: psubject, Message: mpub}
		npub.Publish()
		time.Sleep(2 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP) // kill that
	}()

	nsub := NatsClass{DefaultURL: natsURL, SubsSubject: psubject}
	nsub.SubscribeListen()
}