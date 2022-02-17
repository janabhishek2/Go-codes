package main
import (
	"fmt"
	"strconv"
	"sort"
)
func main(){

	var slice=make([]int,3,3)
	var input string
	for {
		fmt.Printf("Enter input : ")
		_,err:=fmt.Scan(&input)
		if(err!=nil){
			fmt.Print("Error in IO , exiting...")
			return
		}
		if(len(input)==1 && string(input[0])=="X"){
			fmt.Printf("Final slice : %v\n",slice)
			break
		}else{
			ans,err:=strconv.Atoi(input)
			if(err!=nil){
				fmt.Println("Error found , please enter right characters")
				break
			}
			slice=append(slice,ans)
			sort.Ints(slice)
			fmt.Printf("%v\n",slice)
		}
	}
}