package kvalparse

//https://blog.gopheracademy.com/advent-2014/parsers-lexers/
//https://github.com/fatih/hcl/blob/8f83adfc08e6d7162ef328a06cf00ee5fb865f30/scanner/scanner.go

import (
   "fmt"
   "regexp"
   "strings"
)

//maintain state
//queries run temporally buckets >> key >> value
var (
   keyword bool
   bucket bool
   key bool
   value bool
   newname bool
)

func setupstate() {
   keyword = true
   bucket = false
   key = false
   value = false
   newname = false
}

func Parse(query string) (KQUERY, error) {

   setupstate()

   var kq KQUERY
   var LITCACHE string

   var PATTERN = false
   var PATCACHE string

   s := NewScanner(strings.NewReader(query))
   
   var tok Token
   var lit string

   for tok != EOF {
      tok, lit = s.Scan()
      if tok == LITERAL {
         if PATTERN == true {
            PATCACHE = PATCACHE + " " + lit
         } else {
            LITCACHE = LITCACHE + " " + lit
         }
      } else if tok == OPATT {
         PATTERN = true
      } else if tok == CPATT {
         //validate patern
         //Add it to Key or Value as appropriate
         pattern, err := validatepattern(strings.TrimSpace(PATCACHE))
         if err != nil {
            return kq, err
         }
         kq, err = deconstruct(kq, LITERAL, pattern)
         if err != nil {
            return kq, err
         }
         kq.regex = true
         PATTERN = false
      } else if tok != WS {
         var err error
         if LITCACHE != "" {
            //Literal can be A bucket name, key name, or value name
            kq, err = deconstruct(kq, LITERAL, LITCACHE)
            if err != nil {
               return kq, err
            }
            LITCACHE = ""
         }
         if tok != EOF {
            //Keyword dictates the type of operation
            //Operator dictates where in the struct we need to place the value 
            kq, err = deconstruct(kq, tok, lit)
            if err != nil {
               return kq, err
            }
         }
      } 
   }   

   return kq, nil
}

func deconstruct(kq KQUERY, tok Token, lit string) (KQUERY, error) {

   lit = strings.TrimSpace(lit)

   //seek function keyword first
   if keyword == true {
      if KeywordMap[lit] == 0 {
         return kq, fmt.Errorf("Invalid function: %s", lit)
      } else {
         kq.function = tok
         keyword = false
         bucket = true
         return kq, nil
      }
   }

   if bucket == true {
      if tok == BUCKEY {
         bucket = false
         key = true
      } else if tok == BUCBUC {
         //bucket to bucket relationship, do nothing
      } else if tok == ASSIGN {
         //looking to rename bucket 
         bucket = false
         newname = true
      } else {
         kq.buckets = extendslice(kq.buckets, lit)
      }
      return kq, nil
   }

   if key == true {
      kq.key = lit
      key = false       //key added, can only be one
      return kq, nil
   }

   if value == true {
      kq.value = lit
      value = false        //value added, can only be one
      return kq, nil
   }

   if tok == KEYVAL {
      key = false
      value = true
      return kq, nil
   }

   if tok == ASSIGN {
      bucket = false
      key = false
      value = false      
      newname = true
      return kq, nil
   }

   if newname == true {
      kq.newname = lit
      return kq, nil
   }

   return kq, fmt.Errorf("Invalid query: Parsed without finding any new tokens.")
}

//Attempt to compile the pattern to see if it is valid and return itself
func validatepattern(pattern string) (string, error) {
   _, err := regexp.Compile(pattern)      //n.b. CompilePOSIX() too
   if err != nil {
      err = fmt.Errorf("Invalid regex: Cannot compile regular expression.")
   }
   return pattern, err
}

func extendslice(slice []string, element string) []string {
    n := len(slice)
    if n == cap(slice) {
        // Slice is full; must grow.
        // We double its size and add 1, so if the size is zero we still grow.
        newSlice := make([]string, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}