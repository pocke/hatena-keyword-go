package keyword

import (
	"github.com/kolo/xmlrpc"
)

func rpcClient() (*xmlrpc.Client, error) {
	return xmlrpc.NewClient("http://d.hatena.ne.jp/xmlrpc", nil)
}
