package main

import (
	"fmt"
	geoindex "github.com/hailocab/go-geoindex"
)

type Driver struct {
	id            string
	lat           float64
	lon           float64
	canAcceptJobs bool
}

//http://www.gpsvisualizer.com/calculators
func main() {
	index := geoindex.NewPointsIndex(geoindex.Km(0.5))
	//	geoindex.NewGeoPoint("id", 51.51, -0.11)
	//	index.Add(&Driver{id: "id1", lat: 23.333, lon: 123.3333, canAcceptJobs: true})
	//index.Add(geoindex.NewGeoPoint("id", 23.333, 123.333))
	index.Add(geoindex.NewGeoPoint("id", 23.333, 123.333))  //301m
	index.Add(geoindex.NewGeoPoint("id1", 23.335, 123.355)) // 2.045km
	index.Add(geoindex.NewGeoPoint("id2", 23.336, 123.356)) //2.15km
	index.Add(geoindex.NewGeoPoint("id3", 23.536, 123.556)) //31km
	//a := index.KNearest(geoindex.NewGeoPoint("ida", 23.335, 123.335), 5, geoindex.Km(5), func(p geoindex.Point) bool { return true })
	a := index.KNearest(geoindex.NewGeoPoint("ida", 23.335, 123.335), 5, geoindex.Km(30), func(p geoindex.Point) bool { return true })
	fmt.Println(a)

}
