package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type headerTransport struct {
	T   http.RoundTripper
	jwt string
}

func (adt *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", "golang-sesam-client")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", adt.jwt))
	req.Header.Add("Accept", "application/json,*/*")
	return adt.T.RoundTrip(req)
}

//API Sesam API struct
type API struct {
	node       string
	httpClient *http.Client
}

func (a *API) doGet(url string) ([]byte, error) {
	resp, err := a.httpClient.Get(fmt.Sprintf("%senv", a.node))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

//NewAPI factory function to obtain Api instance
func NewAPI(node string, jwt string) API {
	roundTrip := &headerTransport{
		T:   http.DefaultTransport,
		jwt: jwt,
	}
	return API{
		node: fmt.Sprintf("https://%s/api/", node),
		httpClient: &http.Client{
			Transport: roundTrip,
		},
	}
}

//GetEnvironmentVariables returns environmental variables for given node
func (a *API) GetEnvironmentVariables() (map[string]interface{}, error) {
	respBytes, err := a.doGet(fmt.Sprintf("%senv", a.node))
	if err != nil {
		return nil, err
	}

	envMap := make(map[string]interface{})

	err = json.Unmarshal(respBytes, &envMap)
	if err != nil {
		return nil, err
	}
	return envMap, nil
}
