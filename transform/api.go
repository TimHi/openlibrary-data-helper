package transform

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/timhi/openlibrary-data-helper/m/v2/model"
)

func GetWorkFromApi(rating model.Rating) (model.Work, error) {
	workItem := model.Work{}
	baseURL := "https://openlibrary.org/"
	resultFormat := ".json"
	resp, err := http.Get(baseURL + rating.WorkKey + resultFormat)
	if err != nil {
		return model.Work{}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	json.Unmarshal([]byte(sb), &workItem)
	return workItem, nil
}
