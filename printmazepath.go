package main

import (
	"fmt"
	"strconv"
)

func printmp(sr , sc , dr , dc int , ans string) {

	if  sr == dr && sc == dc {
		fmt.Println(ans)
		return
	}else if sr > dr || sc > dc {
		return
	}
	for i := 1 ; i <= dr - sr ; i++{
		printmp(sr+i, sc, dr, dc ,ans+"v"+strconv.Itoa(i))
	}
	for i := 1 ; i <= dc - sc ; i++{
		printmp(sr, sc+i, dr, dc ,ans+"h"+strconv.Itoa(i))
	}
	for i := 1 ; i <= dr - sr && i <= dc - sc  ; i++{
		printmp(sr+i, sc+i, dr, dc ,ans+"d"+strconv.Itoa(i))
	}
}
func main() {
	printmp(1,1,3,3,"")
}