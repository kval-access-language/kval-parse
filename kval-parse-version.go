package kvalparse

var kval_boltdb_version = "0.0.0-KVAL-Working-Draft"

//Return an indication of which version of the KVAL language we are
//working from and the version of the library that you are implementing from
func Version() string {
	return kval_boltdb_version
}

