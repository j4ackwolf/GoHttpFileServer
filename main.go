package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Config file not specified")
		os.Exit(1)
	}
	if os.Args[1] != "-c" {
		fmt.Println("Wrong parameter, use -c")
		os.Exit(1)
	}
	if os.Args[2] == "" {
		fmt.Println("Config file not specified")
		os.Exit(1)
	}
	err := LoadConfig(os.Args[2])
	if err != nil {
		fmt.Println("Warning: config file not found, using default config")
		cfg = DefaultConfig()
		err = SaveConfig(os.Args[2])
		if err != nil {
			fmt.Println("Error: unable to save default config")
			os.Exit(1)
		}
	}

	// Api handlers
	http.HandleFunc(endpoint, apiHandler)
	http.HandleFunc(endpoint+"/upload", apiHandlerUpload)
	http.HandleFunc(endpoint+"/download", apiHandlerUpload)
	http.HandleFunc(endpoint+"/copy", apiHandlerCopy)
	http.HandleFunc(endpoint+"/move", apiHandlerMove)

	// handler for the root path "/"
	http.Handle("/", http.FileServer(http.Dir(cfg.Workdir)))

	// Start the server
	if cfg.Tls {
		err = http.ListenAndServeTLS(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port), cfg.ServerCert, cfg.ServerKey, nil)
	} else {
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Address, cfg.Port), nil)
	}
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}
