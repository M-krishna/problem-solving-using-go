package main
import "fmt"

func SimpleAdding(num int) int {
  
  // code goes here  
  sum := 0
	for i := 1; i <= num; i++{
		sum = sum + i
	}
	return sum

}

func main() {
  fmt.Println(SimpleAdding(10))

}