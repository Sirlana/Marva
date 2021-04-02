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
	cfg, err := libs.NewConfig()
	log := libs.NewLogger()
	if err != nil {
		log.E("Error load config file.")
		panic(err)
	}

	// Init mysql
	mysql := db.NewMySQL(cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)

	// Open database connection.
	if cfg.Database.Driver == "mysql" {
		if err := mysql.Connect(); err != nil {
			log.E("Database connection error.")
			panic(err)
		}
	} else {
		log.E("Database driver undefined.")
		panic("Error database driver name.")
	}

	// Init router.
	r := mux.NewRouter()
	mux.CORSMethodMiddleware(r)
	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./www/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/sir/").Handler(http.StripPrefix("/sir/", http.FileServer(http.Dir("./www/"))))

	// Run Services
	app.Run(r, mysql.GetDB())

	// Load services.
	if err := cfg.LoadTLSServices(r); err != nil {
		log.E("Service Error.")
		panic(err.Error())
	}

	defer mysql.Close()
}
