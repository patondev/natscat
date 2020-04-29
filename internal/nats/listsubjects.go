package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type SubobjStruct struct {
	Subject string `json:"subject"`
	QGroup  string `json:"qgroup"`
	SID	string `json:"sid"`
	Msgs	int    `json:"msgs"`
	CID	int    `json:"cid"`
}

type SubsStruct struct {
	Total int `json:"total"`
	Subscriptions []SubobjStruct `json:"subscriptions_list"`
}

func (n NatsClass) ListSubjects() {
	var text string

	response, err := http.Get(n.DefaultURL)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		text = string(data)
	}

	textBytes := []byte(text)
	subs := SubsStruct{}
	err = json.Unmarshal(textBytes, &subs)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Total: %v\n",subs.Total)
	if (subs.Total == 0){
		fmt.Println("No online subjects available")
	} else {
		fmt.Printf("Online Subjects: ")
		for _, v := range subs.Subscriptions {
			fmt.Printf("%v  ",v.Subject)
		}
		fmt.Println()	
	}
}