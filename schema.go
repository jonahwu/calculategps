package main

import "fmt"

type Gps2dLoc struct {
	Long float64
	Lati float64
}

type Gps4dLoc struct {
	Long      float64
	Lati      float64
	Timestamp int
}

type UserGpsInfo struct {
	CurrentLoc      Gps4dLoc
	CameraLoc       []Gps2dLoc
	TargetCameraLoc []Gps4dLoc
	PrevLoc         Gps4dLoc
	createLoc       Gps4dLoc
}

func main() {
	a := Gps4dLoc{}
	a.Lati = 23.333
	a.Long = 123.3333

	ugi := UserGpsInfo{}
	ugi.CurrentLoc.Lati = 12.3
	ugi.CurrentLoc.Long = 123.333

}
