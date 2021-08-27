package main

import "strconv"

func printencoding (s string , ans string , encoding[] string){


	ch1 := s[0];
	var ch2 byte
	if len(s) > 1{
		ch2 = s[1]
	}

	if ch1 == '2' || ch1 == '1'{
		val1,_ := strconv.Atoi(string(ch1))
		val2,_ := strconv.Atoi(string(ch1)+string(ch2))
		printencoding(s[1:] , ans + encoding[val1-1] , encoding)
		printencoding(s[2:] , ans + encoding[val1-1] , encoding)
	}

}