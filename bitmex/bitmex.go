package bitmex

import (
	"encoding/json"
	"github.com/adi1382/Bitmex-Mirror-Bot/websocket"
	"go.uber.org/atomic"
	"go.uber.org/zap"
)

type Response interface {
	getTable() string
}

func (v Res) getTable() string {
	return v.Table
}

func (v OrderResponse) getTable() string {
	return v.Table
}

func (v PositionResponse) getTable() string {
	return v.Table
}

func (v MarginResponse) getTable() string {
	return v.Table
}

type Res struct {
	Success   bool        `json:"success,omitempty"`
	Subscribe string      `json:"subscribe,omitempty"`
	Request   interface{} `json:"request,omitempty"`
	Table     string      `json:"table,omitempty"`
	//Action    string      `json:"action,omitempty"`
	//Data      interface{} `json:"data,omitempty"`
}

type PositionResponse struct {
	Table  string                           `json:"table,omitempty"`
	Action string                           `json:"action,omitempty"`
	Keys   []string                         `json:"constants,omitempty"`
	Data   []websocket.PositionResponseData `json:"data,omitempty"`
}

type OrderResponse struct {
	Table  string                        `json:"table,omitempty"`
	Action string                        `json:"action,omitempty"`
	Keys   []string                      `json:"constants,omitempty"`
	Data   []websocket.OrderResponseData `json:"data,omitempty"`
}

type MarginResponse struct {
	Table  string                         `json:"table,omitempty"`
	Action string                         `json:"action,omitempty"`
	Keys   []string                       `json:"constants,omitempty"`
	Data   []websocket.MarginResponseData `json:"data,omitempty"`
}

func DecodeMessage(message []byte, logger *zap.Logger, restartRequired *atomic.Bool) (Response, string) {

	logger.Debug("Decoding Socket Message")

	var response Response
	var table string

	var res Res
	var positionResponse PositionResponse
	var orderResponse OrderResponse
	var marginResponse MarginResponse
	err := json.Unmarshal(message, &res)
	if err != nil {
		logger.Error("UnMarshal Error", zap.Error(err))
		restartRequired.Store(true)
		return response, table
	}

	table = res.Table

	if res.Table == "position" {
		err = json.Unmarshal(message, &positionResponse)
		if checkError(err, restartRequired, logger) {
			return response, ""
		}
		response = positionResponse
	} else if res.Table == "order" {
		err = json.Unmarshal(message, &orderResponse)
		if checkError(err, restartRequired, logger) {
			return response, ""
		}
		response = orderResponse
	} else if res.Table == "margin" {
		err = json.Unmarshal(message, &marginResponse)
		if checkError(err, restartRequired, logger) {
			return response, ""
		}
		response = marginResponse
	} else {
		err = json.Unmarshal(message, &res)
		if checkError(err, restartRequired, logger) {
			return response, ""
		}
		response = res
		table = ""
	}

	logger.Debug("Socket Message Decoded")

	return response, table
}

func checkError(err error, restartRequired *atomic.Bool, logger *zap.Logger) bool {
	if err != nil {
		logger.Error("New Error in Bitmex Package", zap.Error(err))
		restartRequired.Store(true)
		return true
	}
	return false
}
