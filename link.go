package keyword

type LinkParam struct {
	Score   int      `xmlrpc:"score"`
	CName   []string `xmlrpc:"cname"`
	ATarget string   `xmlrpc:"a_target"`
	AClass  string   `xmlrpc:"a_class"`
}

func (p *LinkParam) linkParam(body string) *linkParam {
	r := &linkParam{}
	if p.Score != 0 {
		r.Score = &p.Score
	}
	if p.CName != nil {
		r.CName = p.CName
	}
	if p.ATarget != "" {
		r.ATarget = &p.ATarget
	}
	if p.AClass != "" {
		r.AClass = &p.ATarget
	}
	r.Body = body

	return r
}

type linkParam struct {
	Body    string   `xmlrpc:"body"`
	Score   *int     `xmlrpc:"score"`
	CName   []string `xmlrpc:"cname"`
	ATarget *string  `xmlrpc:"a_target"`
	AClass  *string  `xmlrpc:"a_class"`
	Mode    *string  `xmlrpc:"mode"`
}

type LinkResp string

type LinkLiteResp struct {
	Wordlist []struct {
		Word     string `xmlrpc:"word"`
		Score    int    `xmlrpc:"score"`
		Refcount int    `xmlrpc:"refcount"`
		CName    string `xmlrpc:"cname"`
	} `xmlrpc:"wordlist"`
}

func Link(body string, param *LinkParam) (string, error) {
	c, err := rpcClient()
	if err != nil {
		return "", err
	}

	p := param.linkParam(body)

	var resp LinkResp
	err = c.Call("hatena.setKeywordLink", p, &resp)
	if err != nil {
		return "", err
	}

	return string(resp), nil
}

// TODO: うまく動かない
func LinkLite(body string, param *LinkParam) (*LinkLiteResp, error) {
	c, err := rpcClient()
	if err != nil {
		return nil, err
	}
	p := param.linkParam(body)
	mode := "lite"
	p.Mode = &mode
	resp := &LinkLiteResp{}

	err = c.Call("hatena.setKeywordLink", p, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
