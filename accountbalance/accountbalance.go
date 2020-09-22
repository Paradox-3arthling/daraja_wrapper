package accountbalance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/paradox-3arthling/daraja_wrapper"
)

type Account_balance struct {
	Auth    *daraja_wrapper.Auth
	Payload *Acc_bal_payload
}

type Acc_bal_payload struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	PartyA             string
	IdentifierType     string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
}

func Init(key, secret string, prod bool) *Account_balance {
	auth := &daraja_wrapper.Auth{
		Key:    key,
		Secret: secret,
		Prod:   prod,
	}
	payload := &Acc_bal_payload{}
	acc_bal := &Account_balance{
		Auth:    auth,
		Payload: payload,
	}
	return acc_bal
}

type Daraja_response map[string]interface{}

const ACC_URL = "https://sandbox.safaricom.co.ke/mpesa/accountbalance/v1/query"

func (a *Account_balance) GetAcccountBalance() (*Daraja_response, error) {
	var resp *Daraja_response
	a.Payload.CommandID = "AccountBalance"
	a.Payload.IdentifierType = "4"
	req_payload, err := json.Marshal(a.Payload)
	if err != nil {
		return resp, fmt.Errorf("`json.Marshal/1` got error: %q", err)
	}

	requester := &daraja_wrapper.Requester{
		Url:     ACC_URL,
		Payload: req_payload,
		Auth:    a.Auth,
	}
	resp_daraja, err := requester.MakeRequest()
	defer resp_daraja.Body.Close()
	body, err := ioutil.ReadAll(resp_daraja.Body)
	if err != nil {
		return resp, fmt.Errorf("`ioutil.ReadAll/1` got error: %q", err)
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, fmt.Errorf("`json.Unmarshal/2` got error: %q", err)
	}
	return resp, nil
}
