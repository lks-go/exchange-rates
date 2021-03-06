package app

import (
	"fmt"

	"github.com/lks-go/exchange-rates/pkg/exchangeratesapi"
)

type APICient interface {
	Rates(symbols ...string) (*exchangeratesapi.RatesResponse, error)
}

func Run(args *Args, api APICient) (string, error) {

	r, err := api.Rates(args.From, args.To)
	if err != nil {
		return "", err
	}

	if _, ok := r.Rates[args.From]; !ok {
		return "", fmt.Errorf("symbol %s not available", args.From)
	}

	if _, ok := r.Rates[args.To]; !ok {
		return "", fmt.Errorf("symbol %s not available", args.To)
	}

	result := (1 / r.Rates[args.From]) / (1 / r.Rates[args.To]) * args.Amount

	return fmt.Sprintf("%.2f", result), nil
}
