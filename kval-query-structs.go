package kvalparse

import "github.com/kval-access-language/kval-scanner"

type KQUERY struct {
	Function kvalscanner.Token
	Buckets  []string
	Key      string
	Value    string
	Newname  string
	Regex    bool
}
