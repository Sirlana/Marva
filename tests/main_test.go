package tests

import (
	"fmt"
	"testing"

	"sirlana.com/sirlana/sso/libs"
)

func TestMain(t *testing.T) {
	util := libs.NewUtil()
	// log := libs.NewLogger()

	// Config test
	// ConfigTest(util, log)

	// Parsing test
	ParsingTest(util)
}

func ConfigTest(util *libs.Util, log *libs.Logger) {
	cfg := libs.NewConfig("../config.json", util, log)
	apiPrivate := cfg.API()["private"].(map[string]interface{})["endpoints"].([]interface{})
	apiPublic := cfg.API()["public"].(map[string]interface{})["endpoints"].([]interface{})
	for _, val := range apiPrivate {
		fmt.Println(val)
	}
	for _, val := range apiPublic {
		fmt.Println(val)
	}
}

func ParsingTest(util *libs.Util) {
	data := util.ParseData("[string:email_username],[string:password],[string:product_id]")
	for _, val := range data {
		fmt.Println(util.GetDataKey(val))
	}
}
