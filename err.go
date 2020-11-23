package auth

import "fmt"

type Err struct{}

func (Err) Msg(msg string) error {
	return fmt.Errorf("[auth_err]->(%s)", msg)
}
