package main

func create3DArray() any {
	var array = [4][4][4]float32{1: {0: {2: 88.6}}} // modify only this line

	a := [4]int

	a[3] = 1

	return array
}
