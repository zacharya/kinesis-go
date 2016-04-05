package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {

	kc := kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	shard_itr_params := &kinesis.GetShardIteratorInput{
		ShardId:           aws.String("0"),            // Required
		ShardIteratorType: aws.String("TRIM_HORIZON"), // Required
		StreamName:        aws.String("zachstream"),   // Required
	}
	shard_itr_raw, err := kc.GetShardIterator(shard_itr_params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	shard_itr := shard_itr_raw.ShardIterator

	records_params := &kinesis.GetRecordsInput{
		ShardIterator: shard_itr, // Required
	}
	records_raw, err := kc.GetRecords(records_params)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < len(records_raw.Records); i += 1 {
		fmt.Println(string(records_raw.Records[i].Data[:]))
	}
}
