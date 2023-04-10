package extract

import (
	scriptish "github.com/ganbarodigital/go_scriptish"
)

func extractZip() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("unzip", "$1"),
	)
	pipeline.Exec()
}

func extractTar() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tar", "xf", "$1"),
	)
	pipeline.Exec()
}

func extractTarGz() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tar", "xzf", "$1"),
	)
	pipeline.Exec()
}

func extractTarBz2() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tar", "xjf", "$1"),
	)
	pipeline.Exec()
}

func extractBz2() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("bunzip", "$1"),
	)
	pipeline.Exec()
}

func extractRar() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("rar", "r", "$1"),
	)
	pipeline.Exec()
}

func extractGz() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("gunzip", "$1"),
	)
	pipeline.Exec()
}

func extractTbz2() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tar", "xjf", "$1"),
	)
	pipeline.Exec()
}

func extractTgz() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("tar", "xzf", "$1"),
	)
	pipeline.Exec()
}

func extractZ() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("uncompress", "$1"),
	)
	pipeline.Exec()
}

func extract7z() {
	pipeline := scriptish.NewPipeline(
		scriptish.Exec("7z", "x", "$1"),
	)
	pipeline.Exec()
}
