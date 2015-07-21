package dnspod

import (
	"errors"
	"net/url"
	"github.com/h2object/rpc"
	"github.com/h2object/h2object/util"
)
// "domain": {
//         "id": "6",
//         "name": "dnspod.com",
//         "punycode": "dnspod.com",
//         "grade": "DP_Free",
//         "grade_title": "Free",
//         "status": "enable",
//         "ext_status": "notexist",
//         "records": "3",
//         "group_id": "1",
//         "is_mark": "no",
//         "remark": false,
//         "is_vip": "no",
//         "searchengine_push": "yes",
//         "beian": "no",
//         "user_id": "730060",
//         "created_on": "2014-06-04 16:19:31",
//         "updated_on": "2014-06-04 16:20:05",
//         "ttl": "600",
//         "owner": "yizero@qq.com"
//     }

type DomainInfo struct{
	ID string `json:"id"`
	Name string `json:"name"`
	PunyCode string `json:"punycode"`
	Grade string `json:"grade"`
	GradeTitle string `json:"grade_title"`
	Status string `json:"status"`
	ExtStatus string `json:"ext_status"`
	Records string `json:"records"`
	GroupID string `json:"group_id"`
	IsMark string `json:"is_mark"`
	Remark string `json:"remark"`
	IsVIP string `json:"is_vip"`
	UserID string `json:"user_id"`
	Owner string `json:"owner"`
	TTL string `json:"ttl"`
}

func (client *DNSPodClient) GetDomainInfo(domain string, info *DomainInfo) error {
	t, err := client.token(client.email, client.password)
	if err != nil {
		return err
	}

	data := url.Values{}
	data.Set("user_token", t)
	data.Set("domain", domain)
	data.Set("format", "json")

	var ret map[string]interface{}
	u := rpc.BuildHttpsURL(client.addr, "/Domain.Info", nil)
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

	if err := util.Convert(ret["domain"], info); err != nil {
		return err
	}
	return nil
}