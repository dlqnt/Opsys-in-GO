package errors

import (
	"fmt"
)

/*
Task 5: Errors needed for multiwriter

You may find this blog post useful:
http://blog.golang.org/error-handling-and-go

Similar to a the Stringer interface, the error interface also defines a
method that returns a string.

type error interface {
    Error() string
}

Thus also the error type can describe itself as a string. The fmt package (and
many others) use this Error() method to print errors.

Implement the Error() method for the Errors type defined above.

The following conditions should be covered:

1. When there are no errors in the slice, it should return:

"(0 errors)"

2. When there is one error in the slice, it should return:

The error string return by the corresponding Error() method.

3. When there are two errors in the slice, it should return:

The first error + " (and 1 other error)"

4. When there are X>1 errors in the slice, it should return:

The first error + " (and X other errors)"
*/
func (m Errors) Error() string {
	result := ""
	counter := 0
	var errorListe []error
	//Filtrer ut nil
	for i := 0; i < len(m); i++ {
		if m[i] != nil {
			errorListe = append(errorListe, m[i])
			counter++
		}
	}

	for j := 0; j < len(errorListe); j++ {
		result = errorListe[j].Error()
	}

	if len(errorListe) == 0 {
		result = "(0 errors)"
	}
	if len(errorListe) == 1 {
		result = errorListe[0].Error()
	}
	if len(errorListe) == 2 {
		if errorListe[0] != nil && errorListe[1] != nil {
			result += " (and 1 other error)"
		}
	} else if len(errorListe) >= 3 {
		counter--
		result += " (and " + fmt.Sprint(counter) + " other errors)"
	}
	return result
}
