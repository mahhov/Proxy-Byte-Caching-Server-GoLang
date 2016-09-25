# Proxy-Byte-Caching-Server-GoLang  

## Proxy Server  
#### Overview  
#### Example Inputs  
localhost:8080/?s=http://techslides.com/demos/sample-videos/small.mp4&range=0-  
localhost:8080/?s=http://techslides.com/demos/sample-videos/small.mp4&range=300-  
localhost:8080/?s=http://techslides.com/demos/sample-videos/small.mp4&range=0-1000  
localhost:8080/?s=http://techslides.com/demos/sample-videos/small.mp4&range=300-1000  


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
#### Alphabet Example  
	How FillFromCache Works:                                                                                
		If our data is the alphabet                                                                           
		And our cache contents are c.cache = [A, B, E, I, J]                                                  
		And we request the 1st to 8th letters of the alphabet                                                 
		f, rem := c.FillFromCache("alpha", 0, 8)                                                                       
		Then f = [A, B, _, _, E, _, _, _, I] // A, B, E, and I were found within the cache                    
		And rem = (2,3), (5,7) // C to D, and F to H are missing                                              
		Once we retrieve the missing letters (C, D, F, G, H)                                                  
		We store add them to the cache                                                                        
		By calling c.WriteToCache("alpha", 2, [C, D])                                                                  
		And c.WriteToCache("alpha", 5, [F, G, H])                                                                      
		Alternatively (recommended), c.FillCache("alpha", 0, rem, [A, B, C, D, E, F, G, H, I])                            
		Which will then look at rem and do the above 2 calls to c.WriteToCache                                
	
	How WriteToCache Works:                                                                          
		// 01234567890                                                                                 
		// ABCDEFGHIJK                                                                                 
		Imagine cache = [B, C, D, E, I, J, K]                                                          
		And nextCache = 0                                                                              
		And cacheMap = ["alpha" : [[1,4]:[0,3], [8,10]:[4,6]] ]                                                    
		Where [1,4] represents alpha[1:4] = B to E, [8,10] repreesents alpha[8:10] = I to K            
		And [0,3] represents cache[0:3] and [4,6] represents cache[4,6]                                
		Then calling WriteToCache("alpha", 5, [F, G, H])                                                        
		Will move nextCache to 3                                                                       
		And alter cache to [F, G, H, E, I, J, K]                                                       
		And add [5,7]:[0,2] (F to H) to cacheMap                                                       
		And replace [1,4]:[0,3] (B to E) with [4,4]:[3,3] (E) in cacheMap                              
		Resulting in cacheMap of [[4,4]:[3,3], [8,10]:[4,6], [5,7]:[0,2]]                              


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
