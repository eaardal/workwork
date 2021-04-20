package commands_test

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
	"workwork/src/commands"
)

func TestParseAndValidateGoToCommandArgs_Key_KeyIsSetAndEnvironmentIsEmptyString(t *testing.T) {
	ctx := createContextWithArgs(t, "goto", "docs")

	args, err := commands.ParseAndValidateGoToCommandArgs(ctx)
	assert.Nil(t, err)

	assert.Equal(t, "docs", args.UrlKey)
	assert.Empty(t, args.Environment)
	assert.False(t, args.HasEnvironment)
}

func TestParseAndValidateGoToCommandArgs_KeyAndEnvironment_KeyAndEnvironmentIsSet(t *testing.T) {
	ctx := createContextWithArgs(t, "goto", "prod", "docs")

	args, err := commands.ParseAndValidateGoToCommandArgs(ctx)
	assert.Nil(t, err)

	assert.Equal(t, "docs", args.UrlKey)
	assert.Equal(t, "prod", args.Environment)
	assert.True(t, args.HasEnvironment)
}

func TestParseAndValidateGoToCommandArgs_KeyAndEnvironmentDotSeparated_KeyAndEnvironmentIsSet(t *testing.T) {
	ctx := createContextWithArgs(t, "goto", "prod.docs")

	args, err := commands.ParseAndValidateGoToCommandArgs(ctx)
	assert.Nil(t, err)

	assert.Equal(t, "docs", args.UrlKey)
	assert.Equal(t, "prod", args.Environment)
	assert.True(t, args.HasEnvironment)
}

func TestParseAndValidateGoToCommandArgs_KeyAndEnvironmentDotSeparatedInvalidFormat_ReturnsError(t *testing.T) {
	ctx := createContextWithArgs(t, "goto", "prod.docs.more") // only 2 parts separated by dot allowed (for now)

	args, err := commands.ParseAndValidateGoToCommandArgs(ctx)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected format for key")
	assert.Nil(t, args)
}

func createContextWithArgs(t *testing.T, commandName string, arg ...string) *cli.Context {
	set := flag.NewFlagSet("test", flag.ContinueOnError)
	ctx := cli.NewContext(nil, set, nil)
	ctx.Command = &cli.Command{Name: commandName}
	err := set.Parse(arg)
	assert.Nil(t, err)
	assert.Equal(t, len(arg), ctx.NArg())
	return ctx
}
