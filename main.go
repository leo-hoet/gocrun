package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

const version = "1.0.0"

var logger *slog.Logger

func homeHelloWord(w http.ResponseWriter, r *http.Request, l *slog.Logger) {
	clientIP := r.RemoteAddr

	if l != nil {
		l.Info("Serving request in /", "clientIp", clientIP)
	}

	fmt.Fprintf(w, `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Hello World</title>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    text-align: center;
                    padding: 50px;
                    background-color: #f4f4f4;
                }
                h1 {
                    color: #333;
                    font-size: 36px;
                    margin-bottom: 20px;
                }
                p {
                    color: #666;
                    font-size: 18px;
                }
            </style>
        </head>
        <body>
            <h1>Hello World!</h1>
            <p>version: %s</p>
        </body>
        </html>
        `, version)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { homeHelloWord(w, r, logger) })

	fmt.Println("Starting server on port 4000")
	http.ListenAndServe(":"+port, nil)
}
