package daraja_wrapper

import (
	"bytes"
	"fmt"
	"net/http"
)

var daraja_client http.Client

type Requester struct {
	Url     string
	Payload []byte
	Auth    *Auth
}

func (r *Requester) MakeRequest() (*http.Response, error) {
	r.Url = r.Auth.setUrl(r.Url)
	token, err := r.Auth.GetAuthKey()
	if err != nil {
		return nil, fmt.Errorf("`Auth.GetAuthKey()` got: %q", err)
	}
	req, err := http.NewRequest("POST", r.Url, bytes.NewBuffer(r.Payload))
	if err != nil {
		return nil, fmt.Errorf("`http.NewRequest/3` got: %q", err)
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token.Token)

	return daraja_client.Do(req)
	//	return nil, nil
}
