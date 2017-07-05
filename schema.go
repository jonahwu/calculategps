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

type ListGps2dLocTim struct {
	Loc       []Gps2dLoc
	Timestamp int
}

type UserGpsInfo struct {
	CurrentLoc      Gps4dLoc
	CameraLoc       ListGps2dLocTim
	TargetCameraLoc ListGps2dLocTim
	// > 5 sec to update
	PrevLoc Gps4dLoc
	// as map filter used
	CreateLoc Gps4dLoc
}

//shall we need CreatedTargetCameraLoc ? Yes I think we need, it had better we define a target range ~500m, over 500m will regenerate new targets
//what's means 500m, it shoud related to car speed to sec

// initial get fake data and timestamp
//filter
// MapDistanceFilter (createLoc and CurrentLoc) distance
// MapTimeFilter (createLoc and CurrentLoc) with time
// Distance Filter (CurrentLoc and CamearaLoc) -> TargetCameraLoc
// approach Filter (TargetCameraLoc and CurrentLoc, and PrevLoc) to make sure it's getting closed
// orientation Filter
// return target, distance, (arround, near, passed)
// UpdateFilter (CurrentLoc to PrevLoc if >5sec), in the beginning we just check Prev is zero and put current loc there.

func main() {
	a := Gps4dLoc{}
	a.Lati = 23.333
	a.Long = 123.3333

	ugi := UserGpsInfo{}
	ugi.CurrentLoc.Lati = 12.3
	ugi.CurrentLoc.Long = 123.333
	testloc := Gps2dLoc{12.3, 22.3}

	if ugi.CameraLoc.Loc == nil {
		fmt.Println("cameraloc is nil")
		fmt.Println(ugi.CameraLoc)
		ugi.CameraLoc.Loc = append(ugi.CameraLoc.Loc, testloc)
		fmt.Println(len(ugi.CameraLoc.Loc))
		fmt.Println("shoelati", ugi.CameraLoc.Loc[0].Lati)
	} else {
		fmt.Println("shoelati", ugi.CameraLoc.Loc[0].Lati)
	}

	if ugi.CreateLoc.Timestamp == 0 {
		fmt.Println("create loc is nil")
	}

}
