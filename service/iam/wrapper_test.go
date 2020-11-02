package iam

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	region = "region"
	host   = "hosts"
	testAk = "ak"
	testSk = "sk"
)

func TestIAM_ListUsers(t *testing.T) {
	DefaultInstance.Client.SetAccessKey(testAk)
	DefaultInstance.Client.SetSecretKey(testSk)
	DefaultInstance.SetHost(host)
	DefaultInstance.SetRegion(region)

	list, status, err := DefaultInstance.ListUsers(nil, nil)
	fmt.Println(status, err)
	b, _ := json.Marshal(list)
	fmt.Println(string(b))
}
