package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type AWSEMRCluster_ScalingAction struct {

	// Market AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-market
	Market string `json:"Market"`

	// SimpleScalingPolicyConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-simplescalingpolicyconfiguration
	SimpleScalingPolicyConfiguration AWSEMRCluster_SimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingAction) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
