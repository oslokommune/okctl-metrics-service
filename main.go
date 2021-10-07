/*
 * okctl-metrics-backend
 *
 * A service for collecting metrics in okctl
 *
 * API version: 0.0.1
 * Contact: okctl@oslo.kommune.no
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/oslokommune/okctl-metrics-service/pkg/config"

	sw "github.com/oslokommune/okctl-metrics-service/pkg/router"
)

//go:embed specification.yaml
var specification []byte

func main() {
	log.Printf("Server started")

	cfg, err := config.Generate()
	if err != nil {
		panic(err.Error())
	}

	err = cfg.Validate()
	if err != nil {
		panic(err.Error())
	}

	router := sw.New(cfg, specification)

	println("delete me")

	log.Fatal(router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
