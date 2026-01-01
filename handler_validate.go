package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// validate chirp
func handlerValidate(w http.ResponseWriter, r *http.Request) {
	const maxChirpLength = 140
	type parameters struct {
		Body string `json:"body"`
	}

	type validReturn struct {
		Cleaned_Body string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleaned_body := profaneCleaner(params.Body)

	respondWithJSON(w, http.StatusOK, validReturn{
		Cleaned_Body: cleaned_body,
	})
}

func profaneCleaner(chirp string) string {
	censor_list := map[string]bool{
		"kerfuffle": true,
		"sharbert":  true,
		"fornax":    true,
	}
	words := strings.Split(chirp, " ")
	for idx, word := range words {
		if censor_list[strings.ToLower(word)] {
			words[idx] = "****"
		}
	}
	return strings.Join(words, " ")

}
