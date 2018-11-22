package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// An AdBlockingStatus represents a current or desired state of add blocking.
type AdBlockingStatus struct {
	Disabled bool // whether ad blocking is disabled
}

// HandleAdBlockingStatus gets the current ad blocking status for a GET request or updates the ad
// blocking status for a PATCH request.
func HandleAdBlockingStatus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w)
		return
	case "PATCH":
		handlePatch(w, r)
		return
	}

	err := fmt.Sprintf("Method %s not supported.", r.Method)
	http.Error(w, err, 400)
}

func handleGet(w http.ResponseWriter) {
	adBlockingStatus := getAdBlockingStatus()
	enc := json.NewEncoder(w)
	enc.Encode(adBlockingStatus)
}

func handlePatch(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newAdBlockingStatus AdBlockingStatus
	err := decoder.Decode(&newAdBlockingStatus)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	updateAdBlockingStatus(newAdBlockingStatus)
	w.WriteHeader(http.StatusOK)
}

func getAdBlockingStatus() *AdBlockingStatus {
	output, _ := exec.Command("pihole", "status").Output()
	parsed := parseStatus(output)
	return parsed
}

// todo...
func updateAdBlockingStatus(newAdBlockingStatus AdBlockingStatus) *AdBlockingStatus {
	command := "enable"
	if newAdBlockingStatus.Disabled {
		command = "disable"
	}

	output, _ := exec.Command("pihole", command).Output()
	return parseStatus(output)
}

func parseStatus(commandOutput []byte) *AdBlockingStatus {
	status := string(commandOutput)
	trimmedStatus := strings.TrimSpace(status)
	disabled := strings.HasSuffix(trimmedStatus, "Disabled")
	return &AdBlockingStatus{Disabled: disabled}
}
