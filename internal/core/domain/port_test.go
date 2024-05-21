package domain

import (
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewPort(t *testing.T) {
	t.Parallel()

	portID := faker.ID
	portName := faker.NAME
	portCode := faker.NAME
	portTimezone := faker.Timezone()

	t.Run("valid", func(t *testing.T) {
		port, err := NewPort(portID, portName, portCode, "", "",
			nil, nil, nil, "", portTimezone, nil)
		require.NoError(t, err)

		require.Equal(t, portID, port.Id())
		require.Equal(t, portCode, port.Code())
		require.Equal(t, portName, port.Name())
		require.Equal(t, portTimezone, port.Timezone())
	})

	t.Run("missing port ID", func(t *testing.T) {
		_, err := NewPort("", portName, "", "", "",
			nil, nil, nil, "", "", nil)
		require.Error(t, err)
	})
}
