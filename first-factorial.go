package main
import "fmt"

func FirstFactorial(num int) int {
  
  // code goes here
  f := 1
  for num > 0{
    f = f * num
    num--
  }
  return f;

}

func main() {
  fmt.Println(FirstFactorial(5))

}