package main

import "fmt"

type ConsumerMessage struct {
	Headers []*RecordHeader
}

type RecordHeader struct {
	Key   string
	Value []byte
}

func main() {
	test := RecordHeader{
		Key: "1",
	}

	//test := RecordHeader{}

	var msg *ConsumerMessage
	msg = &ConsumerMessage{
		Headers: []*RecordHeader{&test},
	}

	fmt.Println(msg.Headers[0].Key, msg.Headers[0].Value)
}
