package nats

import (
	"L0_Task/pkg/model"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"L0_Task/pkg/controllers"
	"log"
)

var Sc stan.Conn

func Stan() {
	ConnectStan("nice-nice")
}

func ConnectStan(clientID string) {
	clusterID := "test-cluster"    // nats cluster id
	url := "nats://127.0.0.1:4222" // nats url
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url),
		stan.Pings(1, 3),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}))
	if err != nil {
		log.Fatalf("Can not connect: %v", err)
	}

	log.Println("Connected!")

	Sc = sc
}

func TakeMessage(subject, queueGroup string, reg *controllers.Reg) {
	mcb := func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Printf("Error while publishing message:%v", err)
		}
		var order model.Order
		err := json.Unmarshal(msg.Data, &order)
		if err != nil {
			log.Printf("Error with json %v", err)
			return
		}
		_, err = reg.Order.Create(&order)
		if err != nil {
			log.Fatalf("Error with data, %v", err)
		}
		log.Println("NEW ORDER, ORDER ID - ", order.OrderUid)
	}

	_, err := Sc.QueueSubscribe(subject,
		queueGroup, mcb,
		stan.SetManualAckMode())
	if err != nil {
		log.Println(err)
	}

}
