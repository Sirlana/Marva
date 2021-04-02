package libs

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	Database struct {
		Driver   string `json:"driver"`
		Username string `json:"username"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		DBName   string `json:"dbname"`
	} `json:"database"`
	Migration struct {
		Driver string `json:"driver"`
		Folder string `json:"folder"`
	} `json:"migration"`
	Server struct {
		Host        string `json:"host"`
		Port        string `json:"port"`
		Certificate string `json:"certificate"`
		Privkey     string `json:"privkey"`
	} `json:"server"`
}

func NewConfig() (*Config, error) {
	file, _ := os.Open("config.sir")
	defer file.Close()
	cfg := Config{}
	err := json.NewDecoder(file).Decode(&cfg)
	return &cfg, err
}

func (cfg Config) LoadTLSServices(r *mux.Router) error {
	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		TLSConfig:    tlsConfig,
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	return srv.ListenAndServeTLS(cfg.Server.Certificate, cfg.Server.Privkey)
}

func (cfg Config) LoadServices(r *mux.Router) error {
	srv := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	return srv.ListenAndServe()
}
