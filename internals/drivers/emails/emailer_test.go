package emails_test

import (
	"go-scaffold/internals/drivers/emails"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	err := emails.NewEmailerEnv().SendCode("Signup verification code", "signup", "daqiancode@gmail.com", "123", "", 10)
	assert.Nil(t, err)

}
