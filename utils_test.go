package gormpgtime

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTimeToLocation(t *testing.T) {
	loc, err := TimeToLocation("01:02")
	require.NoError(t, err)

	_, offset := time.Now().In(loc).Zone()
	require.Equal(t, 3720, offset)

	_, err = TimeToLocation("0102")
	require.Error(t, err)

	loc, err = TimeToLocation("-01:02")
	require.NoError(t, err)

	_, offset = time.Now().In(loc).Zone()
	require.Equal(t, -3720, offset)
}
