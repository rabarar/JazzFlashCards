package main

import (
        "log"
        "net/http"
        "os"
        "path/filepath"
)

func main() {
        directory := "./static"
        fileServer := http.FileServer(http.Dir(directory))

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                path := filepath.Join(directory, filepath.Clean(r.URL.Path))

                // Serve existing files normally.
                if info, err := os.Stat(path); err == nil && !info.IsDir() {
                        fileServer.ServeHTTP(w, r)
                        return
                }

                // "/" serves index.html.
                if r.URL.Path == "/" {
                        http.ServeFile(w, r, filepath.Join(directory, "index.html"))
                        return
                }

                // Everything else gets the custom 404 page with a real HTTP 404 status.
                w.WriteHeader(http.StatusNotFound)
                http.ServeFile(w, r, filepath.Join(directory, "404.html"))
        })

        log.Printf("Starting HTTP server on :28001 serving %s\n", directory)

        if err := http.ListenAndServe("0.0.0.0:28001", nil); err != nil {
                log.Fatal("Server failed to start: ", err)
        }
}
