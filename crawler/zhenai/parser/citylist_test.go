package parser

import (
	"testing"
		"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝","City 阿克苏","City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Requests))
	}
	for i,url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d: %s;but was %s",i,url,result.Requests[i].Url)
		}
	}

	if len(result.Item) != resultSize {
		t.Errorf("result should have %d requests;but had %d",resultSize,len(result.Item))
	}
	for i,city := range expectedCities {
		if result.Item[i].(string)!= city {
			t.Errorf("excepted city #%d: %s;but was %s",i,city,result.Item[i].(string))
		}
	}

}