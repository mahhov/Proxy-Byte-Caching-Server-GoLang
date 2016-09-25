package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"cache2"
	"linkRange"
	"strings"
	"strconv"
)

// 

var cacheSize int = 64 * 1024 * 1024 // 64 MB. use cacheSize < 20 when verbose = true
var c *cache2.Cache = cache2.New(cacheSize)
var verbose bool = true

func getFromSource(f []byte, rem *linkRange.LinkRange, start int, client *http.Client, req *http.Request) {
	for rem != nil {
		// make http request
		req.Header.Set("Range", "bytes=" + strconv.Itoa(rem.Start) + "-" + strconv.Itoa(rem.End))
		res, _ := client.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		
		// fill f and do sanity check
		t := copy(f[rem.Start - start:], body)
		if t != rem.End - rem.Start + 1{
			fmt.Println("err?", t, rem.End - rem.Start + 1)
		}
		
		// itterate
		rem = rem.Next
	}
}

func getFromSourceNoRange(f []byte, start int, client *http.Client, req *http.Request) {
	res, _ := client.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	
	copy(f, body[start:])
}

// write out requested bytes
func outBytes (w http.ResponseWriter, r *http.Request, s string, rangee string) {
	// http header request
	client := &http.Client{}
	hreq, _ := http.NewRequest("HEAD", s, nil)
	res, _ := client.Do(hreq)
	
	// check header for Accept-Ranges
	_, headRange := res.Header["Accept-Ranges"]
	headLength, _ := res.Header["Content-Length"]
	if verbose {
		fmt.Println("\ncontent-length ", headLength)
	}
	
	// parse range
	rang := strings.Split(rangee, "-")
	start, _ := strconv.Atoi(rang[0])
	maxEnd, _ := strconv.Atoi(headLength[0])
	maxEnd -= 1
	end := maxEnd
	if len(rang) > 1 && rang[1] != "" {
		end, _ = strconv.Atoi(rang[1])
		if end > maxEnd {
			fmt.Println("Warning, requested byte range exceeds source byte range")
		}
	}
	if verbose {
		fmt.Println("request:     ", start, end)
	}
	
	// first check cache for request
	f, rem := c.FillFromCache(s, start, end)
	if verbose {
		if end - start < 20 {
			fmt.Println("from cache: ", f)
		}
		fmt.Print(  "remaining:   ")
		rem.Print()
		fmt.Println()
	}
	
	// check header for allowRange
	// and make appropiate (with or without range) http request
	if headRange {
		req, _ := http.NewRequest("GET", s, nil)
		getFromSource(f, rem, start, client, req)
	} else {
		fmt.Println("Warning, the requested url " + s + " does not allow range requests")
		req, _ := http.NewRequest("GET", s, nil)
		getFromSourceNoRange(f, start, client, req)
	}

	// store in cache
	c.FillCache(s, start, rem, f)
	if verbose && cacheSize < 20{
		c.Print()
	}
	
	// return body
	// w.Header().Set("Content-type", "video/mp4")
	if verbose {
		if end - start < 20 {
			fmt.Println("result:     ", f)
		} else {
			fmt.Println("result:     ", "len -", len(f))
		}
	}
	w.Write(f)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// parse query
	q := r.URL.Query()
	if _, ok := q["s"]; !ok {
		return
	}
	if _, ok := q["range"]; !ok {
		return
	}

	s := q["s"][0]
	rangee := q["range"][0]
	
	// write out requested bytes
	outBytes(w, r, s, rangee)
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
