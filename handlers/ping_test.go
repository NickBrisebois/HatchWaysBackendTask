package handlers

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"testing"
)

func TestGetPingResponse(t *testing.T) {
	expectedResponse := gin.H{"success": true}
	actualResponse := getPingResponse()

	if !reflect.DeepEqual(expectedResponse, actualResponse) {
		t.Error("Expected ping response not the same as actual ping response")
		t.Fail()
	}
}