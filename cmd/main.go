package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	templateIndex "github.com/jkandasa/static-file-server/pkg/template"
	"github.com/jkandasa/static-file-server/pkg/utils"
)

const (
	ENV_BRAND_NAME     = "BRAND_NAME"
	DEFAULT_BRAND_NAME = "Lightweight Static File Server"
)

var (
	baseDir         string
	brandName       string
	webpageTemplate *template.Template
)

func main() {
	port := flag.String("port", "8080", "port to serve on")
	directory := flag.String("dir", "/data", "the static directory to host")
	flag.Parse()

	baseDir = *directory
	brandName = os.Getenv(ENV_BRAND_NAME)
	if brandName == "" {
		brandName = DEFAULT_BRAND_NAME
	}

	// compile webpage template
	compiledTemplate, err := template.New("webpage").Parse(templateIndex.INDEX)
	if err != nil {
		panic(err)
	}
	webpageTemplate = compiledTemplate

	http.HandleFunc("/", handleRequests)
	fmt.Printf("serving %s on HTTP port: %s\n", *directory, *port)
	address := fmt.Sprintf(":%s", *port)

	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path

	fileFullPath, valid := utils.IsValidPath(baseDir, filename)
	if !valid {
		writeFileNotFound(w)
		return
	}

	file := utils.IsFileExists(fileFullPath)
	if file == nil {
		writeFileNotFound(w)
		return
	}

	if file.IsDir {
		// check the index.html is available
		indexFile := filepath.Join(fileFullPath, "index.html")
		if utils.IsFileExists(indexFile) != nil {
			http.ServeFile(w, r, indexFile)
			return
		}

		files, err := utils.ListFiles(baseDir, fileFullPath)
		if err != nil {
			fmt.Println(err)
			writeFileNotFound(w)
			return
		}

		baseReference := strings.TrimPrefix(path.Dir(fileFullPath), baseDir)
		var tplBuffer bytes.Buffer
		err = webpageTemplate.Execute(&tplBuffer, map[string]interface{}{"Files": files, "Dir": baseReference, "Brand": brandName})

		if err != nil {
			_, err = w.Write(tplBuffer.Bytes())
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		_, err = w.Write(tplBuffer.Bytes())
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	http.ServeFile(w, r, fileFullPath)
}

func writeFileNotFound(w http.ResponseWriter) {
	http.Error(w, "404 page not found", http.StatusNotFound)
}
