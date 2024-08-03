// Package traefikrealip
package traefikrealip

import (
	"context"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct{}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}


type RealIpPlugin struct {
	next http.Handler
	name string
}


// New created a new RealIpPlugin plugin.
func New(ctx context.Context, next http.Handler, _ *Config, name string) (http.Handler, error) {
	return &RealIpPlugin{
		next: next,
		name: name,
	}, nil
}

func (a *RealIpPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	xForwardedFor := req.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			req.Header.Set("X-Real-IP", strings.TrimSpace(ips[0]))
		}
	}
	m.next.ServeHTTP(rw, req)
}
