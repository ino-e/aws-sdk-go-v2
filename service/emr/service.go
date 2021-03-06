// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package emr

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// EMR provides the API operation methods for making requests to
// Amazon Elastic MapReduce. See this package's package overview docs
// for details on the service.
//
// EMR methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type EMR struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*EMR)

// Used for custom request initialization logic
var initRequest func(*EMR, *aws.Request)

// Service information constants
const (
	ServiceName = "elasticmapreduce" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName        // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the EMR client with a config.
//
// Example:
//     // Create a EMR client from just a config.
//     svc := emr.New(myConfig)
func New(config aws.Config) *EMR {
	var signingName string
	signingRegion := config.Region

	svc := &EMR{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2009-03-31",
				JSONVersion:   "1.1",
				TargetPrefix:  "ElasticMapReduce",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a EMR operation and runs any
// custom request initialization.
func (c *EMR) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
