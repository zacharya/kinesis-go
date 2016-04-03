package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

func main() {

	kc := kinesis.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

	descParams := &kinesis.DescribeStreamInput{
		StreamName: aws.String("zachstream"), // Required
		Limit:      aws.Int64(1),
	}
	descRaw, err := kc.DescribeStream(descParams)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(descRaw)

	var records []*kinesis.PutRecordsRequestEntry
	for i := 0; i < 100; i++ {

		putrecordsEntry := &kinesis.PutRecordsRequestEntry{
			PartitionKey: aws.String(string(i)),
			Data:         []byte("zachpayload1" + string(i)),
		}
		records = append(records, putrecordsEntry)
	}

	putrecordsParams := &kinesis.PutRecordsInput{
		Records:    records,                  // Required
		StreamName: aws.String("zachstream"), // Required
	}
	putrecordsRaw, err := kc.PutRecords(putrecordsParams)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(putrecordsRaw)

	// shard_itr_params := &kinesis.GetShardIteratorInput{
	// 	ShardId:           aws.String("0"),          // Required
	// 	ShardIteratorType: aws.String("LATEST"),     // Required
	// 	StreamName:        aws.String("zachstream"), // Required
	// }
	// shard_itr_raw, err := kc.GetShardIterator(shard_itr_params)
	//
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	//
	// fmt.Println(shard_itr_raw)
	// shard_itr := shard_itr_raw.ShardIterator
	//
	// records_params := &kinesis.GetRecordsInput{
	// 	ShardIterator: shard_itr, // Required
	// 	Limit:         aws.Int64(1),
	// }
	// records_raw, err := kc.GetRecords(records_params)
	//
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(records_raw.Records)
	//
	// for i := 0; i < len(records_raw.Records); i += 1 {
	// 	fmt.Println(string(records_raw.Records[i].Data[:]))
	// }
}
