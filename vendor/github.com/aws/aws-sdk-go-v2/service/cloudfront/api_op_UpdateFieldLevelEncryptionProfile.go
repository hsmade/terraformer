// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFieldLevelEncryptionProfileInput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfileConfig"`

	// Request to update a field-level encryption profile.
	//
	// FieldLevelEncryptionProfileConfig is a required field
	FieldLevelEncryptionProfileConfig *FieldLevelEncryptionProfileConfig `locationName:"FieldLevelEncryptionProfileConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`

	// The ID of the field-level encryption profile request.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the profile
	// identity to update. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s UpdateFieldLevelEncryptionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFieldLevelEncryptionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFieldLevelEncryptionProfileInput"}

	if s.FieldLevelEncryptionProfileConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("FieldLevelEncryptionProfileConfig"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.FieldLevelEncryptionProfileConfig != nil {
		if err := s.FieldLevelEncryptionProfileConfig.Validate(); err != nil {
			invalidParams.AddNested("FieldLevelEncryptionProfileConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFieldLevelEncryptionProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryptionProfileConfig != nil {
		v := s.FieldLevelEncryptionProfileConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2019-03-26/"}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfileConfig", v, metadata)
	}
	return nil
}

type UpdateFieldLevelEncryptionProfileOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryptionProfile"`

	// The result of the field-level encryption profile request.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the results of updating the profile.
	FieldLevelEncryptionProfile *FieldLevelEncryptionProfile `type:"structure"`
}

// String returns the string representation
func (s UpdateFieldLevelEncryptionProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFieldLevelEncryptionProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.FieldLevelEncryptionProfile != nil {
		v := s.FieldLevelEncryptionProfile

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "FieldLevelEncryptionProfile", v, metadata)
	}
	return nil
}

const opUpdateFieldLevelEncryptionProfile = "UpdateFieldLevelEncryptionProfile2019_03_26"

// UpdateFieldLevelEncryptionProfileRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Update a field-level encryption profile.
//
//    // Example sending a request using UpdateFieldLevelEncryptionProfileRequest.
//    req := client.UpdateFieldLevelEncryptionProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2019-03-26/UpdateFieldLevelEncryptionProfile
func (c *Client) UpdateFieldLevelEncryptionProfileRequest(input *UpdateFieldLevelEncryptionProfileInput) UpdateFieldLevelEncryptionProfileRequest {
	op := &aws.Operation{
		Name:       opUpdateFieldLevelEncryptionProfile,
		HTTPMethod: "PUT",
		HTTPPath:   "/2019-03-26/field-level-encryption-profile/{Id}/config",
	}

	if input == nil {
		input = &UpdateFieldLevelEncryptionProfileInput{}
	}

	req := c.newRequest(op, input, &UpdateFieldLevelEncryptionProfileOutput{})
	return UpdateFieldLevelEncryptionProfileRequest{Request: req, Input: input, Copy: c.UpdateFieldLevelEncryptionProfileRequest}
}

// UpdateFieldLevelEncryptionProfileRequest is the request type for the
// UpdateFieldLevelEncryptionProfile API operation.
type UpdateFieldLevelEncryptionProfileRequest struct {
	*aws.Request
	Input *UpdateFieldLevelEncryptionProfileInput
	Copy  func(*UpdateFieldLevelEncryptionProfileInput) UpdateFieldLevelEncryptionProfileRequest
}

// Send marshals and sends the UpdateFieldLevelEncryptionProfile API request.
func (r UpdateFieldLevelEncryptionProfileRequest) Send(ctx context.Context) (*UpdateFieldLevelEncryptionProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFieldLevelEncryptionProfileResponse{
		UpdateFieldLevelEncryptionProfileOutput: r.Request.Data.(*UpdateFieldLevelEncryptionProfileOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFieldLevelEncryptionProfileResponse is the response type for the
// UpdateFieldLevelEncryptionProfile API operation.
type UpdateFieldLevelEncryptionProfileResponse struct {
	*UpdateFieldLevelEncryptionProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFieldLevelEncryptionProfile request.
func (r *UpdateFieldLevelEncryptionProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}