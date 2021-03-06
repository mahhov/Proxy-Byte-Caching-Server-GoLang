package linkRange

import (
	"os"
	"bufio"
	"fmt"
	"testing"
)

// go test linkRange -v

// helper
func printRemove(r *LinkRange, s int, e int) *LinkRange {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
	fmt.Println("x ", s, e)
	r = r.RemoveRange(s, e)
	r.Print()
	fmt.Println()
	return r
}

// testing
func TestX(t *testing.T) {
	hundred := New(0, 100)
	hundred.Print()
	fmt.Println()
	
	hundred = printRemove(hundred, 6, 14)	// 0-5, 15-100
	hundred = printRemove(hundred, 31, 54)	// 0-5, 15-30, 55-100
	hundred = printRemove(hundred, 21, 30)	// 0-5, 15-20, 55-100
	hundred = printRemove(hundred, 55, 79)	// 0-5, 15-20, 80-100
	hundred = printRemove(hundred, 15, 20)	// 0-5, 80-100
	hundred = printRemove(hundred, 0, 5) // 80-100
	hundred = printRemove(hundred, 80, 100) //
}

