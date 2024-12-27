package httpError

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type LogType int

const (
	INFO = iota + 1
	SUCCESS
	REDIRECTION
	WARNING
	ERROR
)

type ErrorRes struct {
	Code        int    `json:"code"`
	Message     string `json:"msg"`
	Info        string `json:"info,omitempty"`
	Success     string `json:"success,omitempty"`
	Redirection string `json:"redirection,omitempty"`
	Warning     string `json:"warning,omitempty"`
	Error       string `json:"error,omitempty"`
}

func InternalServerError(w http.ResponseWriter, errMsg string) {
	httpLogger(ERROR, errMsg)
	httpResponse(w, http.StatusInternalServerError, "Internal Server Error", errMsg)
}

func NotFoundError(w http.ResponseWriter, errMsg string) {
	httpLogger(WARNING, errMsg)
	httpResponse(w, http.StatusNotFound, "Not Found", errMsg)
}

func UnauthorizedError(w http.ResponseWriter, errMsg string) {
	httpLogger(WARNING, errMsg)
	httpResponse(w, http.StatusUnauthorized, "Unauthorized", errMsg)
}

func ConflictError(w http.ResponseWriter, errMsg string) {
	httpLogger(WARNING, errMsg)
	httpResponse(w, http.StatusConflict, "Conflict", errMsg)
}

func BadRequestError(w http.ResponseWriter, errMsg string) {
	httpLogger(WARNING, errMsg)
	httpResponse(w, http.StatusBadRequest, "Bad Request", errMsg)
}

func UnprocessableEntityError(w http.ResponseWriter, errMsg string) {
	httpLogger(WARNING, errMsg)
	httpResponse(w, http.StatusBadRequest, "", errMsg)
}

func StatusCreated(w http.ResponseWriter, errMsg string) {
	httpLogger(SUCCESS, errMsg)
	httpResponse(w, http.StatusUnprocessableEntity, "Unprocessable Entity", errMsg)
}

func httpLogger(logType LogType, errMsg string) {
	// TODO Need to disable for prod
	bgRed := color.New(color.BgRed).SprintFunc()
	bgBlue := color.New(color.BgBlue).SprintFunc()
	bgCyan := color.New(color.BgCyan).SprintFunc()
	bgYellow := color.New(color.BgYellow).SprintFunc()
	bgGreen := color.New(color.BgGreen).SprintFunc()
	var logHeader string
	switch logType {
	case INFO:
		logHeader = bgBlue("HTTP-INFO:")
	case SUCCESS:
		logHeader = bgGreen("HTTP-SUCCESS:")
	case REDIRECTION:
		logHeader = bgCyan("HTTP-REDIR:")
	case WARNING:
		logHeader = bgYellow("HTTP-WARNING:")
	case ERROR:
		logHeader = bgRed("HTTP-ERROR:")
	}
	log.Printf(
		"%s %s\n",
		logHeader,
		errMsg,
	)
}

func httpResponse(w http.ResponseWriter, statusCode int, message string, detailedMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := ErrorRes{
		Code:    statusCode,
		Message: message,
	}
	var logType LogType
	if statusCode >= 100 && statusCode <= 199 {
		logType = INFO
	} else if statusCode >= 200 && statusCode <= 299 {
		logType = SUCCESS
	} else if statusCode >= 300 && statusCode <= 399 {
		logType = REDIRECTION
	} else if statusCode >= 400 && statusCode <= 499 {
		logType = WARNING
	} else if statusCode >= 500 && statusCode <= 599 {
		logType = ERROR
	}

	if os.Getenv("DEBUG") != "" {
		switch logType {
		case INFO:
			res.Info = detailedMsg
		case SUCCESS:
			res.Success = detailedMsg
		case WARNING:
			res.Warning = detailedMsg
		case ERROR:
			res.Error = detailedMsg
		}
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		log.Println(fmt.Errorf("Error: utls/httpError:httpResponse: Failed marshaling JSON error %w", err))
		// If JSON marshalling fails, fallback to plain text error
		http.Error(w, "500: Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonResponse)
}
