package kvalparse

import "fmt"

/*
type KQUERY struct { 
   function Token
   buckets []string  
   key string
   value string
   newname string
   regex bool
}
*/

var (
   kq01 = KQUERY{INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "Value", "", false}
   kq02 = KQUERY{INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq03 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq04 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq05 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "PAT", "", "", true}
   kq06 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "Value", "", false}
   kq07 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "PAT", "", true}
   kq08 = KQUERY{LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq09 = KQUERY{LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq0a = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq0b = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq0c = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "_", "", false}
   kq0d = KQUERY{REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "New Key", false}
   kq0e = KQUERY{REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "Third Bucket", false}
   kq0f = KQUERY{INS, []string{"Prime Bucket"}, "", "", "", false}
   kq10 = KQUERY{INS, []string{"Prime Bucket"}, "key", "", "", false}
   kq11 = KQUERY{INS, []string{"Prime Bucket"}, "key", "value", "", false}
)

//Queries that should work according to the KVAL specification
var GoodQueryMap = map[string]string {
   "kq01_insert_value": "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: Value",
   "kq02_insert_value": "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
   "kq03_get_bucket_contents": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket",  
   "kq04_get_value": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq05_get_value_from_key_pattern": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}",
   "kq06_get_key_from_value": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value",
   "kq07_get_key_from_value_pattern": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}",
   "kq08_does_key_exist": "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq09_does_bucket_exist": "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
   "kq0a_delete_bucket": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
   "kq0b_delete_key": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq0c_delete_value": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: _ ",
   "kq0d_rename_key": "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key => New Key",
   "kq0e_rename_bucket": "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket => Third Bucket",   
   "kq0f_one_bucket": "INS Prime Bucket",   
   "kq10_one_bucket_key": "INS Prime Bucket >>>> key",   
   "kq11_one_bucket_key_value": "INS Prime Bucket >>>> key :: value",   
}

var GoodQueryExpected = map[string]KQUERY {
   "kq01_insert_value": kq01,
   "kq02_insert_value": kq02,
   "kq03_get_bucket_contents": kq03,
   "kq04_get_value": kq04,
   "kq05_get_value_from_key_pattern": kq05,
   "kq06_get_key_from_value": kq06,
   "kq07_get_key_from_value_pattern": kq07,
   "kq08_does_key_exist": kq08,
   "kq09_does_bucket_exist": kq09,
   "kq0a_delete_bucket": kq0a,
   "kq0b_delete_key": kq0b,
   "kq0c_delete_value": kq0c,
   "kq0d_rename_key": kq0d,
   "kq0e_rename_bucket": kq0e,
   "kq0f_one_bucket": kq0f,   
   "kq10_one_bucket_key": kq10,   
   "kq11_one_bucket_key_value": kq11,
}

var BadQueryMap = map[string]string {
   "badkq01_no_buckets": "INS",
   "badkq02_ins_regex": "INS Prime Bucket >>>> {PATT}",
   "badkq03_ins_regex": "INS Prime Bucket >>>> key :: {PATT}",
   "badkq04_ins_regex": "INS Prime Bucket >>>> {PATT} :: {PATT}",
   "badkq05_get_val": "GET Prime Bucket >>>> known :: unknown",              //if we know value, we don't need get
   "badkq06_lis_val": "LIS Prime Bucket >>>> known :: unknown",              //validate for yourself, for many reasons!   
   "badkq07_get_unknown": "GET Prime Bucket >>>> _",
   "badkq08_lis_unknown": "LIS Prime Bucket >>>> _",
}

var BadQueryExpected = map[string]error {
   "badkq01_no_buckets": fmt.Errorf("Zero buckets: No buckets specified in input query."),
   "badkq02_ins_regex": fmt.Errorf("Invalid Pattern use: Can't have regex on insert."),
   "badkq03_ins_regex": fmt.Errorf("Invalid Pattern use: Can't have regex on insert."),
   "badkq04_ins_regex": fmt.Errorf("Invalid Pattern use: Can't have regex on insert."),
   "badkq05_get_val": fmt.Errorf("Known Value: No need to GET a known value."), 
   "badkq06_lis_val": fmt.Errorf("Known Value: No need to LIS a known value."),     
   "badkq07_get_unknown": fmt.Errorf("Unknown unknown: Cannot seek unknown key and value."),
   "badkq08_lis_unknown": fmt.Errorf("Unknown unknown: Cannot seek unknown key and value."),
}
