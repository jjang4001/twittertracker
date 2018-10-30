package models

import (
	"time"
)

type Example struct {
	ExampleId         string `json:"exampleId" bson:"exampleId"`
	ExampleValue      string `json:"exampleValue" bson:"exampleValue"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
}

func NewExample(exampleId string, exampleValue string) *Example {
	timestamp := time.Now().Unix()
	ex := Example{ExampleId: exampleId, ExampleValue: exampleValue, TimestampCreated: timestamp}
	return &ex
}
