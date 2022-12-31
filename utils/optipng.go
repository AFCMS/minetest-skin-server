package utils

import (
	"io"
	"os"
	"os/exec"
)

// Return true if OptiPNG is present on the system
//
// Try to execute `optipng --version`
func OptiPNGPresent() bool {
	c := exec.Command("optipng", "--version")
	c.Stdout = nil // /dev/null

	return c.Run() == nil
}

// Run OptiPNG on a byte array
//
// Make use of temporary files
func OptiPNGBytes(input []byte) (output []byte, err error) {
	// Create a temporary file
	tmp_file, err := os.CreateTemp("", "optipng-")
	tmp_file.Write(input)

	defer tmp_file.Close()
	defer os.Remove(tmp_file.Name())

	// Run OptiPNG on the image in place
	cmd := exec.Command("optipng", "-o7", "-zm1-9", "-nc", "-strip", "all", "-clobber", tmp_file.Name())

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
	tmp_file_r, err := os.Open(tmp_file.Name())

	if err != nil {
		return nil, err
	}

	defer tmp_file_r.Close()

	// Read the file content, return output
	output, err = io.ReadAll(tmp_file_r)
	if err != nil {
		return nil, err
	}

	return output, nil
}
