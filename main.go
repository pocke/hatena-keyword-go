package keyword

import (
	"github.com/kolo/xmlrpc"
)

func rpcClient() (*xmlrpc.Client, error) {
	return xmlrpc.NewClient("http://d.hatena.ne.jp/xmlrpc", nil)
}

type similarWordParam struct {
	Wordlist []string `xmlrpc:"wordlist"`
}

type similarWordResp struct {
	Wordlist []struct {
		Word string `xmlrpc:"word"`
	} `xmlrpc:"wordlist"`
}

func (r *similarWordResp) toStringSlice() []string {
	resp := make([]string, 0, len(r.Wordlist))

	for _, w := range r.Wordlist {
		resp = append(resp, w.Word)
	}

	return resp
}

func SimilarWord(wordlist []string) ([]string, error) {
	c, err := rpcClient()
	if err != nil {
		return nil, err
	}

	param := &similarWordParam{
		Wordlist: wordlist,
	}

	resp := &similarWordResp{}
	err = c.Call("hatena.getSimilarWord", param, resp)
	if err != nil {
		return nil, err
	}

	return resp.toStringSlice(), nil
}
