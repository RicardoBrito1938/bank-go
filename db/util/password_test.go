package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(hashedPassword, wrongPassword)
	require.Error(t, err)

	hashedPassword2, err := Hash(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)

	require.NotEqual(t, hashedPassword2, hashedPassword)
}
