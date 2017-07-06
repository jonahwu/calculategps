package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

var bb string
var a string
var cc string

func FromStringToList(locx string) []float64 {
	c := make(map[string]interface{})
	e := json.Unmarshal([]byte(locx), &c)
	if e != nil {
		panic(e)
	}

	k := make([]string, len(c))
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	// output result to STDOUT
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

func main() {
	bb = `{"a":"b", "c":"d"}`
	//cc = `[{"a":"b", "c":"d"}, {"e":"f", "g":"h"}]`
	cc = `[{"a":"b", "c":"d"}, {"e":"f", "g":"h"}]`

	a = `[{"metric":"rrrrealgps","tags":{"loc":"x","id":"jonah"},"aggregateTags":[],"dps":{"1435716527":23.12339973449707,"1435716627":23.33329963684082,"1435717000":23.33329963684082}},{"metric":"rrrrealgps","tags":{"loc":"y","id":"jonah"},"aggregateTags":[],"dps":{"1435716527":123.22219848632812,"1435716627":123.6666030883789,"1435717000":123.6666030883789}}]`
	//a := "{\"a\":\"b\", \"c\":\"d\"}"
	//fmt.Println(a)
	//	fmt.Println(bb)
	//	fmt.Println(gjson.Get(bb, "a"))
	fmt.Println(gjson.Get(cc, "0.a"))
	fmt.Println(gjson.Get(cc, "1.e"))
	locx := gjson.Get(a, "0.dps").String()
	locy := gjson.Get(a, "1.dps").String()
	//locy := gjson.Get(a, "1.dps").String()
	posx := FromStringToList(locx)
	posy := FromStringToList(locy)
	for i := 0; i < len(posx); i++ {
		fmt.Println(posx[i])
	}
	for i := 0; i < len(posy); i++ {
		fmt.Println(posy[i])
	}

}
