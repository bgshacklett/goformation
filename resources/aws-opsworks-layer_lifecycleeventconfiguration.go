package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.LifecycleEventConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type AWSOpsWorksLayer_LifecycleEventConfiguration struct {

	// ShutdownEventConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
	ShutdownEventConfiguration AWSOpsWorksLayer_ShutdownEventConfiguration `json:"ShutdownEventConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.LifecycleEventConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
