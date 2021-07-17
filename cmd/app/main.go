package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lks-go/exchange-rates/pkg/exchangeratesapi"

	"github.com/lks-go/exchange-rates/internal/app"
)

func main() {
	args, err := app.ParseArgs()
	if err != nil {
		log.Fatalf("invalid arguments: %s", err)
	}

	cfg := app.Config{}

	if err := app.Configure(&cfg); err != nil {
		log.Fatalf("failed to configure: %s", err)
	}

	api := exchangeratesapi.New(&cfg.API)

	res, err := app.Run(args, api)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(os.Stdout, res)
}
