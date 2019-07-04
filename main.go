package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	apiHost string
	apiPort string
	debug   bool
)

func init() {
	flag.StringVar(&apiHost, "host", "localhost", "Hostname")
	flag.StringVar(&apiPort, "port", "8000", "Port")
	flag.BoolVar(&debug, "debug", false, "Debug mode")
}

func main() {
	flag.Parse()

	router := mux.NewRouter()

	router.HandleFunc("/", RedirectToChloe).Queries("name", "{name}")
	router.HandleFunc("/", Welcome)
	router.HandleFunc("/d/{name}", Chloe)
	router.HandleFunc("/{name}", Chloe)

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", apiHost, apiPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting Chloe on http://%s:%s\n", apiHost, apiPort)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func RedirectToChloe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if debug {
		fmt.Printf("%+v\n", r)
	}
	if fmt.Sprintf("%s:%s", apiHost, apiPort) == r.Host {
		http.Redirect(w, r, fmt.Sprintf("/%s", vars["name"]), 301)
	} else {
		http.Redirect(w, r, fmt.Sprintf("https://a.cow.name/d/%s", vars["name"]), 301)
	}
}

func Chloe(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Println(vars["name"])
	buffer := new(bytes.Buffer)
	Stamp(buffer, vars["name"])
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func Welcome(writer http.ResponseWriter, request *http.Request) {
	b, _ := Asset("assets/index.html")
	writer.Write(b)
}
