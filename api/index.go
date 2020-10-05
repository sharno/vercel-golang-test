package handler

import (
	"fmt"
	"net/http"
	"time"

	firebase "firebase.google.com/go"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, currentTime)

	// app, err := firebase.NewApp(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("error initializing app: %v\n", err)
	// }
	fmt.Fprintf(w, firebase.Version)
}
