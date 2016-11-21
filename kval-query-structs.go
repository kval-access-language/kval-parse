package kvalparse

import "github.com/kval-access-language/kval-scanner"

type KQuery struct {
	Function kvalscanner.Token
	Buckets  []string
	Key      string
	Value    string
	Newname  string
	Regex    bool
}
