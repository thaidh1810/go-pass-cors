package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

var (
	letEncryptFolder = "/etc/letsencrypt/live/"
	projects         = map[string]string{
		"mgmt-dev.raedahgroup.com": "8000",
		"msgs.raedahgroup.com":     "8001",
		"ista.raedahgroup.com":     "8002",
		"ista.vibros.co":     "8002",
		"dcrcare.raedahgroup.com":  "8080",
	}
	port     int
	redirect bool
)

func main() {
	flag.IntVar(&port, "port", 80, "http port. Default is 80")
	flag.BoolVar(&redirect, "redirect", true, "redirect from http to https. Default is true")
	flag.Parse()

	for name, port := range projects {
		vhost, err := url.Parse("http://127.0.0.1:" + port)
		if err != nil {
			panic(err)
		}
		proxy := httputil.NewSingleHostReverseProxy(vhost)
		http.HandleFunc(name+"/", handler(proxy))
	}

	go func() {
		tlsConfig := &tls.Config{}
		files, err := ioutil.ReadDir(letEncryptFolder)
		for _, f := range files {
			if f.IsDir() {
				if cert, err := tls.LoadX509KeyPair(letEncryptFolder+f.Name()+"/cert.pem",
					letEncryptFolder+f.Name()+"/privkey.pem"); err == nil {
					tlsConfig.Certificates = append(tlsConfig.Certificates, cert)
				} else {
					fmt.Println("load cert fail: ", err)
				}
			}
		}

		tlsConfig.BuildNameToCertificate()
		server := &http.Server{
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
			TLSConfig:      tlsConfig,
		}
		listener, err := tls.Listen("tcp", ":443", tlsConfig)

		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(server.Serve(listener))
	}()
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		panic(err)
	}

}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteURL, err  := url.Parse(r.Header.Get("Remote-URL"))
		proxy := httputil.NewSingleHostReverseProxy(remoteURL)
		proxy.ServeHTTP(w, r)
	}
}
