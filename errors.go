package main

import (
	"fmt"
)

var ErrNotFound = fmt.Errorf("Not Found")
var ErrBadUrl = fmt.Errorf("differ can only fetch http or https urls")
