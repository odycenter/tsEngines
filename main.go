package main

import (
	"log"
	_ "tsEngine/tsFuzzy"
	_ "tsEngine/tsMap"
	_ "tsEngine/tsSlice"
	_ "tsEngine/tsString"
	"tsEngine/tsTime"
	_ "tsEngine/tsTime"
)

func main() {
	// log.Println(tsTime.MakeTimeSe(int64()))
	log.Println(tsTime.CurrNs())
}
