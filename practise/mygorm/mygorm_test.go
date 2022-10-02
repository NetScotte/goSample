package mygorm

import (
	"testing"
)

func Test_Migrate(t *testing.T) {
	if err := Migrate(); err != nil {
		t.Error(err)
	}
}

func Test_Create(t *testing.T) {
	if err := Create(); err != nil {
		t.Fatal(err)
	}
}

func Test_GetById(t *testing.T) {
	tests := []struct {
		name  string
		input uint
		want  string
	}{
		{
			name:  "基础测试",
			input: 1,
			want:  "liufy47",
		},
		{
			name:  "基础测试2",
			input: 2,
			want:  "zhangly97",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user, err := GetById(test.input)
			if err != nil {
				t.Error(err)
			}
			if user.Username != test.want {
				t.Errorf("input id: %v, want username: %v, get: %v\n", test.input, test.want, user.Username)
			}
		})
	}
}

func Test_GetList(t *testing.T) {
	userList, err := GetList(1, 10, "")
	if err != nil {
		t.Error(err)
	}
	if len(userList) == 0 {
		t.Errorf("结果为空")
	}

}

func TestDisableUser(t *testing.T) {
	err := DisableUser(3)
	if err != nil {
		t.Error(err)
	}
}
