package kvalparse

//https://blog.gopheracademy.com/advent-2014/parsers-lexers/
//https://github.com/fatih/hcl/blob/8f83adfc08e6d7162ef328a06cf00ee5fb865f30/scanner/scanner.go

import (
   "regexp"
   "strings"
   "strconv"
   "unicode/utf8"
   "github.com/pkg/errors"
   "github.com/kval-access-language/kval-scanner"
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

   s := kvalscanner.NewScanner(strings.NewReader(query))
   
   var tok kvalscanner.Token
   var lit string

   for tok != kvalscanner.EOF {
      tok, lit = s.Scan()
      if tok != kvalscanner.ILLEGAL {
         if tok == kvalscanner.LITERAL {
            if PATTERN == true {
               PATCACHE = PATCACHE + lit
            } else {
               LITCACHE = LITCACHE + lit
            }
         } else if tok == kvalscanner.WS {
            LITCACHE = LITCACHE + lit    //repatriate whitespace from input        
         } else if tok == kvalscanner.USCORE {
            LITCACHE = LITCACHE + lit
         }else if tok == kvalscanner.OPATT {
            PATTERN = true
         } else if tok == kvalscanner.CPATT {
            //validate patern
            //Add it to Key or Value as appropriate
            pattern, err := validatepattern(strings.TrimSpace(PATCACHE))
            if err != nil {
               return kq, err
            }
            kq, err = deconstruct(kq, kvalscanner.LITERAL, pattern)
            if err != nil {
               return kq, err
            }
            kq.Regex = true
            PATTERN = false
         } else if tok != kvalscanner.WS {
            var err error
            if LITCACHE != "" {
               //kvalscanner.LITERAL: can be A bucket name, key name, or value name
               kq, err = deconstruct(kq, kvalscanner.LITERAL, LITCACHE)
               if err != nil {
                  return kq, err
               }
               LITCACHE = ""
            }
            if tok != kvalscanner.EOF {
               //Keyword dictates the type of operation
               //Operator dictates where in the struct we need to place the value 
               kq, err = deconstruct(kq, tok, lit)
               if err != nil {
                  return kq, err
               }
            }
         }
      } else {
         r, s := utf8.DecodeRune([]byte(lit))
         if s != 0 {
            unicode := strconv.QuoteRuneToASCII(r)
            return kq, errors.Wrapf(err_illegal_token, "'%s', %s.\n", lit, unicode)
         } else {
            return kq, errors.Wrapf(err_illegal_token, "'%s'.\n", lit)
         }
      }
   }

   //return kq, nil
   return validatequerystruct(kq)   
}

func deconstruct(kq KQUERY, tok kvalscanner.Token, lit string) (KQUERY, error) {

   lit = strings.TrimSpace(lit)

   if !value {
      //seek function keyword first
      if keyword == true {
         if kvalscanner.KeywordMap[lit] == 0 {
            return kq, err_invalid_function
         } else {
            kq.Function = tok
            keyword = false
            bucket = true
            return kq, nil
         }
      }

      if bucket == true {
         if tok == kvalscanner.BUCKEY {
            bucket = false
            key = true
         } else if tok == kvalscanner.BUCBUC {
            //bucket to bucket relationship, do nothing
         } else if tok == kvalscanner.ASSIGN {
            //looking to rename bucket 
            bucket = false
            newname = true
         } else {
            kq.Buckets = extendslice(kq.Buckets, lit)
         }
         return kq, nil
      }

      if key == true {
         kq.Key = lit
         key = false       //key added, can only be one
         return kq, nil
      }

      if tok == kvalscanner.KEYVAL {
         key = false
         value = true
         return kq, nil
      }
   }

   if tok == kvalscanner.ASSIGN {
      bucket = false
      key = false
      value = false      
      newname = true
      return kq, nil
   }

   if newname == true {
      kq.Newname = lit
      return kq, nil
   }

   if tok == kvalscanner.LITERAL && lit == "" {
      //no error, just nothing else to do...
      return kq, nil
   }

   //once value flag is set, treat it *all* as a value until EOM...
   if value == true {
      if kq.Value == "" {
         kq.Value = lit
      } else {
         kq.Value = kq.Value + " " + lit
      }
      return kq, nil
   }

   return kq, errors.Wrapf(err_parsed_no_new_tokens, "'%v', '%v'", tok, lit)
   //return kq, nil
}

//Attempt to compile the pattern to see if it is valid and return itself
func validatepattern(pattern string) (string, error) {
   _, err := regexp.Compile(pattern)      //n.b. CompilePOSIX() too
   if err != nil {
      err = err_compile_regex
   }
   return pattern, err
}

func validatequerystruct(kq KQUERY) (KQUERY, error) {
   //check for buckets
   if len(kq.Buckets) < 1 {
      return kq, err_zero_buckets
   }   
   if kq.Function == kvalscanner.INS && kq.Regex == true {
      return kq, err_ins_regex
   }
   if (kq.Function == kvalscanner.LIS || kq.Function == kvalscanner.GET) && kq.Regex != true {
      if kq.Key != "" && kq.Key != "_" {
         //trying to use REGEX for a key that is known...
         if kq.Value != "" {
            if kq.Function == kvalscanner.GET {
               return kq, err_key_get_regex  
            } else {
               return kq, err_key_lis_regex
            }
         }
      }    
   }
   //unless we want this to be a synonym for getting all values from a bucket...
   if (kq.Function == kvalscanner.GET || kq.Function == kvalscanner.LIS) && (kq.Key == "_" && kq.Value == "") {
      return kq, err_unk_unk
   }
   //rename capability
   if kq.Function == kvalscanner.REN && kq.Newname == "" {
      return kq, err_no_name_rename   
   }
   //searching for a known
   if kq.Function == kvalscanner.GET && (kq.Key != "_" && kq.Key != "") && kq.Value != "" {
      return kq, err_key_get_regex
   }
   return kq, nil
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
