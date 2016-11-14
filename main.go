package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var filedir string

func init() {
	flag.StringVar(&filedir, "d", "", "Directory of files to share")
}

func main() {

	flag.Parse()

	l, err := net.Listen("tcp", ":4444")

	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/list", fileLister)
	mux.HandleFunc("/cast", caster)

	mux.Handle("/f/", http.StripPrefix("/f/", http.FileServer(http.Dir(filedir))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	log.Print(http.Serve(l, mux))

}

func fileLister(w http.ResponseWriter, r *http.Request) {

	list, err := listFiles(filedir)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Write([]byte("<html><body>"))

	for k := range list {

		if !validFile(list[k]) {
			continue
		}

		filei := fmt.Sprintf("%s [%d] <a href=\"http://192.168.0.2:4444/f/%s\" target=\"_blank\">link</a> <br />", list[k].Name(), list[k].Size(), list[k].Name())
		w.Write([]byte(filei))

	}

	w.Write([]byte("</body></html>"))

}

func listFiles(dir string) (filelist []os.FileInfo, err error) {
	return ioutil.ReadDir(dir)
}

func validFile(file os.FileInfo) bool {

	if file.IsDir() {
		return false
	}

	if file.Size() == 0 {
		return false
	}

	sName := strings.Split(file.Name(), ".")

	if sName[len(sName)-1] != "mp4" {
		return false
	}

	return true

}
