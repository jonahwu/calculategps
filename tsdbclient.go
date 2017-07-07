package main

import (
	"fmt"
	//"net/http"
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"math"
	"os/exec"
	"sort"
	"strconv"
	"time"
)

var tmpgettsdb string

func FromStringToList(locx string) ([]float64, []int) {
	c := make(map[string]interface{})
	e := json.Unmarshal([]byte(locx), &c)
	if e != nil {
		panic(e)
	}

	k := make([]int, len(c))
	i := 0
	fmt.Println("unmarshal------------- ")
	fmt.Println(c)
	// copy c's keys into k
	for s, _ := range c {
		k[i], _ = strconv.Atoi(s)
		i++
	}
	sort.Ints(k)

	// output result to STDOUT
	fmt.Println("print timestamp key: is it sequential by time?")
	fmt.Printf("%#v\n", k)
	var pos []float64
	//	for _, kk := range k {
	for kk := 0; kk < len(k); kk++ {
		skk := strconv.Itoa(k[kk])

		fv := float64(c[skk].(float64))
		//fmt.Println(fv)
		fmt.Println(fv)
		pos = append(pos, fv)
	}
	return pos, k

}

func FromStringToListOld(locx string) []float64 {
	c := make(map[string]interface{})
	e := json.Unmarshal([]byte(locx), &c)
	if e != nil {
		panic(e)
	}

	k := make([]string, len(c))
	i := 0
	fmt.Println("unmarshal------------- ")
	fmt.Println(c)
	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	// output result to STDOUT
	fmt.Println("print timestamp key: is it sequential by time?")
	fmt.Printf("%#v\n", k)
	var pos []float64
	for _, kk := range k {
		fv := float64(c[kk].(float64))
		//fmt.Println(fv)
		fmt.Println(fv)
		pos = append(pos, fv)
	}
	return pos

}

func CurlCommand(remoteip string, remoteport string, tmpgettsdb string) string {
	sshcmd1 := "-k"
	sshcmd2 := "-H"
	sshcmd3 := "\"Content-Type: application/json; charset=UTF-8\""
	//	sshcmd4 := "https://172.18.0.1:443/apis/apps/v1beta1/namespaces/inq/statefulsets"
	sshcmd4 := "-d"
	sshcmd5 := tmpgettsdb
	sshcmd6 := fmt.Sprintf("http://%s:%s/api/query", remoteip, remoteport)

	runcommand := exec.Command("curl", sshcmd1, sshcmd2, sshcmd3, sshcmd4, sshcmd5, sshcmd6)
	fmt.Println(runcommand.Args)
	var out bytes.Buffer
	runcommand.Stdout = &out

	err := runcommand.Run()
	fmt.Println(out.String())
	//value := gjson.Get(out.String(), "items.0.metadata")
	fmt.Println(err)
	return out.String()
}

func GetTSDBData() ([]float64, []float64, []int) {
	tmpgettsdb = `{
    "start": 1435716526,
    "queries": [
        {
            "metric": "testgps",
            "aggregator": "avg",
            "tags": {
                "id": "*0707*",
                "loc":"*"
                           }        }
    ]}`
	remoteip := "35.189.170.202"
	remoteport := "14242"
	loc := CurlCommand(remoteip, remoteport, tmpgettsdb)
	fmt.Println("initial raw data-------------")
	fmt.Println(loc)
	locx := gjson.Get(loc, "0.dps").String()
	locy := gjson.Get(loc, "1.dps").String()
	posx, ts := FromStringToList(locx)
	posy, ts := FromStringToList(locy)
	fmt.Println("result ---------------------")
	fmt.Println(posx)
	fmt.Println(posy)
	fmt.Println(ts)
	if len(posx) != len(posy) {
		fmt.Println("length of two data is different")
	} else {
		fmt.Println("length of two data is the same")
	}
	return posx, posy, ts
}

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
func CalGPSDistance(posx []float64, posy []float64, ts []int) {
	tarx := 25.080223
	tary := 121.697908
	fmt.Println("---------------------------------")
	for i := 0; i < len(posx); i++ {
		gpsx := posx[i]
		gpsy := posy[i]
		tss := ts[i]
		dist := CalDistance(gpsy, gpsx, tary, tarx)
		fmt.Println(dist, tss, gpsx, gpsy, tarx, tary)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	//25.080223, 121.697908
	posx, posy, ts := GetTSDBData()
	fmt.Println(posx)
	fmt.Println(posy)
	CalGPSDistance(posx, posy, ts)

}
