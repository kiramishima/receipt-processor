package main

import (
	"github.com/kiramishima/receipt-processor/bootstrap"

	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
