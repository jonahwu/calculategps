package main

//1499299239
import (
	"fmt"
	"time"
)

func main() {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	fmt.Println(timestamp)
	fmt.Println(tm)
	fmt.Println(tm.Format("2006-01-02 03:04:05 PM"))
	fmt.Println(time.Date(2017, 7, 6, 12, 5, 30, 0, time.UTC).Unix())

}
