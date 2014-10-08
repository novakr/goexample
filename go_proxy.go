package main

import (
	//"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = ""
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	log.Println("[Request URL]:", r.URL)
	log.Println("[Request Host]:", r.Host)
	log.Println("[Requst Header]:", r.Header)
	log.Println("[Requst Cookies]:", r.Cookies())

	if "www.baidu.com" == r.Host {
		http.Redirect(w, r, "/www.google.com", http.StatusFound)
	}

	//resp, err := http.DefaultClient.Do(r)
	resp, err := client.Do(r)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal("req do", err)
	}
	log.Println("[Response Header]:", resp.Header)
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}

	//log.Println("[Response Cookies]:", resp.Cookies())
	//for _, c := range resp.Cookies() {
	//	w.Header().Add("Set-Cookie", c.Raw)
	//}

	//_, ok := resp.Header["Content-Length"]
	//if !ok && resp.ContentLength > 0 {
	//	w.Header().Add("Content-Length", fmt.Sprint(resp.ContentLength))
	//}

	w.WriteHeader(resp.StatusCode)

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		log.Fatal("resp body", err)
	}
	w.Write(result)
}
func proxy(ch chan bool) {
	http.HandleFunc("/", handler)
	log.Println("Start serving on port 8888")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		ch <- false
	}
}

func main() {
	ch := make(chan bool)
	go proxy(ch)
	f := <-ch
	if false == f {
		log.Fatal(" failed")
	}
	os.Exit(0)
}
