package file

import (
	_ "embed"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/tr00datp00nar/file/compress"
	"github.com/tr00datp00nar/file/extract"
)

var Cmd = &Z.Cmd{
	Name:        `file`,
	Summary:     help.S(_file),
	Description: help.D(_file),
	Usage:       `COMMAND`,
	Version:     `v0.1.0`,
	License:     `Apache-2.0`,
	Source:      `git@github.com:tr00datp00nar/c/find.git`,
	Copyright:   `Copyright Micah Nadler 2023`,
	Issues:      `github.com/tr00datp00nar/c/issues`,
	Commands: []*Z.Cmd{
		compress.Cmd,
		extract.Cmd,
		help.Cmd,
	},
}
