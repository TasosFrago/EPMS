package router

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{
			ResponseWriter: w,
			status:         http.StatusOK,
		}
		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		methodColor := color.New(color.BgBlue).SprintFunc()
		statusColor := color.New(color.BgGreen).SprintFunc()
		if rec.status >= 400 && rec.status < 500 {
			statusColor = color.New(color.BgYellow).SprintFunc()
		} else if rec.status >= 500 {
			statusColor = color.New(color.BgRed).SprintFunc()
		}

		log.Printf(
			"%s %s\t%s %s  %s::%s  %s",
			methodColor(" "+r.Method+" "),
			r.URL.Path,
			r.RemoteAddr,
			r.UserAgent(),
			statusColor(" "+strconv.Itoa(rec.status)+" "),
			statusColor(" "+http.StatusText(rec.status)+" "),
			duration,
		)
		log.Printf("Headers: %+v", r.Header)
	})
}

type AvailableEndpoint struct {
	Path    string   `json:"path"`
	Methods []string `json:"methods"`
}

func LogAvailableEndpoints(router *mux.Router) {
	var endpoints []AvailableEndpoint
	bgCyan := color.New(color.BgCyan).SprintFunc()
	color.Cyan("Available Endpoints:")

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err != nil {
			return nil
		}
		methods, _ := route.GetMethods()
		if len(methods) == 0 {
			return nil
		}
		log.Printf("- %s [%s]", pathTemplate, bgCyan(" "+strings.Join(methods[:], ",")+" "))

		endpoints = append(endpoints, AvailableEndpoint{
			Path:    pathTemplate,
			Methods: methods,
		})

		return nil
	})

	// fmt.Println("\n\n[")
	// for _, endpoint := range endpoints {
	// 	jsonData, err := json.MarshalIndent(endpoint, "", "  ")
	// 	if err != nil {
	// 		fmt.Println("Error marshalling to JSON:", err)
	// 		continue
	// 	}
	// 	fmt.Println(string(jsonData) + ",")
	// }
	// fmt.Print("]\n\n")
}
