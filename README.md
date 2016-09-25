# Proxy-Byte-Caching-Server-GoLang

## Proxy Server

## Cache  
#### Overview  
Byte array representing a cache.
#### Public Fields:  
None
#### Public Methods:  
New(cacheSize int),  
Print(),  
WriteToCache(source string, start int, body [] byte),  
FillCache(source string, start int, rem *linkRange.LinkRange, body [] byte),  
#### Example  


## Link Range  
#### Overview  
Double linked list representing a range of values. supports range subtraction.  
#### Public Fields:  
x.Start, x.End, x.Next  
#### Public Methods:  
New(start int, end int),   
Print(),  
RemoveRange(start int, end int)  
#### Example  
x := linkRange.New(0, 100) // x = 0 -> 100;  
x = x.RemoveRange(6, 14)  // x = 15 -> 100; 0 -> 5;  
x = x.RemoveRange(31, 54)  // x = 55 -> 100; 15 -> 30; 0 -> 5;  
x = x.RemoveRange(21, 30)  // x = 55 -> 100; 15 -> 20; 0 -> 5;  
x = x.RemoveRange(55, 79)  // x = 80 -> 100; 15 -> 20; 0 -> 5;  
x = x.RemoveRange(15, 20)  // 80 -> 100; 0 -> 5;  
x = x.RemoveRange(0, 5)  // x = 80 -> 100;  
x = x.RemoveRange(80, 100)  // x = ;  
