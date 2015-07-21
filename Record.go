package dnspod

import (
	"errors"
	"net/url"
	"github.com/h2object/rpc"
	"github.com/h2object/h2object/util"
)

// {
//              "id": "47",
//              "name": "@",
//              "line": "Default",
//              "type": "NS",
//              "ttl": "600",
//              "value": "b.dnspod.com.",
//              "mx": "0",
//              "enabled": "1",
//              "status": "enabled",
//              "monitor_status": "",
//              "remark": "",
//              "updated_on": "2014-06-05 09:47:40",
//              "hold": "hold"
//          },
type RecordInfo struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Line string `json:"line"`
	Type string `json:"type"`
	TTL string `json:"ttl"`
	Value string `json:"value"`
	MX string `json:"mx"`
	Enabled string `json:"enabled"`
	Status string `json:"status"`
	UpdateAt string `json:"updated_on"`
	Hold string `json:"hold"`
}


func (client *DNSPodClient) AddRecord(domainID string, info *RecordInfo) error {
	t, err := client.token(client.email, client.password)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("user_token", t)
	data.Set("domain_id", domainID)
	data.Set("sub_domain", info.Name)
	data.Set("record_type", info.Type)
	data.Set("record_line", info.Line)
	data.Set("value", info.Value)
	if info.MX != "" {
		data.Set("mx", info.MX)
	}
	if info.TTL != "" {
		data.Set("ttl", info.TTL)
	}
	data.Set("format", "json")

	var ret map[string]interface{}
	u := rpc.BuildHttpsURL(client.addr, "/Record.Create", nil)
	if err := client.conn.PostForm(nil, u, data, &ret); err != nil {
		return err
	}
	
	var status Status
	if err := util.Convert(ret["status"], &status); err != nil {
		return err
	}

	if status.Code != "1" {
		return errors.New(status.Message)
	}

	if err := util.Convert(ret["record"], info); err != nil {
		return err
	}
	return nil
}

func (client *DNSPodClient) ModRecord(domainID string, info *RecordInfo) error {
	t, err := client.token(client.email, client.password)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("user_token", t)
	data.Set("domain_id", domainID)
	data.Set("sub_domain", info.Name)
	data.Set("record_type", info.Type)
	data.Set("record_line", info.Line)
	data.Set("value", info.Value)
	if info.MX != "" {
		data.Set("mx", info.MX)
	}
	if info.TTL != "" {
		data.Set("ttl", info.TTL)
	}
	data.Set("format", "json")

	var ret map[string]interface{}
	u := rpc.BuildHttpsURL(client.addr, "/Record.Modify", nil)
	if err := client.conn.PostForm(nil, u, data, &ret); err != nil {
		return err
	}
	
	var status Status
	if err := util.Convert(ret["status"], &status); err != nil {
		return err
	}

	if status.Code != "1" {
		return errors.New(status.Message)
	}

	if err := util.Convert(ret["record"], info); err != nil {
		return err
	}
	return nil
}

func (client *DNSPodClient) DelRecord(domainID string, recordID string) error {
	t, err := client.token(client.email, client.password)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("user_token", t)
	data.Set("domain_id", domainID)
	data.Set("record_id", recordID)
	data.Set("format", "json")

	var ret map[string]interface{}
	u := rpc.BuildHttpsURL(client.addr, "/Record.Remove", nil)
	if err := client.conn.PostForm(nil, u, data, &ret); err != nil {
		return err
	}
	
	var status Status
	if err := util.Convert(ret["status"], &status); err != nil {
		return err
	}

	if status.Code != "1" {
		return errors.New(status.Message)
	}
	return nil	
}

