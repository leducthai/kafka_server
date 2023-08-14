package server

import (
	"log"
	"os"
	"producer/server/config"

	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
)

func parseFlags() *config.Options {
	var conf config.Options

	configuations := []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Server Configuation",
			LongDescription:  "Server Configuation",
			Options:          &conf,
		},
	}

	parser := flags.NewParser(nil, flags.Default)
	for _, optGroup := range configuations {
		if _, err := parser.AddGroup(optGroup.ShortDescription, optGroup.LongDescription, optGroup.Options); err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok && fe.Type == flags.ErrHelp {
			code = 0
		}
		os.Exit(code)
	}
	return &conf
}
