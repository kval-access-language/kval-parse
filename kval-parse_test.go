package kvalparse

import (
	"reflect"
	"testing"
)

//Exported Functions
var ExportDeconstruct = deconstruct
var ExportValidatePattern = validatepattern
var ExportExtendSlice = extendslice

//Test Exported function Parse for basic functionality...
func TestParse(t *testing.T) {
	for key, query := range GoodQueryMap {
		kq, err := Parse(query)
		if err != nil {
			t.Errorf("FAIL: Parse error \n %s \n %s", query, err)
		}
		if !reflect.DeepEqual(kq, GoodQueryExpected[key]) {
			t.Errorf("FAIL: Good query parsed incorrectly: \n %s \n Received: %v \n Expected: %v", query, kq, GoodQueryExpected[key])
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

//Test that base64 encoding and decoding work as required...
func TestBase64Input(t *testing.T) {
	kq, err := Parse(INS_base64_img_1)
	if err != nil {
		t.Errorf("FAIL: Parse error \n %s \n %s", INS_base64_img_1, err)
	}

	//we want to preserve the integrity of the values, so test those
	if kq.Value != GET_base64value_res_1 {
		//TODO: is outputting the base64 encoded value helpful?
		t.Errorf("FAIL: Parse error, base64 encoded value not preserved on parsing: \n %s \n %s", GET_base64value_res_1, kq.Value)
	}

	kq, err = Parse(INS_base64_img_2)
	if err != nil {
		t.Errorf("FAIL: Parse error \n %s \n %s", INS_base64_img_1, err)
	}

	//we want to preserve the integrity of the values, so test those
	if kq.Value != GET_base64value_res_2 {
		//TODO: is outputting the base64 encoded value helpful?
		t.Errorf("FAIL: Parse error, base64 encoded value not preserved on parsing: \n %s \n %s", GET_base64value_res_2, kq.Value)
	}

	kq, err = Parse(INS_base64_img_3)
	if err != nil {
		t.Errorf("FAIL: Parse error \n %s \n %s", INS_base64_img_1, err)
	}

	//we want to preserve the integrity of the values, so test those
	if kq.Value != GET_base64value_res_3 {
		//TODO: is outputting the base64 encoded value helpful?
		t.Errorf("FAIL: Parse error, base64 encoded value not preserved on parsing: \n %s \n %s", GET_base64value_res_3, kq.Value)
	}
}

//Test that big unicode strings can be Parsed correctly...
//TODO: Add some more exception characters to the test cases...
func TestBigString(t *testing.T) {
	kq, err := Parse("INS bucket one >> bucket two >>>> bigstring :: " + bigstring_one)
	if err != nil {
		t.Errorf("FAIL: Parse error, unicode string incorrectly not allowed: \n %v\n", err)
	} else if kq.Value != bigstring_one {
		t.Errorf("FAIL: Parse error, bigstring warped on input.\n", err)
	}

	kq, err = Parse("INS bucket one >> bucket two >>>> bigstring :: " + bigstring_two)
	if err != nil {
		t.Errorf("FAIL: Parse error, unicode string incorrectly not allowed: \n %v\n", err)
	} else if kq.Value != bigstring_two {
		t.Errorf("FAIL: Parse error, bigstring warped on input.\n", err)
	}
}

func TestExceptionQueries(t *testing.T) {
	var qstring = "GET ABC >> DEF >>>> GHI :: JKL"
	kq, err := Parse("INS bucket one >> bucket two >>>> bigstring :: " + qstring)
	if err != nil {
		t.Errorf("FAIL: Parse error, string not allowed: \n %v\n", err)
	} else if kq.Value != qstring {
		t.Errorf("FAIL: Parse error, example query string warped on input:\nReceived:%v\nExpected: %v\n", kq.Value, qstring)
	}

	qstring = ">> DEF >>>> GHI :: JKL"
	kq, err = Parse("INS bucket one >> bucket two >>>> bigstring :: " + qstring)
	if err != nil {
		t.Errorf("FAIL: Parse error, string not allowed: \n %v\n", err)
	} else if kq.Value != qstring {
		t.Errorf("FAIL: Parse error, example query string warped on input:\nReceived:%v\nExpected: %v\n", kq.Value, qstring)
	}
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
