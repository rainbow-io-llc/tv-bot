package ngrok

import (
	"context"
	"net"
	"os"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

const (
	USRNAME_ENV_VAR = "NGROK_USRNAME"
	PASSWD_ENV_VAR = "NGROK_PASSWD"
	TOKEN_ENV_VAR = "NGROK_AUTHTOKEN"
)

var Tunnel tunnel

type tunnel struct {
	username string
	password string
}

func init() {
	token := os.Getenv(TOKEN_ENV_VAR)
	if token == "" {
		panic("[NGROK-error] no token setup for tunnel connect")
	}

	username := os.Getenv(USRNAME_ENV_VAR)
	password := os.Getenv(PASSWD_ENV_VAR)
	if username == "" || password == "" {
		panic("[NGROK-error] no username or password setup for basic auth")
	}

	Tunnel.username = username	
	Tunnel.password = password
}

func (t *tunnel) Get(ctx context.Context) (net.Listener, error) {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithBasicAuth(t.username, t.password),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return nil, err
	}
	return tun, nil
}