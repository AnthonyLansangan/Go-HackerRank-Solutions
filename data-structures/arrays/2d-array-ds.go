/*
Given a 6 x 6 2D Array, :

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0
0 0 0 0 0 0

We define an hourglass in A to be a subset of values with indices falling in this pattern in 's graphical representation:

a b c
  d
e f g

There are 16 hourglasses in A, and an hourglass sum is the sum of an hourglass' values.

Task
Calculate the hourglass sum for every hourglass in A, then print the maximum hourglass sum.

Output Format

Print the largest (maximum) hourglass sum found in A.

Sample Input

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Sample Output

19

Explanation

A contains the following hourglasses:

1 1 1   1 1 0   1 0 0   0 0 0
  1       0       0       0
1 1 1   1 1 0   1 0 0   0 0 0

0 1 0   1 0 0   0 0 0   0 0 0
  1       1       0       0
0 0 2   0 2 4   2 4 4   4 4 0

1 1 1   1 1 0   1 0 0   0 0 0
  0       2       4       4
0 0 0   0 0 2   0 2 0   2 0 0

0 0 2   0 2 4   2 4 4   4 4 0
  0       0       2       0
0 0 1   0 1 2   1 2 4   2 4 0

The hourglass with the maximum sum (19) is:

2 4 4
  2
1 2 4

 */

package main
import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	io := bufio.NewReader(os.Stdin)
	var a [6][6]int
	for i := 0; i < 6; i++{
		for j := 0; j < 6; j++{
			var v int
			fmt.Fscan(io, &v)
			a[i][j] = v
		}
	}
	max := -1000
	for i := 0; i < 4; i++{
		for j := 0; j < 4; j++{
			top := a[i][j] + a[i][j+1] + a[i][j+2]
			mid := a[i+1][j+1]
			bottom := a[i+2][j] + a[i+2][j+1] + a[i+2][j+2]
			sum := top + mid + bottom
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Print(max)
}
