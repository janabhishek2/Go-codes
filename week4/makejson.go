package main
import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
)
func main(){
	var name string
	var address string
	scanner := bufio.NewScanner(os.Stdin)
   
	sampleMap:=make(map[string]string)
	fmt.Print("Enter name : ")
	 scanner.Scan() // use `for scanner.Scan()` to keep reading
    name = scanner.Text()
	
	fmt.Print("\nEnter address : ")
	 scanner.Scan() // use `for scanner.Scan()` to keep reading
    address = scanner.Text()
	sampleMap["Name"]=name
	sampleMap["Address"]=address
	barr,err:=json.Marshal(sampleMap)
	if(err!=nil){
		return
	}
	fmt.Println(string(barr))
}