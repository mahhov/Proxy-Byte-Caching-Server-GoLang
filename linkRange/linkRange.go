package linkRange

import (
	"fmt"
)

/*
How To Use:
// create a new range (0, 100) inclusive
x := linkRange.New(0, 100)

// print contents of range
x.Print()

// subtract range (10, 49) from (0, 100) = (0, 9) & (50, 100)
x = x.RemoveRange(10, 90)

// Parameter Requirement
// in order to call x.RemoveRange(i, j), wherethe range (i,j) MUST be contained within range x

// Invarient
// it is safe to assume RemoveRange will never result in split adjacent ranges, such as (0, 50) & (51, 100); instead it will contain (0, 100)
*/


type LinkRange struct {
	Start int
	End int
	Next *LinkRange
	prev *LinkRange
}

func New(start int, end int) *LinkRange {
	return &LinkRange{start, end, nil, nil}
}

func (r *LinkRange) Print() {
	if r != nil {
		fmt.Print(r.Start, " -> ", r.End, "; ")
		r.Next.Print()
	}
}

func (r *LinkRange) add(s int, e int) {
	a := &LinkRange{s, e, r.Next, r}
	
	if r.Next != nil {
		r.Next.prev = a
	}
	r.Next = a
}

func (r *LinkRange) remove() *LinkRange {
	if r.prev != nil {
		r.prev.Next = r.Next
	}
	if r.Next != nil {
		r.Next.prev = r.prev
	}
	return r.Next
}

func (r *LinkRange) RemoveRange(s int, e int) *LinkRange {
	if overlap(s, e, r.Start, r.End) {
		if s > r.Start {
			r.add(r.Start, s - 1)
		}
		if e < r.End {
			r.add(e + 1, r.End)
		}
		return r.remove()
		
	} else if r.Next != nil {
		r.Next.RemoveRange(s, e)
	}
	return r
}

func overlap(v0 int, v1 int, val0 int, val1 int) bool {
	return v1 >= val0 && v0 <= val1
}

