package libs

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	util *Util
	log  *Logger
	data map[string]interface{}
}

func NewConfig(file string, util *Util, log *Logger) *Config {
	data, err := util.JSONData(file)
	if err != nil {
		log.E(err.Error())
		panic(err.Error())
	}
	return &Config{
		util: util,
		log:  log,
		data: data,
	}
}

func (cfg Config) Database() map[string]interface{} {
	return cfg.data["database"].(map[string]interface{})
}

func (cfg Config) Server() map[string]interface{} {
	return cfg.data["server"].(map[string]interface{})
}

func (cfg Config) API() map[string]interface{} {
	return cfg.data["api"].(map[string]interface{})
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
	server := cfg.Server()
	srv := &http.Server{
		Addr:         server["host"].(string) + ":" + server["port"].(string),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		TLSConfig:    tlsConfig,
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	return srv.ListenAndServeTLS(server["certificate"].(string), server["privkey"].(string))
}

func (cfg Config) LoadServices(r *mux.Router) error {
	server := cfg.Server()
	srv := &http.Server{
		Addr:         server["host"].(string) + ":" + server["port"].(string),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handlers.LoggingHandler(os.Stdout, r),
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	return srv.ListenAndServe()
}
