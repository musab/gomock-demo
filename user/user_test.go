package user_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/musab/gomock-demo/mocks"
	"github.com/musab/gomock-demo/user"
)

func TestUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockGreeter := mocks.NewMockGreeter(mockCtrl)
	testUser := &user.User{Greeter: mockGreeter}

	mockResponse := "Hello GoMock"

	mockGreeter.EXPECT().Greet("GoMock").Return(mockResponse, nil).Times(1)

	s, _ := testUser.Hello()

	if s != mockResponse {
		t.Errorf("Expected %s, got %s", mockResponse, s)
	}
}

func TestUserError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockGreeter := mocks.NewMockGreeter(mockCtrl)
	testUser := &user.User{Greeter: mockGreeter}

	dummyError := errors.New("dummy error")

	mockGreeter.EXPECT().Greet("GoMock").Return("", dummyError).Times(1)

	_, err := testUser.Hello()

	if err != dummyError {
		t.Fail()
	}
}
