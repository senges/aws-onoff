package main

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

/* Power instance(s) on */
func powerOn(instance string) error {

	if exist := checkMachine(instance); exist != true {
		return errors.New("EC2 '" + instance + "' does not exists in config file")
	}

	config := aws.Config{
		Region:      aws.String(ConfigFile.Instance[instance].Region),
		Credentials: credentials.NewStaticCredentials(
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsKeyId,
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsSecretKey,
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsToken,
		),
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            config,
	}))

	svc := ec2.New(sess)

	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{
			aws.String(ConfigFile.Instance[instance].InstanceId),
		},
		DryRun: aws.Bool(true),
	}
	result, err := svc.StartInstances(input)
	awsErr, ok := err.(awserr.Error)

	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StartInstances(input)
		
		if err != nil {
			fmt.Println("Error", err)
			fmt.Println("Couldn't start target")
		} else {
			fmt.Println("Success", result.StartingInstances)
			fmt.Println("Instance(s) properly started")
		}
	} else { 
		fmt.Println("Error", err)
		fmt.Println("Instance does not exists, or you have insufficient permission")
	}

	return nil
}

/* Power instance(s) off */
func powerOff(instance string) error {

	if exist := checkMachine(instance); exist != true {
		return errors.New("EC2 '" + instance + "' does not exists in config file")
	}

	config := aws.Config{
		Region:      aws.String(ConfigFile.Instance[instance].Region),
		Credentials: credentials.NewStaticCredentials(
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsKeyId,
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsSecretKey,
			ConfigFile.AwsAccounts[ConfigFile.Instance[instance].AwsAccount].AwsToken,
		),
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            config,
	}))

	svc := ec2.New(sess)

	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{
			aws.String(ConfigFile.Instance[instance].InstanceId),
		},
		DryRun: aws.Bool(true),
	}

	result, err := svc.StopInstances(input)
	awsErr, ok := err.(awserr.Error)
	
	if ok && awsErr.Code() == "DryRunOperation" {
		input.DryRun = aws.Bool(false)
		result, err = svc.StopInstances(input)
		
		if err != nil {
			fmt.Println("Error", err)
			fmt.Println("Couldn't stop target")
		} else {
			fmt.Println("Success", result.StoppingInstances)
			fmt.Println("Instance(s) properly stopped")
		}

	} else {
		fmt.Println("Error", err)
		fmt.Println("Instance does not exists, or you have insufficient permission")
	}

	return nil
}

/* Check if machine is in config file */
func checkMachine(name string) bool {
	if _, exists := ConfigFile.Instance[name]; exists {
		return true
	}

	return false
}
