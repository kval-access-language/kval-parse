package kvalparse

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota    // used in declarations of incrementing numbers
	EOF                     // EOF
	WS                      // Whitespace

	// Literals
	LITERAL                   // main

   // Other Operators (in order)
   BUCKEY                  // >>>> Bucket to Key Hierarchy
   BUCBUC                  // >> Bucket Hierarchy
   KEYVAL                  // :: Key//Value pair
   ASSIGN                  // => Assignment for rename operations

   // Single character operators
   USCORE                  // _ Return key OR value, for unknown key OR value 
   OPATT                   // { Open a regex pattern
   CPATT                   // | Close a regex pattern

	// Keywords
   INS                     // Insert
   GET                     // Get values
   LIS                     // Check existence
   DEL                     // Delete values
   REN                     // Rename values

   // Regex
   REGEX                   // {PATT} ANy regex pattern inside these two values
)

//map to help validate
//may be unecessary redundancy...
var KeywordMap = map[string]int{
   "INS":   0x1,         // Insert
   "ins":   0x1,

   "GET":   0x2,         // Get values
   "get":   0x2, 

   "LIS":   0x3,         // Check existence
   "lis":   0x3, 

   "DEL":   0x4,         // Delete values
   "del":   0x4,

   "REN":   0x5,         // Rename values   
   "ren":   0x5,
}
