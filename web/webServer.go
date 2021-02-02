package web

import (
	"fmt"
	"net/http"
	"strings"
)

type Config struct {
	ServerParams struct {
		ExternalServer string `yaml:"externalServer"`
		ServerPort     string `yaml:"serverPort"`
		PathToPic      string `yaml:"pathToPic"`
		PathToHTML     string `yaml:"pathToHTML"`
	} `yaml:"serverParams"`
}

var Configuration Config

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if strings.Contains(r.URL.RequestURI(), "email=") {
			http.ServeFile(w, r, Configuration.ServerParams.PathToHTML)
			fmt.Println(r.URL.RequestURI())
			path := Configuration.ServerParams.ExternalServer + r.URL.RequestURI()
			_, err := http.Get(path)
			if err != nil {
				fmt.Println("http.Get error: ", err)
			}
		} else if strings.Contains(r.URL.RequestURI(), "picture=") {
			fmt.Println(r.URL.RequestURI())
			path := Configuration.ServerParams.ExternalServer + r.URL.RequestURI()
			_, err := http.Get(path)
			if err != nil {
				fmt.Println("http.Get error: ", err)
			}
			http.Redirect(w, r, "/pic", 301)
		} else {
			http.ServeFile(w, r, Configuration.ServerParams.PathToHTML)
		}
	} else {
		http.ServeFile(w, r, Configuration.ServerParams.PathToHTML)
	}
}

func PicHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.RequestURI())
	http.ServeFile(w, r, Configuration.ServerParams.PathToPic)

	fmt.Printf("Request connection: %s, path: %s\n", r.Proto, r.URL.Path[1:])
}

func WebServer() {
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/pic", PicHandler)

	fmt.Println("Server is listening on port:", Configuration.ServerParams.ServerPort)
	http.ListenAndServe(":" + Configuration.ServerParams.ServerPort, nil)
}