package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
)

func Headings(c *cli.Context) error {
	reader, err := selectInput(c)
	if err != nil {
		return errors.Wrap(err, "Select input failed")
	}

	headings, err := reader.Read()
	if err != nil {
		return errors.Wrap(err, "Read line failed")
	}

	for i, el := range headings {
		fmt.Printf("%d: %s\n", i+1, el)
	}

	return nil
}
