package compress

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"io"
	"os"

	"github.com/charmbracelet/log"
)

func CreateTarAndGz(fileNames []string, buffer io.Writer) error {
	gzipWriter := gzip.NewWriter(buffer)
	defer gzipWriter.Close()
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()
	for _, f := range fileNames {
		err := addToTar(tarWriter, f)
		if err != nil {
			return err
		}
	}
	return nil
}

func addToTar(tarWriter *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}
	header.Name = filename
	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}

func CompressZip(filenames []string, outFile string) error {

	zipArchive, _ := os.Create(outFile)
	defer zipArchive.Close()
	writer := zip.NewWriter(zipArchive)

	for _, file := range filenames {
		f, err := os.Open(file)
		if err != nil {
			log.Error(err)
		}
		defer f.Close()
		w, err := writer.Create(file)
		if err != nil {
			log.Error(err)
		}
		io.Copy(w, f)

	}
	writer.Close()

	return nil
}
