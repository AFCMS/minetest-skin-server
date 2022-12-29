package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func OptiPNGPresent() bool {
	c := exec.Command("optipng", "--version")
	c.Stdout = nil // /dev/null

	return c.Run() == nil
}

func OptiPNGBytes(input []byte) (output []byte, err error) {
	cmd := exec.Command("optipng", "")
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	err = cmd.Run()

	if err != nil {
		return nil, err
	}

	output = o.Bytes()
	return output, nil
}
