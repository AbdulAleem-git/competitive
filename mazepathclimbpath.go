package main

import (
	"fmt"
	"strconv"
)
func mazepathclimb(sr , sc , dr , dc int) [] string{
	var path [] string
	if sr == dr && sc == dc {
		return []string{""}
	}
	for i := 1 ; i <= dc - sc; i++{	
		hpath := mazepathclimb(sr , sc + i , dr , dc)
		for _ , h := range hpath{
			path = append(path, "h"+strconv.Itoa(i)+h)
		}
	}

	for i := 1 ; i <= dr - sr; i++{
		vpath := mazepathclimb(sr + i , sc  , dr , dc)
		for _ , v := range vpath{
			path = append(path, "v"+strconv.Itoa(i)+v)
		}
	}

	for i := 1 ; i <= dc - sc && i <= dr - sr; i++{
		dpath := mazepathclimb(sr + i , sc + i  , dr , dc)
		for _ , d:= range dpath{
			path = append(path, "d"+strconv.Itoa(i)+d)
		}
	}
	return path
}
func main(){
	fmt.Println(mazepathclimb(1,1,3,3))
}