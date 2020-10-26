package handlers

import (
	"github.com/NickBrisebois/HatchWaysAppBackend/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

// TestSetConfig tests function for setting handler package's configuration
func TestSetConfig(t *testing.T) {
	conf, err := config.LoadConfig(testConfig)
	if err != nil {
		t.Error("Error loading test config: " + err.Error())
		t.Fail()
	}

	SetConfig(conf)

	if !reflect.DeepEqual(conf, serverConfig) {
		t.Error("Set config not equal to expected config")
		t.Fail()
	}
}

// TestIsValidSortBy tests function for checking if a given sort method is valid or not
func TestIsValidSortBy(t *testing.T) {
	expectedTrue := isValidSortBy("id")
	expectedFalse := isValidSortBy("notarealsortmethod")

	if !expectedTrue {
		t.Error("Valid sort method was said to be not")
		t.Fail()
	}

	if expectedFalse {
		t.Error("Non valid sort method was said to be valid")
		t.Fail()
	}
}

func TestFullySetValidQueryContext(t *testing.T) {
	// Craft a minimal gin context with valid URL queries to test if we can extract the queries
	validFullySetQueryContext := &gin.Context{
		Request: &http.Request{
			URL: &url.URL{
				RawQuery: "tags=science,tech&sortBy=id&direction=asc",
			},
		},
	}

	expectedOutput := &queryConfig{
		tags: []string{"science", "tech"},
		sortBy: "id",
		direction: "asc",
	}

	if validQueries, err := getAndValidateQueries(validFullySetQueryContext); err != nil {
		t.Error("failed to get queries from valid configuration: " + err.Error())
		t.Fail()
	} else {
		if !reflect.DeepEqual(validQueries, expectedOutput) {
			t.Error("Returned values for fully set queries did not match expected values")
			t.Fail()
		}
	}
}

func TestPartiallySetValidQueryContext(t *testing.T) {
	validPartiallySetQueryContext := &gin.Context {
		Request: &http.Request {
			URL: &url.URL {
				RawQuery: "tags=science&direction=desc",
			},
		},
	}

	expectedOutput := &queryConfig{
		tags: []string{"science"},
		sortBy: "id",
		direction: "desc",
	}

	if validQueries, err := getAndValidateQueries(validPartiallySetQueryContext); err != nil {
		t.Error("failed to get queries from valid configuration: " + err.Error())
		t.Fail()
	} else {
		if !reflect.DeepEqual(validQueries, expectedOutput) {
			t.Error("Returned values for fully set queries did not match expected values")
			t.Fail()
		}
	}
}
func TestSetInvalidQueryContext(t *testing.T) {
	validPartiallySetQueryContext := &gin.Context {
		Request: &http.Request {
			URL: &url.URL {
				RawQuery: "tags=science&direction=notvalid",
			},
		},
	}

	if _, err := getAndValidateQueries(validPartiallySetQueryContext); err == nil {
		t.Error("function should have returned an error on invalid queries but did not")
		t.Fail()
	}
}
