package keyword

import (
	"strings"
	"testing"
)

func TestLink(t *testing.T) {
	body := "はてなキーワードをリンクして！"
	param := &LinkParam{
		Score:   20,
		CName:   []string{"book", "movie", "hatena"},
		ATarget: "_blank",
		AClass:  "keyword",
	}

	resp, err := Link(body, param)
	if err != nil {
		t.Fatal(err)
	}
	if resp == "" {
		t.Fatal("Response is empty.")
	}

	if !strings.Contains(resp, param.ATarget) {
		t.Fatalf("Response should contain %s, but got %s", param.ATarget, resp)
	}
	if !strings.Contains(resp, param.AClass) {
		t.Fatalf("Response should contain %s, but got %s", param.AClass, resp)
	}
}

func TestLinkLite(t *testing.T) {
	t.Skip()
	body := "はてなのブログははてなダイアリー"
	param := &LinkParam{}

	resp, err := LinkLite(body, param)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Wordlist) == 0 {
		t.Fatal("Response is empty")
	}
}
