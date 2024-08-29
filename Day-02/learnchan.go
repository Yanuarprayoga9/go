// package main

// import (
// 	"fmt"
// 	"time"
// )

// // Order represents a customer order
// type Orderr struct {
// 	ID     int
// 	Status string
// }
// func processOrder(orderId int, statusChannel chan Orderr) {
// 	time.Sleep(time.Second * 2)
// 	statusChannel <- Orderr{ID:orderId,Status: "Completed"};
// }
// func main() {
// 	statusChannel := make(chan Orderr);

// 	for i := 1 ; i <=5 ; i++ {
// 		go processOrder(i,statusChannel)
// 	}

// }
