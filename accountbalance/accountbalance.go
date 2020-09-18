package accountbalance

import (
	"github.com/paradox-3arthling/daraja_wrapper"
)

type Account_balance struct {
	Auth    *daraja_wrapper.Auth
	Payload *Acc_bal_payload
}

type Acc_bal_payload struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	PartyB                 string
	ReceiverIdentifierType string
	Remarks                string
	QueueTimeOutURL        string
	ResultURL              string
	AccountType            string
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

func (a *Account_balance) GetAcccountBalance() (*Daraja_response, error) {
	var resp *Daraja_response

	return resp, nil
}
