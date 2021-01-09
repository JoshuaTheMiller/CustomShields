package handler

import (
	"encoding/json"
	"net/http"
)

func codeDeployWebsiteSeemsToExist() bool {
	const urlToCheck = "https://codedeploy.experiments.joshuamiller.net/"

	_, err := http.Get(urlToCheck)

	if err != nil {
		return false
	}

	return true
}

// Handler is implemented per the Vercel Function spec:
// https://vercel.com/docs/serverless-functions/supported-languages#go
func Handler(w http.ResponseWriter, r *http.Request) {

	var label = "Application is currently"

	currentStateMessage := "DESTROYED"
	currentColor := "red"

	if codeDeployWebsiteSeemsToExist() {
		currentStateMessage = "AVAILABLE"
		currentColor = "green"
	}

	response := ShieldResponse{
		SchemaVersion: 1,
		Label:         label,
		Message:       currentStateMessage,
		Color:         currentColor,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
