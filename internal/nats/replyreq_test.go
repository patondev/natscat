package nats

import (
    "testing"
    "time"
    "syscall"
)

func TestRequestReply(t *testing.T) {
	// To make sure request and reply is running
	psubject := "subject_"+randString(15)
	preply := "reply_"+randString(25)
	go func() {
		nreply := NatsClass{DefaultURL: natsURLForTest, ReplySubject: psubject, ReplyMessage: preply, QueueGroupName: "qgroup"}
		nreply.Reply()
	}()
	time.Sleep(1 * time.Second)
	reqmsg := "request_"+randString(20)
	nreq := NatsClass{DefaultURL: natsURLForTest, RequestSubject: psubject, RequestMessage: reqmsg}
	nreq.Request()
	time.Sleep(2 * time.Second)
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP) // kill that
}

