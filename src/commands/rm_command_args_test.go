package commands_test

import (
	"github.com/eaardal/workwork/src/commands"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAndValidateRMCommandArgs_NoArgs_ReturnsError(t *testing.T) {
	ctx := createContextWithArgs(t, "rm")
	args, err := commands.ParseAndValidateRMCommandArgs(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no url keys specified")
	assert.Nil(t, args)
}

func TestParseAndValidateRMCommandArgs_WithArgs_ReturnsUrlKeysToDelete(t *testing.T) {
	ctx := createContextWithArgs(t, "rm", "docs", "tasks", "issues")
	args, err := commands.ParseAndValidateRMCommandArgs(ctx)
	assert.Nil(t, err)
	assert.Len(t, args.UrlKeysToBeDeleted, 3)
	assert.Equal(t, "docs", args.UrlKeysToBeDeleted[0])
	assert.Equal(t, "tasks", args.UrlKeysToBeDeleted[1])
	assert.Equal(t, "issues", args.UrlKeysToBeDeleted[2])
}
