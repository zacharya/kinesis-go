package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/cloudfoundry/noaa/consumer"
)

const firehoseSubscriptionId = "firehose-a"

var (
	dopplerAddress = os.Getenv("DOPPLER_ADDR")
	authToken      = os.Getenv("CF_ACCESS_TOKEN")
)

func main() {

	kc := kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	consumer := consumer.New(dopplerAddress, &tls.Config{InsecureSkipVerify: true}, nil)
	consumer.SetDebugPrinter(ConsoleDebugPrinter{})

	fmt.Println("===== Streaming Firehose (will only succeed if you have admin credentials)")

	msgChan, errorChan := consumer.Firehose(firehoseSubscriptionId, authToken)

	go func() {
		for err := range errorChan {
			fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		}
	}()

	for envelope := range msgChan {

		fmt.Println(envelope.String())

		params := &kinesis.PutRecordInput{
			Data:         []byte(envelope.String()), // Required
			PartitionKey: aws.String("1"),           // Required
			StreamName:   aws.String("zachstream"),  // Required
		}
		_, err := kc.PutRecord(params)
		if err != nil {
			fmt.Println("error!!!: " + err.Error())
			return
		}

		fmt.Println("Added record!")
	}

}

type ConsoleDebugPrinter struct{}

func (c ConsoleDebugPrinter) Print(title, dump string) {
	println(title)
	println(dump)
}
