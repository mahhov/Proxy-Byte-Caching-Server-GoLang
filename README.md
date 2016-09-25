# Proxy-Byte-Caching-Server-GoLang

## Proxy Server

## Cache  
#### Overview  
Byte array representing a cache. FIFO.
#### Public Fields:  
None
#### Public Methods:  
**New(cacheSize int)**  -  create a new cache of capacity cacheSize bytes  
**Print()**  - print contets of cache and cacheMap (a mapping of query -> cache indices)  
**FillFromCache(source string, start int, end int) ([]byte, *linkRange.LinkRange)**  -  1) search cache mapping for any portions of source within the range start to end. 2) return []byte of of length=(end-start+1) containing the portions found from cache and 0 elsewhere. 3) return a LinkRange representing the portions not found within the cache.  
**WriteToCache(source string, start int, body [] byte)**  -  1) store the contents of body in the cache. 2) add the mapping source->start->index of cache where body was stored. 3) remove/truncate mappings for the cache indices overwritten.  
**FillCache(source string, start int, rem *linkRange.LinkRange, body [] byte)**  -  1) rem contains the portions of body that are not in the cache. 2) add those un-cached portions of body to the cache via WriteToCache(...) and update mappings.  
#### Example  


## Link Range  
#### Overview  
Double linked list representing a range of values. supports range subtraction.  
#### Public Fields:  
x.Start, x.End, x.Next  
#### Public Methods:  
**New(start int, end int)**   
**Print()**  
**RemoveRange(start int, end int)**  
#### Example  
x := linkRange.New(0, 100) // x = 0 -> 100;  
x = x.RemoveRange(6, 14)  // x = 15 -> 100; 0 -> 5;  
x = x.RemoveRange(31, 54)  // x = 55 -> 100; 15 -> 30; 0 -> 5;  
x = x.RemoveRange(21, 30)  // x = 55 -> 100; 15 -> 20; 0 -> 5;  
x = x.RemoveRange(55, 79)  // x = 80 -> 100; 15 -> 20; 0 -> 5;  
x = x.RemoveRange(15, 20)  // 80 -> 100; 0 -> 5;  
x = x.RemoveRange(0, 5)  // x = 80 -> 100;  
x = x.RemoveRange(80, 100)  // x = ;  
