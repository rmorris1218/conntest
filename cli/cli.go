package cli

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/rmorris1218/conntest/internal"
)

func GetCli() *cli.App {
	return &cli.App{
		Name: "conntest",
    	Usage: "test network connectivity to different endpoint/ports based off a config file",
		Flags: []cli.Flag {
			&cli.StringFlag{
			  Name: "config-file",
			  Usage: "path of conntest config file",
			},
		  },
		Action: func(c *cli.Context) error {

			testCases, err := internal.ParseTestFile(c.String("config-file"))
			if err != nil {
				return cli.Exit(fmt.Errorf("%v", err), 1)
			}
			for _, endpoint := range testCases.Tests {
				isReachable := endpoint.TestReachability()
				if !isReachable {
					errMsg := fmt.Sprintf("error: could not reach %s", endpoint.Uri)
					return cli.Exit(errMsg, 1)
				}
				fmt.Printf("successfully reached all ports on %s", endpoint.Uri)
			}
			return nil
		},

	}

}
