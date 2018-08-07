package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func buildTar(src string) (*bytes.Buffer, error) {
	// new buffer
	buf := new(bytes.Buffer)
	tarWriter := tar.NewWriter(buf)
	defer tarWriter.Close()

	if _, err := os.Stat(src); err != nil {
		return nil, fmt.Errorf("Unable to tar files - %v", err.Error())
	}

	err := filepath.Walk(src, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// create file header
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// write the header
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// open files for taring
		f, err := os.Open(file)
		defer f.Close()
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tarWriter, f); err != nil {
			return err
		}

		return nil

		// remove the directoryname from the file
		// is this needed
		// nakedFile := strings.Replace(file, src, "", -1)
		// reallyNakedFile := strings.TrimPrefix(nakedFile, string(filepath.Separator))
	})
	if err != nil {
		return nil, err
	}

	return buf, nil
}
