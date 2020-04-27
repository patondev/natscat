package nats

import (
	"log"
    "testing"
    "math/rand"
    "time"
    "github.com/nats-io/nats.go"
    "syscall"
    "strings"
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
	// To make sure it receives the message
	psubject := "subject_"+randString(5)
	mpub := "message_"+randString(10)

	go func(){
		time.Sleep(2 * time.Second)
		log.Printf("Subject %v with Message %v created!\n",psubject, mpub)
		npub := NatsClass{DefaultURL: natsURL, PubSubject: psubject, Message: mpub}
		npub.Publish()
		time.Sleep(2 * time.Second)
		log.Printf("Subject %v with Message %v created!\n",psubject, mpub)
		npub = NatsClass{DefaultURL: natsURL, PubSubject: psubject, Message: mpub}
		npub.Publish()
		time.Sleep(2 * time.Second)
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP) // kill that
	}()

	nsub := NatsClass{DefaultURL: natsURL, SubsSubject: psubject}
	nsub.SubscribeListen()
}

func TestPubWithVerification(t *testing.T) {
	// To make sure that the message is the same as when it's being sent
	psubject := "subject_"+randString(15)
	mpub := "message_"+randString(25)

	go func(){
		time.Sleep(2 * time.Second)
		log.Printf("Subject %v with Message %v created!\n",psubject, mpub)
		npub := NatsClass{DefaultURL: natsURL, PubSubject: psubject, Message: mpub}
		npub.Publish()
	}()

	nc, _ := nats.Connect(natsURL)
	defer nc.Close()

	sub, _ := nc.SubscribeSync(psubject)
	m, err := sub.NextMsg(3 * time.Second)
	if err == nil {
		recvmsg := string(m.Data)
		if (strings.Compare(recvmsg,mpub) != 0){
			t.Errorf("The string received is not the same, should be %v but got %v",mpub,recvmsg)
		}
	} else {
	    t.Errorf("NextMsg timed out.")
	}
}