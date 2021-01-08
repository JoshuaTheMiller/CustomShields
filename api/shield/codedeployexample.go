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

type ShieldResponse struct {
	SchemaVersion int    `json:"schemaVersion"`
	Label         string `json:"label"`
	Message       string `json:"message"`
	Color         string `json:"color"`
}

func Handler(w http.ResponseWriter, r *http.Request) {

	var label = "Application currently"

	currentStateMessage := "is DESTROYED"
	currentColor := "red"

	if codeDeployWebsiteSeemsToExist() {
		currentStateMessage = "exists"
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
