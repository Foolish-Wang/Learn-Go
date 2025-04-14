package packageone

// This variable is not exported and cannot be accessed outside this package
// Starts with a lowercase letter
var privateVar = "This is a private variable" 

var PublicVar = "This is a public variable"

// This func is not exported and cannot be accessed outside this package
// Starts with a lowercase letter
func notExported()  {
	
}

func Exported() {
	
}
