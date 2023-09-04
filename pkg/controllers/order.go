package controllers

import (
	"L0_Task/pkg/db"
	"html/template"
	"log"
	"net/http"
)


type Reg struct {
	Order db.DB
}

func (re *Reg) OrderReg(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	order, _ := re.Order.FindById(id)
	tmpl, _ := template.ParseFiles("../static/order.html")
	err := tmpl.Execute(w, order)
	if err != nil {
		log.Fatalf("Error while executing template: %v", err)
	}

}
