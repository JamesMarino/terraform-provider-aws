// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package sfn

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	sfn_sdkv1 "github.com/aws/aws-sdk-go/service/sfn"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceActivity,
			TypeName: "aws_sfn_activity",
		},
		{
			Factory:  DataSourceAlias,
			TypeName: "aws_sfn_alias",
		},
		{
			Factory:  DataSourceStateMachine,
			TypeName: "aws_sfn_state_machine",
		},
		{
			Factory:  DataSourceStateMachineVersions,
			TypeName: "aws_sfn_state_machine_versions",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceActivity,
			TypeName: "aws_sfn_activity",
			Name:     "Activity",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceAlias,
			TypeName: "aws_sfn_alias",
		},
		{
			Factory:  ResourceStateMachine,
			TypeName: "aws_sfn_state_machine",
			Name:     "State Machine",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.SFN
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*sfn_sdkv1.SFN, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return sfn_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config["endpoint"].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
