package user

import "github.com/musab/gomock-demo/greeter"

type User struct {
	Greeter greeter.Greeter
}

func (u *User) Hello() (string, error) {
	return u.Greeter.Greet("GoMock")
}
