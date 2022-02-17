package main

import(
				"strings"
	"fmt"
  "bufio"
  "os"
)
func containsIAN (str string) (bool){
	if(len(str)<3){
		return false
	}
	
	if(!(str[0]=='i' || str[0]=='I')){
		return false
	}
	if(!(str[len(str)-1]=='n' || str[len(str)-1]=='N' )){
		return false
	}
	var subStr=str[1:len(str)-1]
	if(!(strings.Contains(subStr,"a") || strings.Contains(subStr,"A"))){
		return false
	}
	return true
}  
func main(){


	fmt.Printf("Enter the string : ")
	var str string
	 scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    str = scanner.Text()

	
	var out =containsIAN(str)
	if(out){
		fmt.Printf("Found!")
	}else{
		fmt.Printf("Not Found!")
	}
}