// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RestoreSecretInput struct {
	_ struct{} `type:"structure"`

	// Specifies the secret that you want to restore from a previously scheduled
	// deletion. You can specify either the Amazon Resource Name (ARN) or the friendly
	// name of the secret.
	//
	// If you specify an ARN, we generally recommend that you specify a complete
	// ARN. You can specify a partial ARN too—for example, if you don’t include
	// the final hyphen and six random characters that Secrets Manager adds at the
	// end of the ARN when you created the secret. A partial ARN match can work
	// as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as
	// a partial ARN, then those characters cause Secrets Manager to assume that
	// you’re specifying a complete ARN. This confusion can cause unexpected results.
	// To avoid this situation, we recommend that you don’t create secret names
	// that end with a hyphen followed by six characters.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreSecretInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreSecretInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreSecretInput"}

	if s.SecretId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecretId"))
	}
	if s.SecretId != nil && len(*s.SecretId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RestoreSecretOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the secret that was restored.
	ARN *string `min:"20" type:"string"`

	// The friendly name of the secret that was restored.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s RestoreSecretOutput) String() string {
	return awsutil.Prettify(s)
}

const opRestoreSecret = "RestoreSecret"

// RestoreSecretRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Cancels the scheduled deletion of a secret by removing the DeletedDate time
// stamp. This makes the secret accessible to query once again.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:RestoreSecret
//
// Related operations
//
//    * To delete a secret, use DeleteSecret.
//
//    // Example sending a request using RestoreSecretRequest.
//    req := client.RestoreSecretRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/RestoreSecret
func (c *Client) RestoreSecretRequest(input *RestoreSecretInput) RestoreSecretRequest {
	op := &aws.Operation{
		Name:       opRestoreSecret,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreSecretInput{}
	}

	req := c.newRequest(op, input, &RestoreSecretOutput{})
	return RestoreSecretRequest{Request: req, Input: input, Copy: c.RestoreSecretRequest}
}

// RestoreSecretRequest is the request type for the
// RestoreSecret API operation.
type RestoreSecretRequest struct {
	*aws.Request
	Input *RestoreSecretInput
	Copy  func(*RestoreSecretInput) RestoreSecretRequest
}

// Send marshals and sends the RestoreSecret API request.
func (r RestoreSecretRequest) Send(ctx context.Context) (*RestoreSecretResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreSecretResponse{
		RestoreSecretOutput: r.Request.Data.(*RestoreSecretOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreSecretResponse is the response type for the
// RestoreSecret API operation.
type RestoreSecretResponse struct {
	*RestoreSecretOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreSecret request.
func (r *RestoreSecretResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}