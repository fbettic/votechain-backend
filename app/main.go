package main

import (
	"flag"

	"github.com/fbettic/votechain-backend/app/pipelines"
	"github.com/fbettic/votechain-backend/app/webserver"
	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func main() {

	hydratedConfig := webserver.Config{Port: 8080}

	// Server init
	withData := flag.Bool("withData", false, "initialize the api with some options")
	flag.Parse()

	var options map[string]*dto.Option
	if *withData {
		println("initializing the api with some options")
		options = sampledata.Options
	} else {
		println("initializing the api without any options")
	}

	srv, err := webserver.New(hydratedConfig, options)
	if err != nil {
		panic(err)
	}
	srv.Start(pipelines.BuildPipeline)

}
