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
   {"*", ILLEGAL}, 
   {" ", WS}, 
   {"_", USCORE}, 
}

func TestScan(t *testing.T) {
   //Test simple scan results to being with... more complex scans later.
   for _, expected := range scannerTests {
      s := NewScanner(strings.NewReader(expected.scanvalue))
      var tok Token
      for tok != EOF {
         tok, _ = s.Scan()
         if (tok != EOF && expected.expected != EOF) && tok != expected.expected {
            //EOF returned each scan, so ignore if it's not what we're testing...      
            t.Errorf("FAIL: Got %d when %d was expected.", tok, expected.expected)
         }
      }
   }
}


