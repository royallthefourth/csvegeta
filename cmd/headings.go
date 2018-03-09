package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"strconv"
	"strings"
)

func Headings(c *cli.Context) error {
	scanner, err := selectInput(c)

	if err != nil {
		return errors.Wrap(err, "Select input failed")
	}

	scanner.Scan()
	line := scanner.Text()

	headings := strings.Split(line, ",")

	for i, el := range headings {
		if strings.ContainsAny(el, "\"'") {
			el, _ = strconv.Unquote(el)
		}
		fmt.Printf("%d: %s\n", i, el)
	}

	return nil
}
