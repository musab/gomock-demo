package greeter

//go:generate mockgen -package mocks -source=$GOFILE -destination=../mocks/mocks.go Greeter

type Greeter interface {
	Greet(string) (string, error)
}
