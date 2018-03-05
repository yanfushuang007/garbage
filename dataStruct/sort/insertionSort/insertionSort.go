package main

import "fmt"

func InsertionSort( data []int) []int{
	var i, j, tmp int

	var count = len(data)

	if(count < 2){
		return data
	}

	for i=1 ; i<count; i++ {
		tmp = data[i]
		for j=i; j>0 && data[j-1] > tmp; j-- {
			data[j] = data[j-1]
			fmt.Println("for2 j=", j, ", i=", i, ", data=", data);
		}
		data[j] = tmp
		fmt.Println("j=", j, ", i=", i, ", data=", data);
	}

	return data
}
