package logic

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"rusProfile/internal/app"
	"time"
)

const url = "https://www.rusprofile.ru/ajax.php"

type RusProfileLogic struct {
	httpClient *http.Client
}

func NewRusProfileLogic() app.RusProfileLogic {
	HTTPClient := http.Client{
		Transport: &http.Transport{MaxConnsPerHost: 10},
		Timeout:   6 * time.Second,
	}
	return &RusProfileLogic{httpClient: &HTTPClient}
}

func (r RusProfileLogic) GetCompanyByINN(inn string) (app.RusProfileData, error) {
	companyInfo := app.RusProfileData{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return companyInfo, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("content-type", "application/json")
	q := req.URL.Query()
	q.Add("query", inn)
	q.Add("action", "search")
	req.URL.RawQuery = q.Encode()
	response, err := r.httpClient.Do(req)
	if err != nil {
		return companyInfo, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return companyInfo, err
	}
	m := map[string]interface{}{}
	err = json.Unmarshal(data, &m)
	if v, ok := m["ul"]; ok {
		err = mapstructure.Decode(v.([]interface{})[0].(map[string]interface{}), &companyInfo)
		if err != nil {
			return companyInfo, err
		}
	}
	return companyInfo, nil
}
