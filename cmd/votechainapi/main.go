package main

//PUTO EL QUE LEE

import (
	"flag"
	"log"
	"net/http"

	sample "github.com/fbettic/votechain-backend/cmd/sample-data" // import the sample package
	option "github.com/fbettic/votechain-backend/pkg"             // import the option package
	"github.com/fbettic/votechain-backend/pkg/server"             // import the server package
	"github.com/fbettic/votechain-backend/pkg/storage/inmem"
)

func main() {
	withData := flag.Bool("withData", false, "initialize the api with some options")
	flag.Parse()

	var options map[string]*option.Option
	if *withData {
		println("initializing the api with some options")
		options = sample.Options
	}else{
		println("initializing the api without any options")
	}

	repository := inmem.NewOptionRepository(options)

	s := server.New(repository)
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
