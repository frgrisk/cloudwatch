package cloudwatch

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

func (c *mockClient) PutLogEvents(ctx context.Context, input *cloudwatchlogs.PutLogEventsInput, optFns ...func(options *cloudwatchlogs.Options)) (*cloudwatchlogs.PutLogEventsOutput, error) {
	args := c.Called(input)
	return args.Get(0).(*cloudwatchlogs.PutLogEventsOutput), args.Error(1)
}

func (c *mockClient) CreateLogStream(ctx context.Context, input *cloudwatchlogs.CreateLogStreamInput, optFns ...func(options *cloudwatchlogs.Options)) (*cloudwatchlogs.CreateLogStreamOutput, error) {
	args := c.Called(input)
	return args.Get(0).(*cloudwatchlogs.CreateLogStreamOutput), args.Error(1)
}

func (c *mockClient) GetLogEvents(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput, optFns ...func(options *cloudwatchlogs.Options)) (*cloudwatchlogs.GetLogEventsOutput, error) {
	args := c.Called(input)
	return args.Get(0).(*cloudwatchlogs.GetLogEventsOutput), args.Error(1)
}
