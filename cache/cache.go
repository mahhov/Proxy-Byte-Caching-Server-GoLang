package cache

import (
	"linkRange"
	"fmt"
)

/*
HOW TO USE:
	// create new cache of size n (int) Bytes
	c := cache.New(n) 

	// print cache and cacheMap
	c.Print()
	
	// try to find data[start:end] in cache
	// return the portions found in f ([]byte)
	// and the indices of the portions not found in rem (*LinkRange)
	f, rem := c.FillFromCache(start, end)

	// add d ([]byte) to cache, where d = data[start : start+len(data)]
	// all of d will be added to the cache
	// there is really no reason this method should be used unless manual control of what is added to cache and what isn't is required
	c.WriteToCache(start, d)
	
	// add d ([]byte) to cache, where d = data[start:end], and rem is the 2nd returned value from c.FillFromCache(start, end)
	// this will look at rem and effectively add only the parts of d that were not already in the cache rather than all of d
	c.FillCache(rem, d)

ALPHABET EXAMPLE:
	How FillFromCache Works:
		If our data is the alphabet
		And our cache contents are c.cache = [A, B, E, I, J]
		And we request the 1st to 8th letters of the alphabet
		f, rem := c.FillFromCache(0, 8)
		Then f = [A, B, _, _, E, _, _, _, I] // A, B, E, and I were found within the cache
		And rem = (2,3), (5,7) // C to D, and F to H are missing
		Once we retrieve the missing letters (C, D, F, G, H)
		We store add them to the cache
		By calling c.WriteToCache(2, [C, D])
		And c.WriteToCache(5, [F, G, H])
		Alternatively (recommended), c.FillCache(rem, [A, B, C, D, E, F, G, H, I])
		Which will then look at rem and do the above 2 calls to c.WriteToCache
	
	How WriteToCache Works:
		// 01234567890
		// ABCDEFGHIJK
		Imagine cache = [B, C, D, E, I, J, K]
		And nextCache = 0
		And cacheMap = [[1,4]:[0,3], [8,10]:[4,6]]
		Where [1,4] represents alpha[1:4] = B to E, [8,10] repreesents alpha[8:10] = I to K
		And [0,3] represents cache[0:3] and [4,6] represents cache[4,6]
		Then calling WriteToCache(5, [F, G, H])
		Will move nextCache to 3
		And alter cache to [F, G, H, E, I, J, K]
		And add [5,7]:[0,2] (F to H) to cacheMap
		And replace [1,4]:[0,3] (B to E) with [4,4]:[3,3] (E) in cacheMap
		Resulting in cacheMap of [[4,4]:[3,3], [8,10]:[4,6], [5,7]:[0,2]]
*/

type Cache struct {
	cacheSize int 
	cache []byte
	nextCache int
	
	// data range -> cached location
	cacheMap map[[2]int][2]int
}

func New(cacheSize int) *Cache {
	return &Cache{cacheSize, make([]byte, cacheSize), 0, make(map[[2]int][2]int)}
}

func (c *Cache) Print() {
	fmt.Print(  "cache map:   ")
	fmt.Println(c.cacheMap)
	
	fmt.Print(  "cache:       ")
	fmt.Println(c.cache)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *Cache) FillFromCache(start int, end int) ([]byte, *linkRange.LinkRange) {
	f := make([]byte, end - start + 1)
	rem := linkRange.New(start, end)
	
	for k, v := range c.cacheMap {
		if overlap(k[0], k[1], start, end) {
			s := max(k[0], start)
			e := min(k[1], end)
			d := e - s
			for i := 0; i <= d; i++ {
				f[s + i - start] = c.cache[v[0] + s - k[0] + i]
			}
			rem = rem.RemoveRange(s, e) 
		}
	}
	
	return f, rem
}

func overlap(v0 int, v1 int, val0 int, val1 int) bool {
	return v1 >= val0 && v0 <= val1
}

func (c *Cache) addCacheMap2(key [2]int, val [2]int) {
	for k, v := range c.cacheMap {
		if overlap(v[0], v[1], val[0], val[1]) {
			delete(c.cacheMap, k)
			
			if val[1] < v[1]  {
				t := val[1] - v[0] + 1
				v[0] += t
				k[0] += t
				c.cacheMap[k] = v
			}
		} 
	}
	
	c.cacheMap[key] = val
}

func (c *Cache) addCacheMap(key0 int, val0 int, dist int) {
	if dist != 0 {
		dist -= 1
		c.addCacheMap2 ([2]int {key0, key0 + dist}, [2]int {val0, val0 + dist})
	}
}

func (c *Cache) WriteToCache(start int, body [] byte) {
	if c.cacheSize > len(body) {
		t := copy(c.cache[c.nextCache:], body)
		c.addCacheMap(start, c.nextCache, t)
		c.nextCache += t
		
		if c.nextCache == c.cacheSize {
			c.nextCache = copy(c.cache[0:], body[t:])
			c.addCacheMap(start + t, 0, c.nextCache)
		}
		
	} else {
		c.nextCache = 0
		copy(c.cache[0:], body)
		c.addCacheMap(start, 0, c.cacheSize)
	}
}

func (c *Cache) FillCache(start int, rem *linkRange.LinkRange, body [] byte) {
	for rem != nil {
		c.WriteToCache(rem.Start, body[rem.Start - start : rem.End - start + 1])
		rem = rem.Next
	}
}

