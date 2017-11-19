/*
A left rotation operation on an array of size n shifts each of the array's elements 1 unit to the left.
For example, if 2 left rotations are performed on array [1,2,3,4,5] , then the array would become [3,4,5,1,2].
 */
package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	var n int
	var d int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	fmt.Fscan(io, &d)
	a := make([]int, n)
	for i := 0; i < n; i++{
		fmt.Fscan(io, &a[i])
	}
	for i := 0; i < d; i++{
		a = append(a[1:n], a[0])
	}
	s := strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), " "), "[]")
	fmt.Print(s)
}