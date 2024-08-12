package ngrok

import (
	"context"
	"net"
	"os"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

const CLIENT_CN = "webhook-server@tradingview.com"

var Plain plainTunnel 

type plainTunnel struct {
}

func init() {
	token := os.Getenv(TOKEN_ENV_VAR)
	if token == "" {
		panic("[NGROK-error] no token setup for tunnel connect")
	}
}

func (t *plainTunnel) Get(ctx context.Context) (net.Listener, error) {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return nil, err
	}
	return tun, nil
}