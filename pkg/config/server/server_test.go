package server

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ServerConfig(t *testing.T) {
	fakeCMD := &cobra.Command{}
	RegisterConfig(fakeCMD)

	cfg := GetConfig()
	assert.Equal(t, configKeyBindAddrDefault, cfg.Bind)
	assert.Equal(t, uint16(configKeyBindPortDefault), cfg.Port)
}
