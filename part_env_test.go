package mc

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestGetenv(t *testing.T) {
	assertions := require.New(t)

	Must(os.Unsetenv("TEST_BLAH_BLAH"))
	assertions.Equal("no", EnvOr("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", ""))
	assertions.Equal("no", EnvOr("TEST_BLAH_BLAH", "no"))

	Must(os.Setenv("TEST_BLAH_BLAH", "yes"))
	assertions.Equal("yes", EnvOr("TEST_BLAH_BLAH", "no"))
}
