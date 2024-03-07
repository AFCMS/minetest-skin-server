package utils

import (
	"io"
	"os"
	"os/exec"
)

// OptiPNGPresent Return true if OptiPNG is present on the system
//
// Try to execute `optipng --version`
func OptiPNGPresent() bool {
	c := exec.Command("optipng", "--version")
	c.Stdout = nil // /dev/null

	return c.Run() == nil
}

// OptiPNGBytes Run OptiPNG on a byte array
//
// Make use of temporary files
func OptiPNGBytes(input []byte) (output []byte, err error) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "optipng-")
	tmpFile.Write(input)

	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// Run OptiPNG on the image in place
	cmd := exec.Command("optipng", "-o7", "-zm1-9", "-nc", "-strip", "all", "-clobber", tmpFile.Name())

	// Redirect the output to the null device
	cmd.Stdout = nil
	cmd.Stderr = nil

	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	err = cmd.Run()

	if err != nil {
		return nil, err
	}

	// Open the temporary file in read mode
	tmpFileR, err := os.Open(tmpFile.Name())

	if err != nil {
		return nil, err
	}

	defer tmpFileR.Close()

	// Read the file content, return output
	output, err = io.ReadAll(tmpFileR)
	if err != nil {
		return nil, err
	}

	return output, nil
}
