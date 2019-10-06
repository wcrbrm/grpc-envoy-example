package main

import (
	"github.com/golang/protobuf/jsonpb"

	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "github.com/grpc-ecosystem/grpc-gateway/utilities"

	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
	pb "github.com/wcrbrm/grpc-envoy-example/core/test"
)

var app = cli.App("auth", "example gRPC service for authentication")

func main() {
	// log.SetLevel(logging.Level(*appLogLevel))
	// log.SetFormatter(logging.DefaultFormatter())

	app.Action = start
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("[ERR]", err)
	}
}

func testEnums() {
	marshaller := jsonpb.Marshaler{}

	paginationJSON, _ := marshaller.MarshalToString(&pb.DisplayLocale{
		Lang: pb.Lang_nl,
	})
	log.Printf("locale: %v", paginationJSON)

	jsonMsg := "{\"lang\": \"fr\"}"
	var locale pb.DisplayLocale
	if err := jsonpb.UnmarshalString(jsonMsg, &locale); err != nil {
		log.WithError(err).Errorf("decode")
	}
	log.WithField("locale", locale.String()).Info()
}

func testOneOf() {
	marshaller := jsonpb.Marshaler{}
	paginationJSON, _ := marshaller.MarshalToString(&pb.DisplayLocale{
		Lang: pb.Lang_nl,
	})
	log.Printf("oneOf: %v", paginationJSON)
}
