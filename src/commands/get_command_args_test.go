package commands_test

import (
	"github.com/eaardal/workwork/src/commands"
	"github.com/eaardal/workwork/src/ww"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAndValidateGetCommandArgs_Key_ReturnsUrlKeys(t *testing.T) {
	ctx := createContextWithArgs(t, "get", "docs")
	args, err := commands.ParseAndValidateGetCommandArgs(ctx)
	assert.Nil(t, err)
	assert.Len(t, args.UrlKeys, 1)
	assert.Len(t, args.UrlKeys[ww.Global], 1)
	assert.Equal(t, "docs", args.UrlKeys[ww.Global][0])
}

func TestParseAndValidateGetCommandArgs_ManyKeys_ReturnsUrlKeys(t *testing.T) {
	ctx := createContextWithArgs(t, "get", "docs", "tasks", "ci")
	args, err := commands.ParseAndValidateGetCommandArgs(ctx)
	assert.Nil(t, err)
	assert.Len(t, args.UrlKeys, 1)
	assert.Len(t, args.UrlKeys[ww.Global], 3)
	assert.Equal(t, "docs", args.UrlKeys[ww.Global][0])
	assert.Equal(t, "tasks", args.UrlKeys[ww.Global][1])
	assert.Equal(t, "ci", args.UrlKeys[ww.Global][2])
}

func TestParseAndValidateGetCommandArgs_ManyKeysVariousEnvironments_ReturnsUrlKeys(t *testing.T) {
	ctx := createContextWithArgs(t, "get", "docs", "local.tasks", "prod.ci", "prod.live")
	args, err := commands.ParseAndValidateGetCommandArgs(ctx)
	assert.Nil(t, err)
	assert.Len(t, args.UrlKeys, 3)
	assert.Len(t, args.UrlKeys[ww.Global], 1)
	assert.Len(t, args.UrlKeys["local"], 1)
	assert.Len(t, args.UrlKeys["prod"], 2)
	assert.Equal(t, "docs", args.UrlKeys[ww.Global][0])
	assert.Equal(t, "tasks", args.UrlKeys["local"][0])
	assert.Equal(t, "ci", args.UrlKeys["prod"][0])
	assert.Equal(t, "live", args.UrlKeys["prod"][1])
}

func TestParseAndValidateGetCommandArgs_KeyDotSeparated_ReturnsUrlKeys(t *testing.T) {
	ctx := createContextWithArgs(t, "get", "prod.docs")
	args, err := commands.ParseAndValidateGetCommandArgs(ctx)
	assert.Nil(t, err)
	assert.Len(t, args.UrlKeys, 1)
	assert.Len(t, args.UrlKeys["prod"], 1)
	assert.Equal(t, "docs", args.UrlKeys["prod"][0])
}

func TestParseAndValidateGetCommandArgs_KeyDotSeparatedInvalidFormat_ReturnsUrlKeys(t *testing.T) {
	ctx := createContextWithArgs(t, "get", "prod.docs.foo") // Only {env}.{key} supported (for now)
	args, err := commands.ParseAndValidateGetCommandArgs(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected format for key")
	assert.Nil(t, args)
}
