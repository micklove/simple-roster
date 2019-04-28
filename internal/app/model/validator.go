package model

type Validator interface {
	Validate(*interface{}) (valid bool, err error)
}

//See file:///Users/lovemi/dev/Books/lets-go/html/11.03-user-signup-and-password-encryption.html

//// Use the regexp.MustCompile() function to parse a pattern and compile a
//// regular expression for sanity checking the format of an email address.
//// This returns a *regexp.Regexp object, or panics in the event of an error.
//// Doing this once at runtime, and storing the compiled regular expression
//// object in a variable, is more performant than re-compiling the pattern with
//// every request.
//var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
//
//type Form struct {
//url.Values
//Errors errors
//}
