/*
 * Exercise 2.4: Write a version of PopCount that counts bits by shifting its argument through 64
 * bit positions, testing the rightmost bit each time.
 * Author: gaohuaid
 * Date: 2019/06/02
 */
package main

import(
	"fmt"
)

func main(){
	fmt.Println(PopCount(8))
}

func PopCount(x uint64) int{
	var mask uint64 = 1
	var count int = 0
	
		for i := 0; i < 64; i++{
			if (x & mask) == 1{
				count ++
			}
			x = x >> 1
		}
	
	return count
}


