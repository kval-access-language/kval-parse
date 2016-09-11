package kvalparse

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
   kq02 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq03 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq04 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "PAT", "", "", true}
   kq05 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "Value", "", false}
   kq06 = KQUERY{GET, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "_", "PAT", "", true}
   kq07 = KQUERY{LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq08 = KQUERY{LIS, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq09 = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "", false}
   kq0a = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "", false}
   kq0b = KQUERY{DEL, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "_", "", false}
   kq0c = KQUERY{REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "Key", "", "New Key", false}
   kq0d = KQUERY{REN, []string{"Prime Bucket", "Secondary Bucket", "Tertiary Bucket"}, "", "", "Third Bucket", false}
)

//Queries that should work according to the KVAL specification
var GoodQueryMap = map[string]string{
   "kq01_insert_value": "INS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: Value",
   "kq02_get_bucket_contents": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket",  
   "kq03_get_value": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq04_get_value_from_key_pattern": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> {PAT}",
   "kq05_get_key_from_value": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: Value",
   "kq06_get_key_from_value_pattern": "GET Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> _ :: {PAT}",
   "kq07_does_key_exist": "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq08_does_bucket_exist": "LIS Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
   "kq09_delete_bucket": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket",
   "kq0a_delete_key": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key",
   "kq0b_delete_value": "DEL Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key :: _ ",
   "kq0c_rename_key": "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket >>>> Key => New Key",
   "kq0d_rename_bucket": "REN Prime Bucket >> Secondary Bucket >> Tertiary Bucket => Third Bucket",   
}

var GoodQueryExpected = map[string]KQUERY{
   "kq01_insert_value": kq01,
   "kq02_get_bucket_contents": kq02,
   "kq03_get_value": kq03,
   "kq04_get_value_from_key_pattern": kq04,
   "kq05_get_key_from_value": kq05,
   "kq06_get_key_from_value_pattern": kq06,
   "kq07_does_key_exist": kq07,
   "kq08_does_bucket_exist": kq08,
   "kq09_delete_bucket": kq09,
   "kq0a_delete_key": kq0a,
   "kq0b_delete_value": kq0b,
   "kq0c_rename_key": kq0c,
   "kq0d_rename_bucket": kq0d,
}