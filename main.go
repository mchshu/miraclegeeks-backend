package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var configJSON = flag.String("conf", "config.json", "config file location")

func main() {
	flag.Parse()

	log.SetFlags(log.Lshortfile | log.Ltime)

	data, err := ioutil.ReadFile(*configJSON)
	if err != nil {
		log.Fatalf("load config err: %s", err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		log.Fatalf("load config err: %s", err)
	}

	initDB(config.Driver, config.Dsn)

	http.HandleFunc("/api/login", HandleLogin)
	http.HandleFunc("/api/register", HandleRegsiter)
	http.HandleFunc("/api/publicProjects", HandleProjects)
	http.HandleFunc("/api/enterpriseProjects", HandleProjects)
	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Printf("listen %s", config.Addr)

	http.ListenAndServe(config.Addr, nil)
}
