package libs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Util struct{}

func NewUtil() *Util {
	return &Util{}
}

func (u Util) JSONData(file string) (map[string]interface{}, error) {
	jsonFile, err := os.Open(file)
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	var arr map[string]interface{}
	err = json.Unmarshal([]byte(jsonData), &arr)
	return arr, err
}

func (u Util) ShowErrorResponse(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(data["code"].(int))
	errors := make(map[string]string)
	errors["message"] = data["message"].(string)
	json.NewEncoder(w).Encode(errors)
}

func (u Util) ShowCustomErrorResponse(w http.ResponseWriter, httpCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	errors := make(map[string]string)
	errors["message"] = message
	json.NewEncoder(w).Encode(errors)
}

func (u Util) GetFormData(r *http.Request, textParam string) map[string]interface{} {
	data := make(map[string]interface{})
	params := u.ParseData(textParam)
	for _, val := range params {
		data[u.GetDataKey(val)] = r.FormValue(u.GetDataKey(val))
	}
	return data
}

func (u Util) ParseParam(r *http.Request, paramType string, dataParam string) map[string]interface{} {
	var data map[string]interface{}
	if strings.EqualFold(paramType, "form-data") {
		data = u.GetFormData(r, dataParam)
	}
	return data
}

func (u Util) ParseData(text string) map[int]string {
	regex := regexp.MustCompile(`\[([^\[\]]*)\]`)
	results := regex.FindAllString(text, -1)
	data := make(map[int]string)
	for i, val := range results {
		val = strings.Trim(val, "[")
		val = strings.Trim(val, "]")
		data[i] = val
	}
	return data
}

func (u Util) GetDataKey(text string) string {
	data := strings.Split(text, ":")
	return data[1]
}
