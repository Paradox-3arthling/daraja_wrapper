package accountbalance

import (
	"fmt"
	"testing"
)

const TEST_SECRET = "secret"
const TEST_KEY = "key"

func TestBasicFields(t *testing.T) {
	acc_bal_req := Init("", "", false)
	expected := fmt.Errorf("`secret` can not be blank!")
	_, err := acc_bal_req.GetAcccountBalance()
	if err != expected {
		t.Errorf("got: %q, expected: %q", err, expected)
	}
	acc_bal_req.Auth.Secret = TEST_SECRET

	expected = fmt.Errorf("`key` can not be blank!")
	_, err = acc_bal_req.GetAcccountBalance()
	if err.Error() != expected.Error() {
		t.Errorf("got: %q, expected: %q", err, expected)
	}
}
