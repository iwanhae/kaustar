package proxy

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/iwanhae/kaustar/pkg/config"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/proxy"
)

type Config struct {
	Server      *url.URL
	ClientCerts []tls.Certificate
	CAPool      *x509.CertPool
}

func LoadConfig(cfg config.KubernetesConfig) (Config, error) {
	certs, err := tls.LoadX509KeyPair(cfg.FrontProxyClientCert, cfg.FrontProxyClientKey)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to load front proxy certificates")
	}
	capool := x509.NewCertPool()

	if b, err := ioutil.ReadFile(cfg.CACert); err != nil {
		return Config{}, errors.Wrap(err, "failed to read CA certificate")
	} else {
		if !capool.AppendCertsFromPEM(b) {
			return Config{}, errors.Wrap(err, "invalid CA certificate")
		}
	}

	server, err := url.Parse(cfg.Server)
	if err != nil {
		return Config{}, errors.Wrap(err, "failed to parse kubernetes server URL")
	}
	return Config{
		Server:      server,
		ClientCerts: []tls.Certificate{certs},
		CAPool:      capool,
	}, nil
}

func KubeAPIServer(cfg Config) echo.HandlerFunc {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			Certificates: cfg.ClientCerts,
			RootCAs:      cfg.CAPool,
		},
	}
	h := proxy.NewUpgradeAwareHandler(cfg.Server, transport, false, false, &responder{})
	h.UseRequestLocation = true
	h.UseLocationHost = true
	h.AppendLocationPath = true
	return func(c echo.Context) error {
		req := c.Request().Clone(context.Background())
		req.Header.Set("X-Remote-User", "test")
		req.Header.Set("X-Remote-Group", "system:masters")
		h.ServeHTTP(c.Response(), req)
		return nil
	}
}

type responder struct{}

func (r *responder) Error(w http.ResponseWriter, req *http.Request, err error) {
	log.Printf("Error while proxying request: %v", err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
