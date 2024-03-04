package main  

import (
  "net/http"
  "log"
)

func main() {
  const port = "8088"
  const rootPath = "."
  mux := http.NewServeMux()
  mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(rootPath))))

  mux.HandleFunc("/healthz", func (w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(http.StatusText(http.StatusOK)))
  })

  corsMux := middlewareCors(mux)
  
  server := &http.Server{
    Addr: ":" + port,
    Handler: corsMux,
  }
  log.Printf("Serving files from %s on port: %s\n", rootPath, port)
  log.Fatal(server.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
  })
}
