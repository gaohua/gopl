/*
 * Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression. Compare the
 * performance of the two version.
 * Author: gaohuaid
 * Date: 2019/06/02
 */
package popcount

var pc [256] byte

func init(){
	for i:=range pc{
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int{
	var count byte = 0
	var i uint8 = 0
	for ; i < 8; i++ {
		count += pc[byte(x >> (i*8))]
	}
	return int(count)
}