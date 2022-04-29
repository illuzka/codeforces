package main

import (
	"fmt"
  "math"
)

func getData(n int) []int {
    i := 0 
    var data []int 
    
    for i < n*2 {
        var line int
        fmt.Scan(&line)
        data = append(data, line)
        i ++ 
    }
    return data 
}

func check(x, y int) string {
    i := 1
    for i <= x { 
        optimized := math.Ceil(float64(y) / (float64(i) + 1.00))
        newDays := i + int(optimized)
        if newDays <= x {
            return "YES"
        }
        i ++ 
    }
    return "NO"
}

func main() {
    var amt int 
    fmt.Scanln(&amt)
    test := getData(amt)
    var result []string  
    
    i := 0
    for i < len(test) {
        if test[i] >= test[i+1] {
            result = append(result, "YES")
        } else { 
            result = append(result, check(test[i], test[i+1]))
        }
        i += 2
    }
    for _, i := range result {
        fmt.Println(i)
    }
}
