package cmd

import (
	"bufio"
	"github.com/urfave/cli"
	"os"
	"github.com/pkg/errors"
)

// selectInput determines whether we should read a file from disk or stdin
func selectInput(c *cli.Context) (*bufio.Scanner, error) {
	if c.NArg() > 0 {
		reader, err := os.Open(c.Args().Get(c.NArg() - 1))

		if err != nil {
			return nil, errors.Wrap(err, "could not open input file")
		}

		return bufio.NewScanner(reader), nil
	}
	return bufio.NewScanner(os.Stdin), nil
}
