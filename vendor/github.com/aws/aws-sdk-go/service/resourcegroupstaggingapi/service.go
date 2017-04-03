// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// This guide describes the API operations for the resource groups tagging.
//
// A tag is a label that you assign to an AWS resource. A tag consists of a
// key and a value, both of which you define. For example, if you have two Amazon
// EC2 instances, you might assign both a tag key of "Stack." But the value
// of "Stack" might be "Testing" for one and "Production" for the other.
//
// Tagging can help you organize your resources and enables you to simplify
// resource management, access management and cost allocation. For more information
// about tagging, see Working with Tag Editor (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/tag-editor.html)
// and Working with Resource Groups (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/resource-groups.html).
// For more information about permissions you need to use the resource groups
// tagging APIs, see Obtaining Permissions for Resource Groups  (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-resource-groups.html)
// and Obtaining Permissions for Tagging  (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html).
//
// You can use the resource groups tagging APIs to complete the following tasks:
//
//    * Tag and untag supported resources located in the specified region for
//    the AWS account
//
//    * Use tag-based filters to search for resources located in the specified
//    region for the AWS account
//
//    * List all existing tag keys in the specified region for the AWS account
//
//    * List all existing values for the specified key in the specified region
//    for the AWS account
//
// Not all resources can have tags. For a list of resources that support tagging,
// see Supported Resources (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/supported-resources.html)
// in the AWS Resource Groups and Tag Editor User Guide.
//
// To make full use of the resource groups tagging APIs, you might need additional
// IAM permissions, including permission to access the resources of individual
// services as well as permission to view and apply tags to those resources.
// For more information, see Obtaining Permissions for Tagging (http://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/obtaining-permissions-for-tagging.html)
// in the AWS Resource Groups and Tag Editor User Guide.
// The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/resourcegroupstaggingapi-2017-01-26
type ResourceGroupsTaggingAPI struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "tagging"   // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ResourceGroupsTaggingAPI client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ResourceGroupsTaggingAPI client from just a session.
//     svc := resourcegroupstaggingapi.New(mySession)
//
//     // Create a ResourceGroupsTaggingAPI client with additional configuration
//     svc := resourcegroupstaggingapi.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ResourceGroupsTaggingAPI {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion, signingName string) *ResourceGroupsTaggingAPI {
	svc := &ResourceGroupsTaggingAPI{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2017-01-26",
				JSONVersion:   "1.1",
				TargetPrefix:  "ResourceGroupsTaggingAPI_20170126",
			},
			handlers,
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
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ResourceGroupsTaggingAPI operation and runs any
// custom request initialization.
func (c *ResourceGroupsTaggingAPI) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}