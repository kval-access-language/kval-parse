package kvalparse

type KQUERY struct { 
   function Token
   buckets []string  
   key string
   value string
   newname string
   regex bool
}
