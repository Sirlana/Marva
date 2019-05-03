package libs

import (
	"net/http"

	"github.com/gorilla/mux"
	"sirlana.com/sirlana/sso/db"
)

type Marva struct {
	util *Util
	log  *Logger
	jwt  *JWT
	r    *mux.Router
	db   db.Database
	fun  map[string]func()
	funi map[string]func() int
	funs map[string]func() string
	funf map[string]func() float64
	funb map[string]func() bool
}

func NewMarva(util *Util, log *Logger, jwt *JWT, r *mux.Router, db db.Database) *Marva {
	return &Marva{
		util: util,
		log:  log,
		jwt:  jwt,
		r:    r,
		db:   db,
	}
}

func (m Marva) Run(file string, isPrivate bool) {
	data, err := m.util.JSONData(file)
	if err != nil {
		m.log.E(err.Error())
		panic(err.Error())
	}
	req := data["request"].(map[string]interface{})

	m.r.HandleFunc(req["path"].(string), func(w http.ResponseWriter, r *http.Request) {
		// Parse params
		reqParam := req["param"].(map[string]interface{})
		_ = m.util.ParseParam(r, reqParam["type"].(string), reqParam["data"].(string))

		if isPrivate {
			token := r.FormValue("X-Session-Token")
			if m.jwt.isValid(token) {

			} else {
				m.util.ShowCustomErrorResponse(w, 401, "Invalid token.")
			}
		} else {

		}
	}).Methods(req["type"].(string))
}

func (m Marva) RegisterFunction(id string, f func()) {
	m.fun[id] = f
}

func (m Marva) RegisterIntFunction(id string, f func() int) {
	m.funi[id] = f
}

func (m Marva) RegisterStringFunction(id string, f func() string) {
	m.funs[id] = f
}

func (m Marva) RegisterFloatFunction(id string, f func() float64) {
	m.funf[id] = f
}

func (m Marva) RegisterBoolFunction(id string, f func() bool) {
	m.funb[id] = f
}
