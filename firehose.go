package main

import (
	"log"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
)

var sess *session.Session
var svc *firehose.Firehose

func init() {
	sess = session.Must(session.NewSession())
	svc = firehose.New(sess)
}

func BatchSend(record *GopherRecord) {
	data, _ := json.Marshal(record)

	params := &firehose.PutRecordInput{
		DeliveryStreamName: aws.String(cfg.FirehoseDeliveryStreamName), // Required
		Record: &firehose.Record{ // Required
			Data: data,
		},
	}
	resp, err := svc.PutRecord(params)

	if err != nil {
		log.Print(err.Error())
		return
	}

	// Pretty-print the response data.
	log.Print(resp)
}
