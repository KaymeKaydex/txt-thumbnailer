package config_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/config"
)

func TestWrapAndFromContext(t *testing.T) {
	ctx := context.TODO()

	require.Nil(t, config.FromContext(ctx))

	ctx = config.WrapContext(ctx, &config.Config{})

	require.NotNil(t, config.FromContext(ctx))
}
