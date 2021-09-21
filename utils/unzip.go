package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func UnZip(src string, dest string) error {
	fmt.Printf("Unpacking file %s...\n", src)

	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename for using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath)
		}

		if f.FileInfo().IsDir() {
			// Make Folder
			fmt.Printf("\tCreating folder %s...\n", fpath)
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		fmt.Printf("\tInflating file %s...\n", fpath)
		err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
		if err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
