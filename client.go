package dnspod

import (
	"github.com/h2object/rpc"
	"github.com/h2object/h2object/util"
	"net/url"
	"time"
	// "log"
	"errors"
)

type Cache interface{
	Set(string, interface{}, time.Duration) 
	Add(string, interface{}, time.Duration) error
	Get(string) (interface{}, bool)
	Delete(string)
	DeleteExpired()
}

type Status struct{
	Code string `json:"code"`
	Message string `json:"message"`
	CreateAt string `json:"created_at"`
}

const DNSPODTOKEN = "DNSPODxTOKENx"

type DNSPodClient struct{
	email 		string 
	password 	string
	cache 		Cache
	addr 		string
	conn *rpc.Client
}


func NewDNSPodClient(addr string, cache Cache) *DNSPodClient {
	connection := rpc.NewClient(NewDNSPODAnalyzer())
	return &DNSPodClient{
		cache: cache,
		addr: addr,
		conn: connection,
	}
}

func (client *DNSPodClient) Authorize(login_email, login_password string) error {
	_, err := client.token(login_email, login_password)
	return err
}

func (client *DNSPodClient) token(login_email, login_password string) (string, error) {
	if client.cache != nil {
		if t, ok := client.cache.Get(DNSPODTOKEN); ok {
			return t.(string), nil
		}
	}

	data := url.Values{}
	data.Set("login_email", login_email)
	data.Set("login_password", login_password)
	data.Set("format", "json")

	var ret map[string]interface{}
	u := rpc.BuildHttpsURL(client.addr, "/Auth", nil)
	if err := client.conn.PostForm(nil, u, data, &ret); err != nil {
		return "",err
	}
	
	var status Status
	if err := util.Convert(ret["status"], &status); err != nil {
		return "",err
	}

	if status.Code != "1" {
		return "",errors.New(status.Message)
	}

	var token string
	if err := util.Convert(ret["user_token"], &token); err != nil {
		return "",err
	}
	if client.cache != nil {
		client.cache.Set(DNSPODTOKEN, token, time.Minute * 10)
	}
	client.email = login_email
	client.password = login_password
	return token, nil
}


// type DomainInfo struct{

// }

// func (client *DNSPodClient) GetDomainInfo(domain string, info *DomainInfo) error {

// }


// func (client *DNSPodClient) GetDomainInfo(domain string, info *DomainInfo) error {

// }

