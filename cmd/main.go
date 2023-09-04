package main

import (
	"L0_Task/pkg/db"
	"L0_Task/pkg/controllers"
	"L0_Task/pkg/nats"
	"log"
	"net/http"
)
func main() {

	go nats.Stan()
	orderRepo := db.New()
	orderRepo.SelectFromDb()
	regi := controllers.Reg{orderRepo}
	nats.TakeMessage("test", "test", &regi)
	http.HandleFunc("/order", regi.OrderReg)

	log.Println("Server started on http://127.0.0.1:8080")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)

}
