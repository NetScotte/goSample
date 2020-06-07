package iniconfig

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf MysqlConfig	`ini:"mysql"`
}

type ServerConfig struct {
	Ip string	`ini:"ip"`
	Port int	`ini:"port"`
}

type MysqlConfig struct {
	UserName string	`ini:"username"`
	Passwd string	`ini:"passwd"`
	Database string	`ini:"database"`
	Host string	`ini:"host"`
	Port int	`ini:"port"`
	Timeout float64 `ini:"timeout"`
}


func TestMarshal(t *testing.T) {
	conf := Config{
		ServerConfig{
			Ip: "192.168.1.1",
			Port: 1818,
		},
		MysqlConfig{
			UserName: "netliu",
			Passwd: "netliu",
			Database: "test",
			Host: "192.168.1.1",
			Port: 3030,
			Timeout: 2.2,
		},
	}
	data, err := Marshal(conf)
	if err != nil {
		t.Errorf("Marshal failed: %v", err)
	}
	t.Logf(string(data))
}

func TestMarshalFile(t *testing.T) {
	conf := Config{
		ServerConfig{
			Ip: "192.168.1.1",
			Port: 1818,
		},
		MysqlConfig{
			UserName: "netliu",
			Passwd: "netliu",
			Database: "test",
			Host: "192.168.1.1",
			Port: 3030,
			Timeout: 2.2,
		},
	}
	filename := "wconfig.ini"
	err := MarshalFile(conf, filename)
	if err != nil {
		t.Errorf("Failed MarshalFile: %v", err)
	}
}

func TestUnMarshal(t *testing.T) {
	var data []byte
	var conf Config
	data, err := ioutil.ReadFile("config.ini")
	if err != nil {
		t.Errorf("failed open config file: %s\n", err)
	}
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Logf("failed UnMarshal config content to Config type: %s", err)
	}
	fmt.Printf("get conf: \n%+v\n", conf)
}

func TestUnMarshalFile(t *testing.T) {
	filename := "config.ini"
	var conf Config
	err := UnMarshalFile(filename, &conf)
	if err != nil {
		t.Errorf("Failed UnMarshalFile: %v", err)
	}
	t.Logf("success UnMarshalFile")
	t.Logf("%+v", conf)
}