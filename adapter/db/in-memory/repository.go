package in_memory

import "go.uber.org/fx"

var Module = fx.Module("db",
	fx.Provide(NewReceiptRepository),
)
