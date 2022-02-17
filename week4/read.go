package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
type FullName struct{
	fname string
	lname string
}
func main() {

	file, err := os.Open("names.txt")

	if err != nil {
		log.Fatalf("failed to open")
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	structArr:=make([]FullName,len(text))
	
	file.Close()

	
	for j, each_ln := range text {
		var lineBreak int
		for i:=0;i<len(each_ln);i++{
			if string(each_ln[i])==" "{
				lineBreak=i
				break
			}
		}
		fname:=each_ln[0:lineBreak]
		lname:=each_ln[lineBreak+1:len(each_ln)]
		structArr[j].fname=fname
		structArr[j].lname=lname
	}
	for i:=0;i<len(structArr);i++{
		fmt.Printf("First Name : %v , Last Name : %v\n",structArr[i].fname,structArr[i].lname)
	}
}
