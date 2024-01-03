package main

import (
	"booking-room-app/delivery"
)

func main() {
	delivery.NewServer().Run()
}
