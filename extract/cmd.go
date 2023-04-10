package extract

import (
	_ "embed"
	"path/filepath"
	"strings"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `extract`,
	Usage:       `FILE`,
	Version:     `v0.1.0`,
	Copyright:   `Copyright Micah Nadler 2023`,
	Summary:     help.S(_extract),
	Description: help.D(_extract),
	Commands: []*Z.Cmd{
		help.Cmd,
	},

	Call: func(_ *Z.Cmd, args ...string) error {
		p := string(args[0])
		f, _ := filepath.Abs(p)

		switch {
		case strings.Contains(f, ".tar.gz2"):
			extractTarBz2()
		case strings.Contains(f, ".tar.gz"):
			extractTarGz()
		case strings.Contains(f, ".bz2"):
			extractBz2()
		case strings.Contains(f, ".rar"):
			extractRar()
		case strings.Contains(f, ".gz"):
			extractGz()
		case strings.Contains(f, ".tar"):
			extractTar()
		case strings.Contains(f, ".tbz2"):
			extractTbz2()
		case strings.Contains(f, ".tgz"):
			extractTgz()
		case strings.Contains(f, ".zip"):
			extractZip()
		case strings.Contains(f, ".Z"):
			extractZ()
		case strings.Contains(f, ".7z"):
			extract7z()
		}
		return nil
	},
}
