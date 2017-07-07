package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	//	"strconv"
)

//km return we might return m so we multiply by 1000
func CalDistance(long1 float64, lati1 float64, long2 float64, lati2 float64) float64 {
	d2r := 0.0174532925199433
	dlong := (long2 - long1) * d2r
	dlat := (lati2 - lati1) * d2r
	a := math.Pow(math.Sin(dlat/2.0), 2) + math.Cos(lati1*d2r)*math.Cos(lati2*d2r)*math.Pow(math.Sin(dlong/2.0), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := 6373 * c
	//return meter
	return d * 1000.

}

func main() {
	csvFile, _ := os.Open("./camera.csv")
	reader := csv.NewReader(csvFile)
	//reader.FieldsPerRecord = -1
	csvData, _ := reader.ReadAll()
	//[5]is long  [6] is lati
	for _, each := range csvData {
		fmt.Println(each[0], each[1], each[2])
		fmt.Println(each[5], each[6], each[7], each[8])
		fmt.Println("------------------------------------")
		//allRecords = append(allRecords, oneRecord)
	}
	fmt.Println("haha")
	lati1 := 45.527517
	long1 := -122.718766
	lati2 := 45.373373
	long2 := -121.693604
	dist := CalDistance(long1, lati1, long2, lati2)
	fmt.Println(dist)

	testlati := 25.065805435180664
	testlong := 121.67511749267578
	tarlati := 25.06611442565918
	tarlong := 121.67586517333984
	fmt.Println(CalDistance(testlong, testlati, tarlong, tarlati))
}
