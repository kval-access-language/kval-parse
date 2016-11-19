package kvalparse

import "github.com/kval-access-language/kval-scanner"

/*
type KQUERY struct { 
   function kvalscanner.Token
   buckets []string  
   key string
   value string
   newname string
   regex bool
}
*/

var (
   kq01 = KQUERY{kvalscanner.INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "Value", "", false}
   kq02 = KQUERY{kvalscanner.INS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq03 = KQUERY{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq04 = KQUERY{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq05 = KQUERY{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "PAT", "", "", true}
   kq06 = KQUERY{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "Value", "", false}
   kq07 = KQUERY{kvalscanner.GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "PAT", "", true}
   kq08 = KQUERY{kvalscanner.LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq09 = KQUERY{kvalscanner.LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq0a = KQUERY{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq0b = KQUERY{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq0c = KQUERY{kvalscanner.DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "_", "", false}
   kq0d = KQUERY{kvalscanner.REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "New Key", false}
   kq0e = KQUERY{kvalscanner.REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "Third Bucket", false}
   //non-reference-queries
   kq0f = KQUERY{kvalscanner.INS, []string{"Prime Bucket"}, "", "", "", false}
   kq10 = KQUERY{kvalscanner.INS, []string{"Prime Bucket"}, "key", "", "", false}
   kq11 = KQUERY{kvalscanner.INS, []string{"Prime Bucket"}, "key", "value", "", false}
   kq12 = KQUERY{kvalscanner.INS, []string{"Prime Bucket"}, "key", "hyphen-value", "", false}
   kq13 = KQUERY{kvalscanner.GET, []string{"Prime Bucket"}, "_", "PATT WITH SPACE", "", false}
   kq14 = KQUERY{kvalscanner.INS, []string{"Prime Bucket"}, "key", "value with space", "", false}
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
   "kq12_one_bucket_hyphen_value": "INS Prime Bucket >>>> key :: hyphen-value",   
   "kq13_regex_spaces": "GET Prime Bucket >>>> _ :: {PATT WITH SPACE}",   
   "kq14_value_spaces": "INS Prime Bucket >>>> key :: value with space",   
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
   "kq12_one_bucket_hyphen_value": kq12,
   "kq13_regex_spaces": kq13,
   "kq14_value_spaces": kq14,
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
   "badkq09_ren_bucket": "REN Prime Bucket => ",
   "badkq0a_ren_key": "REN Prime Buckey >>>> Key => ",
}

var BadQueryExpected = map[string]error {
   "badkq01_no_buckets": err_zero_buckets,
   "badkq02_ins_regex": err_ins_regex,
   "badkq03_ins_regex": err_ins_regex,
   "badkq04_ins_regex": err_ins_regex,
   "badkq05_get_val": err_key_get_regex, 
   "badkq06_lis_val": err_key_lis_regex,     
   "badkq07_get_unknown": err_unk_unk,
   "badkq08_lis_unknown": err_unk_unk,
   "badkq09_ren_bucket": err_no_name_rename,
   "badkq0a_ren_key": err_no_name_rename,
}
