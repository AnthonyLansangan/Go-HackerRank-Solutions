/*
Print all N integers in A in reverse order as a single line of space-separated integers.

Sample Input:
4
1 4 3 2

Sample Output:
2 3 4 1
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
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)
	a := make([]int,n)
	n = n-1
	for n >=0 {
		var temp int
		fmt.Fscan(io, &temp)
		a[n]=temp
		n--
	}
	s := strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), " "), "[]")
	fmt.Print(s)
}