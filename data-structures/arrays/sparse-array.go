/*
There is a collection of N strings ( There can be multiple occurences of a particular string ).
Each string's length is no more than 20 characters. There are also Q queries.
For each query, you are given a string,
and you need to find out how many times this string occurs in the given collection of N strings.
 */
package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var n int
	var q int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	s := make([]string,n)
	for i := 0; i < n; i++{
		fmt.Fscan(io, &s[i])
	}
	fmt.Fscan(io,&q)
	for i:= 0; i < q; i++{
		var qString string
		fmt.Fscan(io,&qString)
		var c int
		for j :=0; j < n; j++ {
			if qString == s[j]{
				c++
			}
		}
		fmt.Println(c)
	}
}