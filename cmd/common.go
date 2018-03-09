package cmd

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"os"
	"encoding/csv"
)

// selectInput determines whether we should read a file from disk or stdin
func selectInput(c *cli.Context) (*csv.Reader, error) {
	if c.NArg() > 0 {
		reader, err := os.Open(c.Args().Get(c.NArg() - 1))

		if err != nil {
			return nil, errors.Wrap(err, "could not open input file")
		}

		return csv.NewReader(reader), nil
	}
	return csv.NewReader(os.Stdin), nil
}
