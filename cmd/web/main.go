package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	session       *sessions.Session
	templateCache map[string]*template.Template
}

type contextKey string

var contextKeyUser = contextKey("user")

func main() {
	// Config & CLI Flags
	addr := flag.String("addr", ":4091", "HTTP network address")
	dsn := flag.String("dsn", "web:pass@/prettyBookmarks?parseTime=true", "MySQL database DSN string")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret session key")
	flag.Parse()

	// Logging
	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// DB Connection
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	// Bookmark Service for DB actions
	// bookmarkModel, err := mysql.NewBookmarkModel(db)
	// if err != nil {
	// 	errorLog.Fatal(err)
	// }

	// Initialise a new template cache
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	// Session Manager
	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour
	session.Secure = true
	session.SameSite = http.SameSiteStrictMode

	// Dependency injection for handlers
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		session:       session,
		templateCache: templateCache,
	}

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	// Start server
	srv := &http.Server{
		Addr:         "localhost" + *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// ListenAndServe is blocking, so an error will only be logged if and when the function returns
	infoLog.Printf("Starting server on %s\n", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	errorLog.Fatal(err)
}

func openDB(dns string) (*sql.DB, error) {
	// Establish pool to hold connections to database
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	// Test that a connection can be made by pinging the database
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
