package utils

import (
	"bytes"
	"os/exec"
)

func OptiPNGPresent() bool {
	c := exec.Command("optipng", "--version")
	c.Stdout = nil // /dev/null

	return c.Run() == nil
}

func OptiPNGBytes(b []byte) []byte {
	c := exec.Command("optipng")

	var out []byte

	c.Stdin = bytes.NewReader(b)
	c.Stdout = bytes.NewBuffer(out)

	return out
}
