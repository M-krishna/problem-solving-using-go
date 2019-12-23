package main
import (
  "fmt"
  "strings"
  "sort"
)

func AlphabetSoup(str string) string {
  
  // code goes here  
  letters := "abcdefghijklmnopqrstuvwxyz"
  newstr := ""
  var a []int
  for _, v := range str{
    i := strings.Index(letters, string(v))
    a = append(a, i)
  }
  sort.Ints(a)
  for _, v := range a {
    newstr += string(letters[v])
  }
  return newstr;

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(AlphabetSoup("hello")) // expected output ehllo

}