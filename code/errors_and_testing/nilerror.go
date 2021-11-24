package main

import "log"

// START1 OMIT
type myError struct{}

func (e *myError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return "ka-booom"
}

func check() *myError {
	return nil
}

func main() {
	var err error
	if err = check(); err != nil {
		log.Fatal("check failed: ", err)
	}
}

// END1 OMIT
