package kvalparse

type KQUERY struct { 
   Function Token
   Buckets []string  
   Key string
   Value string
   Newname string
   Regex bool
}
