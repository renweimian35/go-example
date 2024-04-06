package services

import "fmt"

func Test(code string) (string, error) {
	return fmt.Sprintf("%s:%s", code, "code is cheap,show me the talking"), nil
}
