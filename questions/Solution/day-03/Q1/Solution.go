package main

import "fmt"

func main() {
	x := make([]int, 0, 1)

	prevCap := 1
	count := 0
	for i:=0; i<1000000; i++{
		x = append(x , i)
		curCap:=cap(x)
		if(prevCap != curCap){
			percentageCh := (float64(curCap)-float64(prevCap)) / float64(prevCap) * 100
			fmt.Println("change in capacity : ",percentageCh)
			count = count+1
			prevCap = curCap
		}
	}
	fmt.Println(count)

}


