package daraja_wrapper

import (
	"errors"
	"net/http"
)

type Auth struct {
	Key, Secret string
	prod        bool
}

func (a *Auth) GetAuthKey() (map[string]string, error) {
	client := &http.Client{}
	url := a.setUrl()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(a.Key, a.Secret)
	req.Header.Add("Accept", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	// Do safaricom request here
	// convert response to a map from json
	return make(map[string]string), errors.New("Boiler plate code inserted, will fail")
}
func (a *Auth) setUrl() string {
	var url string
	if a.prod {
		url := "https://safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	} else {
		url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	}
	return url
}

// func makeClient() (*http.Client, error) {

// }
