package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	enableMocks = false
	mocks       = make(map[string]*Mock)
)

//Mock struct
type Mock struct {
	URL        string
	HTTPMethod string
	Response   *http.Response
	Err        error
}

func getMockID(httpMethod string, url string) string {
	return fmt.Sprintf("%s_%s", httpMethod, url)
}

//StartMockups restclient
func StartMockups() {
	enableMocks = true
}

//StopMockups restclient
func StopMockups() {
	enableMocks = false
}

//AddMockup restclient
func AddMockup(mock Mock) {
	mocks[getMockID(mock.HTTPMethod, mock.URL)] = &mock
}

//FlushupMockups restclient
func FlushupMockups() {
	mocks = make(map[string]*Mock)
}

//Post method
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enableMocks {
		mock := mocks[getMockID(http.MethodPost, url)]
		if mock == nil {
			return nil, errors.New("no mockup found for given request")
		}

		return mock.Response, mock.Err
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}
	return client.Do(request)
}
