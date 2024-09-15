package fs

import (
	"io"
	"os"
)

func MoveFile(src, dst string) error {

	// open source
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// create dest
	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// copy contents
	if _, err = io.Copy(destinationFile, sourceFile); err != nil {
		return err
	}

	// remove source
	if err = os.Remove(src); err != nil {
		return err
	}

	return nil
}
