package dnspod

import (
	"encoding/json"
	"net/http"
)

type DNSPODAnalyzer struct {
}

func NewDNSPODAnalyzer() *DNSPODAnalyzer {
	return &DNSPODAnalyzer{}
}

/*
implement interface 
//! interface for the response analyser
type Analyzer interface {
	Analyse(ret interface{}, resp *http.Response) error
}
*/

func (analyser *DNSPODAnalyzer) Analyse(ret interface{}, resp *http.Response) (err error) {
	defer resp.Body.Close()

	if resp.StatusCode/100 == 2 {
		if ret != nil && resp.ContentLength != 0 {
			err = json.NewDecoder(resp.Body).Decode(ret)
			if err != nil {
				return
			}
		}
		if resp.StatusCode == 200 {
			return nil
		}
	}
	return nil
}