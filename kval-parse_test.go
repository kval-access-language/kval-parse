package kvalparse

import (
   "testing"
   "reflect"
   //"fmt"
   //"os"
)

//Exported Functions
var ExportDeconstruct = deconstruct
var ExportValidatePattern = validatepattern
var ExportExtendSlice = extendslice

func TestParse(t *testing.T) {
   //Test exported function Parse() using good queries...
   for key, query := range GoodQueryMap {
      kq, err := Parse(query)
      if err != nil {
         t.Errorf("FAIL: Parse error \n %s \n %s", query, err)
      }

      if !reflect.DeepEqual(kq, GoodQueryExpected[key]) {
         t.Errorf("FAIL: Good query parsed incorrectly: \n %s \n %v \n %v", query, kq, GoodQueryExpected[key])
      }
   }

   //Test exported function Parse() using bad queries... we're *want* errors
   for key, query := range BadQueryMap {
      kq, err := Parse(query)
      if err != nil {
         if !reflect.DeepEqual(err, BadQueryExpected[key]) {
            t.Errorf("FAIL: Bad query parsed incorrectly: \n %s \n %v \n %v \n %v", query, kq, err, BadQueryExpected[key])
         }
      } else if err == nil {
         t.Errorf("FAIL: Expecting error for query: \n %s \n %v", query, BadQueryExpected[key])   
      }
   }
}

func TestBase64Input(t *testing.T) {
   _, err := Parse(INS_base64_img_1)
   if err != nil {
      t.Errorf("FAIL: Parse error \n %s \n %s", INS_base64_img_1, err)
   }

   /*if !reflect.DeepEqual(kq, GoodQueryExpected[key]) {
      t.Errorf("FAIL: Good query parsed incorrectly: \n %s \n %v \n %v", query, kq, GoodQueryExpected[key])
   }*/
}

func TestExportDeconstruct(t *testing.T) {
   //Test non-exported function deconstruct() here.
}

func TestExportValidatePattern(t *testing.T) {
   //Test non-exported function validatepattern() here
}

func TestExportExtendSlice(t *testing.T) {
   //Test non-exported function extendslice() here
}
