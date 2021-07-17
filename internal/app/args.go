package app

import (
	"fmt"
	"os"
	"strconv"
)

type Args struct {
	Amount float64
	From   string
	To     string
}

// ParseArgs parses stdin arguments to struct Args
func ParseArgs() (*Args, error) {

	if len(os.Args) < 4 {
		return nil, fmt.Errorf("not enought arguments")
	}

	amount, err := strconv.ParseFloat(os.Args[1], 2)
	if err != nil {
		return nil, fmt.Errorf("can't parse amount: %s", err)
	}

	//for _, arg := range os.Args[2:4] {
	//	symbol := currency.Symbol(arg)
	//
	//	if !currency.Available(symbol) {
	//		return nil, fmt.Errorf("symbol %s not available", symbol)
	//	}
	//}

	pa := &Args{
		Amount: amount,
		From:   os.Args[2],
		To:     os.Args[3],
	}

	return pa, nil
}
