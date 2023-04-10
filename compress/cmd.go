package compress

import (
	_ "embed"
	"fmt"
	"os"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
)

var Cmd = &Z.Cmd{
	Name:        `compress`,
	Usage:       `COMMAND [help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_compress),
	Description: help.D(_compress),

	Commands: []*Z.Cmd{
		compressTarGz,
		compressZip,
		help.Cmd,
	},
}

var compressTarGz = &Z.Cmd{
	Name:        `targz`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_compresstargz),
	Description: help.D(_compresstargz),

	MinArgs: 1,

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(_ *Z.Cmd, args ...string) error {
		var outName string

		fmt.Println("Output file: ")
		fmt.Scanf("%s", &outName)
		outputFile := outName + ".tar.gz"

		listOfFilesNames := args
		outFile, err := os.Create(outputFile)
		if err != nil {
			panic("Error creating archive file")
		}
		defer outFile.Close()
		err = CreateTarAndGz(listOfFilesNames, outFile)
		if err != nil {
			panic("Error creating archive file.")
		}

		fmt.Println("Archiving and file compression completed.")
		return nil
	},
}

var compressZip = &Z.Cmd{
	Name:        `zip`,
	Usage:       `[help]`,
	Version:     `v0.0.1`,
	Copyright:   `Copyright Micah Nadler 2023`,
	License:     `Apache-2.0`,
	Summary:     help.S(_compresszip),
	Description: help.D(_compresszip),

	Commands: []*Z.Cmd{help.Cmd},

	Call: func(x *Z.Cmd, args ...string) error {
		var outName string

		fmt.Println("Output file: ")
		fmt.Scanf("%s", &outName)
		outputFile := outName + ".zip"

		CompressZip(args, outputFile)
		return nil
	},
}
