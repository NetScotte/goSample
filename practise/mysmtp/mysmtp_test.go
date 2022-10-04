package mysmtp

import "testing"

func TestSendMail(t *testing.T) {
	err := SendMail("867233516@qq.com", "", "test email", "this is a test email", []map[string]string{})
	if err != nil {
		t.Error(err)
	}
}
