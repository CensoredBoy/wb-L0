package main

import "github.com/nats-io/stan.go"

func main() {
	sc, _ := stan.Connect("test-cluster", "nsnsn")
	sc.Publish("test", []byte(`{
  "order_uid": "ID_OF_ORDHDFGSFHKJDGFHJDGFJsdsdsdsER",
  "track_number": "TRACK191919191",
  "entry": "WBIL",
  "delivery": {
    "name": "Danila Efremov",
    "phone": "+97213371488",
    "zip": "3342234",
    "city": "KOALA LUMPUR",
    "address": "Pushkino street, Kalatushkino house",
    "region": "Europe",
    "email": "wb@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "RUB",
    "provider": "wbpay",
    "amount": 191919191919,
    "payment_dt": 16379077222,
    "bank": "sber",
    "delivery_cost": 15000,
    "goods_total": 3172,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 1337228,
      "track_number": "TRACK282828",
      "price": 88,
      "rid": "ab4219087a764ae0btest",
      "name": "Maslory",
      "sale": 30,
      "size": "0",
      "total_price": 3170,
      "nm_id": 2389212,
      "brand": "Luis Karton",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "some customer",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2022-09-26T06:22:19Z",
  "oof_shard": "1"
}`))
}

