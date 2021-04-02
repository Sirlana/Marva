package main

import (
	"net/http"

	"sirlana.com/sirlana/sso/app"
	"sirlana.com/sirlana/sso/db"
	"sirlana.com/sirlana/sso/libs"

	"github.com/gorilla/mux"
)

func main() {
	// Load config file and log.
	r := mux.NewRouter()
	util := libs.NewUtil()
	log := libs.NewLogger()
	cfg := libs.NewConfig("config.json", util, log)

	dbase := cfg.Database()
	jwt := libs.NewJWT(cfg.API()["private"].(map[string]interface{})["token_key"].(string))
	mysql := db.NewMySQL(dbase["username"].(string), dbase["password"].(string), dbase["dbname"].(string))

	marva := libs.NewMarva(util, log, jwt, r, mysql)

	// Open database connection.
	if dbase["driver"].(string) == "mysql" {
		if err := mysql.Connect(); err != nil {
			log.E("Database connection error.")
			panic(err)
		}
	} else {
		log.E("Database driver undefined.")
		panic("Error database driver name.")
	}

	// Init router.
	mux.CORSMethodMiddleware(r)
	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./www/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/sir/").Handler(http.StripPrefix("/sir/", http.FileServer(http.Dir("./www/"))))

	// Run Services
	app.Register(marva)

	apiAuth := cfg.API()["private"].(map[string]interface{})["auth"].([]interface{})
	for _, val := range apiAuth {
		marva.Run(val.(string), true)
	}

	apiPrivate := cfg.API()["private"].(map[string]interface{})["endpoints"].([]interface{})
	for _, val := range apiPrivate {
		marva.Run(val.(string), true)
	}

	apiPublic := cfg.API()["public"].(map[string]interface{})["endpoints"].([]interface{})
	for _, val := range apiPublic {
		marva.Run(val.(string), false)
	}

	// Load services.
	if err := cfg.LoadTLSServices(r); err != nil {
		log.E("Service Error.")
		panic(err.Error())
	}

	defer mysql.Close()
}
