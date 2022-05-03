package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	seq := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		seq[i], _ = strconv.Atoi(scanner.Text())
	}

	if k == n {
		fmt.Println(1000000000)
		return
	}
	sort.Ints(seq)
	if k > 0 && seq[k-1] != seq[k] {
		fmt.Println(seq[k] - 1)
	} else if k == 0 && seq[0] != 1 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
