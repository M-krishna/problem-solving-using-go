package main
import (
  "fmt"
  "unicode"
  "strings"
)

func LetterChanges(str string) string {
  
  // code goes here
  alphabets := "abcdefghijklmnopqrstuvwxyz"
	var newstr = ""
	
	for _, e := range str{
		index := strings.Index(alphabets, strings.ToLower(string(e)))
		if unicode.IsNumber(e) || unicode.IsPunct(e){
			newstr += string(e)
		}
		if unicode.IsSpace(e){
			newstr += " "
		}
		if index != 25 && index != -1 {
			s := string(alphabets[index+1])
			if isVowel(s){
				a := strings.ToUpper(s)
				newstr += a
			} else { 
				newstr += s 
			}
			
		} else if index == 25{
			newstr += "A"
		}
	}
	return newstr
}

func isVowel(str string) bool {
	if str == "a" || str == "e" || str == "i" || str == "o" || str == "u" {
		return true
	}
	return false
}

func main() {
  fmt.Println(LetterChanges("hello world"))
}