package healthz

import "net/http"
import "log"
import "encoding/json"

/*
Config for healthz handler. Used to supply hostnames, etc for services that are part of the health check
*/
type Config struct {
	Hostname string // Hostname of the server this app is running on
}

type handler struct {
	hostname string
	metadata map[string]string
}

/*
Handler for healthz service requests
*/
func Handler(hc *Config) (http.Handler, error) {
	metadata := make(map[string]string)

	h := &handler{hc.Hostname, metadata}
	return h, nil
}

/*
Response to healthz requests
*/
type Response struct {
	Hostname string            `json:"hostname"`
	Metadata map[string]string `json:"metadata"`
	Errors   []Error           `json:"error"`
}

/*
Error messages returned by healthz
*/
type Error struct {
	Description string            `json:"description"`
	Error       string            `json:"error"`
	Metadata    map[string]string `json:"metadata"`
	Type        string            `json:"type"`
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Hostname: h.hostname,
		Metadata: h.metadata,
	}

	statusCode := http.StatusOK

	errors := make([]Error, 0)

	/*
		   Check for signs of bad health here. Database connections, services depended on, etc.

		  errors = append(errors, Error{
			Type:        "Type",
			Description: "Description of error",
			Error: err.Error,
		  })
	*/

	response.Errors = errors
	if len(response.Errors) > 0 {
		statusCode = http.StatusInternalServerError
		for _, e := range response.Errors {
			log.Println(e.Error)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	data, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}
