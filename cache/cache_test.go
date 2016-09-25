package cache

import (
	"fmt"
	"strings"
	"strconv"
	"testing"
)

// go test cache -v

var data = []byte {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

func stringToArray(s string) []string {
	s = strings.TrimSpace(s)
	stringArray := strings.Split(s, " ")
	return stringArray
}

func stringArrayToIntArray(stringArray []string) []int {
	intArray := make([]int, len(stringArray))

	for i, v := range stringArray {
		intArray[i], _ = strconv.Atoi(v)
	}

	return intArray[0:]
}

func getData(start int, end int) []byte {
	return data[start:end+1]
}

func TestX(t *testing.T) {
	c := New(5)
	
	requests := [][2]int{{4, 8},  {5, 7}, {4, 8}, {6, 10}, {4, 8}, {3, 5}, {0, 10}}
	
	for i := 0; i < len(requests); i++ {
		b := requests[i]
		fmt.Println("request> ", b)
		
		f, rem := c.FillFromCache(b[0], b[1])
		data := getData(b[0], b[1])
		c.FillCache(b[0], rem, data)
		
		fmt.Println("from cache: ", f)
		
		fmt.Print(  "remaining:   ")
		rem.Print()
		fmt.Println()
		
		c.Print()
	}
	
}