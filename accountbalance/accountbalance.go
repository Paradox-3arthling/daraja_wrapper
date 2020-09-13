package accountbalance

import (
	auth "github.com/paradox-3arthling/daraja_wrapper"
)

type AccountBalance struct {
	Auth                   *auth.Auth
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
