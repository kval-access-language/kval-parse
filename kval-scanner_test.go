package kvalparse

import (
   "strings"
   "testing"
)

type scannerTest struct {
   scanvalue    string
   expected     Token
}

var scannerTests = []scannerTest {
   {"", EOF},  
   {"*", LITERAL}, 
   {" ", WS}, 
   {"_", USCORE}, 
   {"\x1F", ILLEGAL},
}

func TestScan(t *testing.T) {
   //Test simple scan results to being with... more complex scans later.
   for _, expected := range scannerTests {
      s := NewScanner(strings.NewReader(expected.scanvalue))
      var tok Token
      var lit string 
      for tok != EOF {
         tok, lit = s.Scan()
         if (tok != EOF && expected.expected != EOF) && tok != expected.expected {
            //EOF returned each scan, so ignore if it's not what we're testing...      
            t.Errorf("FAIL: Got %d '%s' when %d '%s' was expected.", tok, lit, expected.expected, expected.scanvalue)
         }
      }
   }
}

func TestValidLiterals(t *testing.T) {
   //Test the full range of literals allowed in KVAL
   var lits = []string{"āēīōūĀĒĪŌŪ", ">>>", "/)(*&^%$#@!>:!@#", "abc123"}
   for _, s := range(lits) {
      s := NewScanner(strings.NewReader(s))
      var tok Token
      var lit string
      for tok != EOF {
         tok, lit = s.Scan()
         if tok == ILLEGAL {
            t.Errorf("FAIL: Illegal token %d '%s' when expecting valid LITERAL value.", tok, lit)
         } 
      }
   }
}