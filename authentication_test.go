package daraja_wrapper

import "testing"

const prod_url = "https://safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
const sandbox_url = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

func TestProductionUrl(t *testing.T) {
	saf_auth := Auth{"key", "secret", true}
	got := saf_auth.setUrl()
	if got != prod_url {
		t.Errorf("Got %s, expected %s \n", got, prod_url)
	}
}
func TestSandboxUrl(t *testing.T) {
	saf_auth := Auth{"key", "secret", false}
	got := saf_auth.setUrl()
	if got != sandbox_url {
		t.Logf("Got %s, expected %s \n", got, prod_url)
	}
}
