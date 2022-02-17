package main
import "fmt"
func main(){

	var userInputFloat float64
	_,err:=fmt.Scan(&userInputFloat)
	if(err!=nil){
		return
	}
	var truncatedInt int32
	truncatedInt=int32(userInputFloat)
	fmt.Printf("%v",truncatedInt)
}