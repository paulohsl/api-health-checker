package main

import (
	"github.com/paulohsl/api-health/healthcheck"
	"github.com/paulohsl/api-health"
)

func main() {

	urls := []string{"http://www.terra.com.br", "http://www.moip.com.br"}

	hc := apihealth.HttpCaller{}

	health.Check(hc, urls)
}
