package myssh

import "testing"

func TestRun(t *testing.T) {
	sshClient, err := GetSSHClient("127.0.0.1", 22, "root", "jerkey")
	if err != nil {
		t.Fatal(err)
	}
	result, err := sshClient.sshRun("ls")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}
