package main

import (
	"fmt"
	"strconv"
)

func TimeConvert(num int) string{
	return strconv.Itoa(num / 60) + ":" + strconv.Itoa(num % 60) 
}


func main(){
	fmt.Println(TimeConvert(126))
}
