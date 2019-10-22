package common

import "fmt"

func ReadValue() string {
	var value string
	fmt.Scanf("%s", &value)
	return value
}
