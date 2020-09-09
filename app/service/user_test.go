package service

import (
	"fmt"
	"github.com/libmonsoon-dev/fasthttp-template/app/infrastructure/logger"
	"testing"
)

func newUserService() *UserService {
	return NewUserService(
		logger.NewStderrLogger(),
		nil,
	)
}

func TestAuthService_ComparePassword(t *testing.T) {
	type Test struct {
		first    []byte
		second   []byte
		expected bool
	}

	tests := []Test{
		{[]byte("test"), []byte("test"), true},
		{[]byte("test"), []byte("test_"), false},
		{[]byte("test"), []byte("tess"), false},
		{[]byte(""), []byte(""), true},
		{[]byte(""), nil, true},
		{nil, []byte(""), true},
		{nil, nil, true},
	}

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprintf("#%v", i+1), func(t *testing.T) {
			t.Parallel()

			service := newUserService()
			hash := service.generatePasswordHash(test.first)
			actual := service.comparePassword(test.second, hash)

			if actual != test.expected {
				t.Errorf("Expected: %v, actual: %v", test.expected, actual)
			}
		})
	}
}
