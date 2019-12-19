package main
import (
  "fmt"
  "regexp"
  "strings"
)

func LongestWord(sen string) string {
  
  // code goes here  
  r := regexp.MustCompile("[^a-zA-Z0-9]+")
	str := r.ReplaceAllString(sen, " ")
	splitted_string := strings.Split(str, " ")
	var word = ""
	for _, e := range splitted_string{
		if len(word) < len(e) {
			word = e
		}
	}
	return word

}

func main() {
  fmt.Println(LongestWord("I&&!! Love Dogs"))
}