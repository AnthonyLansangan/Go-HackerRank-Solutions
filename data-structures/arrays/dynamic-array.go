package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var seqNum int
	var qNum int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &seqNum)
	fmt.Fscan(io, &qNum)
	seqs := make([][]int, seqNum)
	var lastAnswer int
	for i := 0; i < qNum; i++ {
		var qType int
		var x int
		var y int
		fmt.Fscan(io, &qType)
		fmt.Fscan(io, &x)
		fmt.Fscan(io, &y)
		index := (x ^ lastAnswer) % seqNum
		switch qType {
		case 1:
			seqs[index] = append(seqs[index], y)
		case 2:
			size := len(seqs[index])
			lastAnswer = seqs[index][y%size]
			fmt.Println(lastAnswer)
		}
	}
}
