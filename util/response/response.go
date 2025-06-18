package response

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	jsontime "github.com/liamylian/jsontime/v2/v2"
)

var jsoniter = jsontime.ConfigWithCustomTimeFormat

type responseSuccess struct {
	Result interface{} `json:"result"`
}

func parse(result interface{}) json.RawMessage {
	var response json.RawMessage
	bytes, _ := jsoniter.Marshal(result)
	_ = json.Unmarshal(bytes, &response)

	return response
}

func HandleSuccess(c *gin.Context, statusCode int, result interface{}) {
	var response json.RawMessage
	bytes, _ := jsoniter.Marshal(result)
	_ = json.Unmarshal(bytes, &response)

	c.JSON(statusCode, responseSuccess{
		Result: parse(response),
	})
}
