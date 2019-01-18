package test

import (
	"fmt"
	"github.com/Earth-Online/zoomeye-api"
	"github.com/Earth-Online/zoomeye-api/facets"
	"testing"
)

const username = ""
const password = ""
const fake = "test"

func TestLogin(t *testing.T) {
	_, err := zoomeye.NewUser(
		fake,
		fake,
	)
	if err == nil {
		t.Error("登陆失败未报错")
		return
	}

	_, err = zoomeye.NewUser(
		username,
		password,
	)
	if err != nil {
		t.Error("登陆失败")
		return
	}
}

func TestResourcesInfo(t *testing.T) {
	user, err := zoomeye.NewUser(
		username,
		password,
	)
	if err != nil {
		t.Error("登陆失败")
		return
	}
	_, err = user.ResourcesInfo()

	if err != nil {
		t.Error(err)
		return
	}
}

func TestHostSearch(t *testing.T) {
	user, err := zoomeye.NewUser(
		username,
		password,
	)
	if err != nil {
		t.Error("登陆失败")
		return
	}
	info, err := user.HostSearch("title:oneindex", 1, []string{facets.Port})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Print(info)
}

func TestWebSearch(t *testing.T) {
	user, err := zoomeye.NewUser(
		username,
		password,
	)
	if err != nil {
		t.Error("登陆失败")
		return
	}
	info, err := user.WebSearch("title:oneindex", 1, []string{})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Print(info)
}
