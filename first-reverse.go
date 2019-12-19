package main
import (
  "fmt"
  "strings"
)

func FirstReverse(str string) string {
  
  // code goes here
	s := []rune(str)
	var n []string
	strlen := len(s)
	for i := strlen-1; i >= 0; i--{
		n = append(n, string(s[i]))
	} 
  return strings.Join(n, "");

}

func main() {

  
  fmt.Println(FirstReverse("krishna"))

}