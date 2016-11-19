package kvalparse

import "github.com/pkg/errors"

var err_invalid_function = errors.New("Attempting to parse invalid function.")
var err_zero_buckets = errors.New("Zero buckets: No buckets specified in input query.")
var err_ins_regex = errors.New("Invalid Pattern use: Can't have regex on insert.")
var err_key_get_regex = errors.New("Known Value: No need to GET a known value.")
var err_key_lis_regex = errors.New("Known Value: No need to LIS a known value.")
var err_unk_unk = errors.New("Unknown unknown: Cannot seek unknown key and value.")
var err_no_name_rename = errors.New("Rename: Missing newname parameter.")
