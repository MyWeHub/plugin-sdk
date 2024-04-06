package apimanagement

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"os"
)

var (
	logger *zap.Logger
	tracer trace.Tracer
)

func SetTelemetry(l *zap.Logger, t trace.Tracer) {
	logger = l
	tracer = t
}

type Adapter struct {
	cf          *armapimanagement.ClientFactory
	serviceName string
	rgN         string
}

func New() (*Adapter, error) {
	// AZURE_SUBSCRIPTION_ID
	// API_MANAGEMENT_SERVICE_NAME
	// RESOURCE_GROUP_NAME
	subID := os.Getenv("AZURE_SUBSCRIPTION_ID")
	if subID == "" {
		return nil, fmt.Errorf("env 'AZURE_SUBSCRIPTION_ID' not found")
	}
	serviceName := os.Getenv("API_MANAGEMENT_SERVICE_NAME")
	if serviceName == "" {
		return nil, fmt.Errorf("env 'API_MANAGEMENT_SERVICE_NAME' not found")
	}
	rgN := os.Getenv("RESOURCE_GROUP_NAME")
	if rgN == "" {
		return nil, fmt.Errorf("env 'RESOURCE_GROUP_NAME' not found")
	}

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, err
	}

	clientFactory, err := armapimanagement.NewClientFactory(subID, cred, nil)
	if err != nil {
		return nil, err
	}

	return &Adapter{
		cf:          clientFactory,
		serviceName: serviceName,
		rgN:         rgN,
	}, nil
}

// CreateSubscription
// apiID is part of create apigroup response
// notice! the api must exist at first before calling this
// so before this call CreateApiGroup
func (a *Adapter) CreateSubscription(ctx context.Context, orgId string) (*armapimanagement.SubscriptionClientCreateOrUpdateResponse, error) {
	scope := fmt.Sprintf("/apis/%s", createAPIID(orgId))
	sr, err := a.cf.NewSubscriptionClient().CreateOrUpdate(ctx, a.rgN, a.serviceName, orgId, armapimanagement.SubscriptionCreateParameters{
		Properties: &armapimanagement.SubscriptionCreateParameterProperties{
			DisplayName: &[]string{orgId}[0],
			Scope:       &scope,
		},
	}, &armapimanagement.SubscriptionClientCreateOrUpdateOptions{Notify: nil,
		IfMatch: nil,
		AppType: nil,
	})
	if err != nil {
		return nil, err
	}
	return &sr, nil
}
func (a *Adapter) CreateApiGroup(ctx context.Context, orgID string) (*armapimanagement.APIContract, error) {
	apTp := armapimanagement.APITypeHTTP
	name := createAPIID(orgID)
	poller, err := a.cf.NewAPIClient().BeginCreateOrUpdate(ctx, a.rgN, a.serviceName, name, armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			SubscriptionKeyParameterNames: &armapimanagement.SubscriptionKeyParameterNamesContract{
				Header: strToPtr("Ocp-Apim-Subscription-Key"),
				Query:  strToPtr("subscription-key"),
			},
			APIType:     &apTp,
			Path:        &[]string{orgID}[0],
			DisplayName: strToPtr(name),
			Protocols: []*armapimanagement.Protocol{
				&[]armapimanagement.Protocol{armapimanagement.ProtocolHTTPS}[0]},
			ServiceURL:           strToPtr("https://dashboard-uat.weconnecthub.net/module/plugin-http-operations/api/v2/trigger/"),
			SubscriptionRequired: &[]bool{true}[0],
		},
	}, &armapimanagement.APIClientBeginCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		return nil, err
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &res.APIContract, nil
}
func (a *Adapter) CreateAPIOperation(ctx context.Context, orgID, httpNodeId, transformationID string) error {
	_, err := a.cf.NewAPIOperationClient().CreateOrUpdate(ctx, a.rgN, a.serviceName, createAPIID(orgID), httpNodeId, armapimanagement.OperationContract{
		Properties: &armapimanagement.OperationContractProperties{
			Method:      strToPtr("POST"),
			DisplayName: strToPtr(httpNodeId),
			URLTemplate: strToPtr(fmt.Sprintf("/%s/%s", transformationID, httpNodeId)),
		},
	}, &armapimanagement.APIOperationClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		return err
	}
	_, err = a.cf.NewAPIOperationPolicyClient().CreateOrUpdate(ctx, a.rgN, a.serviceName, createAPIID(orgID), httpNodeId, armapimanagement.PolicyIDNamePolicy,
		armapimanagement.PolicyContract{
			Properties: &armapimanagement.PolicyContractProperties{
				Value: strToPtr(
					fmt.Sprintf(
						`<policies><inbound><base /><set-header name="cid" exists-action="override"><value>%s</value></set-header></inbound><backend><base /></backend><outbound><base /></outbound><on-error><base /></on-error></policies>`,
						orgID),
				),
				Format: &[]armapimanagement.PolicyContentFormat{armapimanagement.PolicyContentFormatXML}[0],
			},
		}, nil,
	)
	if err != nil {
		return err
	}
	return nil
}

func (a *Adapter) GetSubscriptionKeys(ctx context.Context, orgId string) (armapimanagement.SubscriptionClientListSecretsResponse, error) {
	subResp, err := a.cf.NewSubscriptionClient().ListSecrets(ctx, a.rgN, a.serviceName, orgId, nil)
	if err != nil {
		return armapimanagement.SubscriptionClientListSecretsResponse{}, nil // TODO: Any reason Mohammad is doing this?
	}
	return subResp, nil
}

func createAPIID(orgID string) string {
	return fmt.Sprintf("http-trigger-%s", orgID)
}

func strToPtr(s string) *string {
	return &s
}
