package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
//	p := fmt.Println
//	p(now.Day())
//	p(now.Year())
//	p(now.Month())
	today := now.Day()
	tyear := now.Year()
	tmonth := now.Month()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ikasi := r1.Intn(100)

	fmt.Print("Henkinen ikäsi on: ",ikasi," ", "tänä kauniina kesäpäivänä ",today,tmonth,tyear,"\n")
}
