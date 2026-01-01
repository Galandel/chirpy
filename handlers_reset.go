package main

import (
	"fmt"
	"log"
	"net/http"
)

// Reset hits metric
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {

	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Reset is only allowed in dev environment.", nil)
		return
	}
	cfg.fileserverHits.Store(0)     // reset counter back to 0
	cfg.db.DeleteUsers(r.Context()) // Delete all users

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := fmt.Fprint(w, "Hits reset to 0 and database reset to initial state."); err != nil {
		log.Printf("error writing response: %v", err)
	}
}
