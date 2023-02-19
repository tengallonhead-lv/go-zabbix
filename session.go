package go_zabbix

import (
	"crypto/tls"
	"github.com/cavaliercoder/go-zabbix"
	"net/http"
)

type ZabbixServer struct {
	Addr     string `json:"address"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Session struct {
	*zabbix.Session
}

func (this *ZabbixServer) NewSession() (*Session, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	url := "http://" + this.Addr + "/zabbix/api_jsonrpc.php"
	session, err := zabbix.CreateClient(url).
		WithCache(cache).
		WithHTTPClient(client).
		WithCredentials(this.User, this.Password).
		Connect()
	if err != nil {
		return nil, err
	}
	cdSession := &Session{
		session,
	}
	return cdSession, nil
}
