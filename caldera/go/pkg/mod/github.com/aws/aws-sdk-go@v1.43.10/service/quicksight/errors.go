// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You don't have access to this item. The provided credentials couldn't be
	// validated. You might not be authorized to carry out the request. Make sure
	// that your account is authorized to use the Amazon QuickSight service, that
	// your policies have the correct permissions, and that you are using the correct
	// access keys.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConcurrentUpdatingException for service response error code
	// "ConcurrentUpdatingException".
	//
	// A resource is already in a state that indicates an operation is happening
	// that must complete before a new update can be applied.
	ErrCodeConcurrentUpdatingException = "ConcurrentUpdatingException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Updating or deleting a resource can cause an inconsistent state.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeDomainNotWhitelistedException for service response error code
	// "DomainNotWhitelistedException".
	//
	// The domain specified isn't on the allow list. All domains for embedded dashboards
	// must be added to the approved list by an Amazon QuickSight admin.
	ErrCodeDomainNotWhitelistedException = "DomainNotWhitelistedException"

	// ErrCodeIdentityTypeNotSupportedException for service response error code
	// "IdentityTypeNotSupportedException".
	//
	// The identity type specified isn't supported. Supported identity types include
	// IAM and QUICKSIGHT.
	ErrCodeIdentityTypeNotSupportedException = "IdentityTypeNotSupportedException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// An internal failure occurred.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The NextToken value isn't valid.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// One or more parameters has a value that isn't valid.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A limit is exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePreconditionNotMetException for service response error code
	// "PreconditionNotMetException".
	//
	// One or more preconditions aren't met.
	ErrCodePreconditionNotMetException = "PreconditionNotMetException"

	// ErrCodeResourceExistsException for service response error code
	// "ResourceExistsException".
	//
	// The resource specified already exists.
	ErrCodeResourceExistsException = "ResourceExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// One or more resources can't be found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceUnavailableException for service response error code
	// "ResourceUnavailableException".
	//
	// This resource is currently unavailable.
	ErrCodeResourceUnavailableException = "ResourceUnavailableException"

	// ErrCodeSessionLifetimeInMinutesInvalidException for service response error code
	// "SessionLifetimeInMinutesInvalidException".
	//
	// The number of minutes specified for the lifetime of a session isn't valid.
	// The session lifetime must be 15-600 minutes.
	ErrCodeSessionLifetimeInMinutesInvalidException = "SessionLifetimeInMinutesInvalidException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Access is throttled.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnsupportedPricingPlanException for service response error code
	// "UnsupportedPricingPlanException".
	//
	// This error indicates that you are calling an embedding operation in Amazon
	// QuickSight without the required pricing plan on your Amazon Web Services
	// account. Before you can use embedding for anonymous users, a QuickSight administrator
	// needs to add capacity pricing to Amazon QuickSight. You can do this on the
	// Manage Amazon QuickSight page.
	//
	// After capacity pricing is added, you can use the GetDashboardEmbedUrl (https://docs.aws.amazon.com/quicksight/latest/APIReference/API_GetDashboardEmbedUrl.html)
	// API operation with the --identity-type ANONYMOUS option.
	ErrCodeUnsupportedPricingPlanException = "UnsupportedPricingPlanException"

	// ErrCodeUnsupportedUserEditionException for service response error code
	// "UnsupportedUserEditionException".
	//
	// This error indicates that you are calling an operation on an Amazon QuickSight
	// subscription where the edition doesn't include support for that operation.
	// Amazon Amazon QuickSight currently has Standard Edition and Enterprise Edition.
	// Not every operation and capability is available in every edition.
	ErrCodeUnsupportedUserEditionException = "UnsupportedUserEditionException"

	// ErrCodeUserNotFoundException for service response error code
	// "QuickSightUserNotFoundException".
	//
	// The user with the provided name isn't found. This error can happen in any
	// operation that requires finding a user based on a provided user name, such
	// as DeleteUser, DescribeUser, and so on.
	ErrCodeUserNotFoundException = "QuickSightUserNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                    newErrorAccessDeniedException,
	"ConcurrentUpdatingException":              newErrorConcurrentUpdatingException,
	"ConflictException":                        newErrorConflictException,
	"DomainNotWhitelistedException":            newErrorDomainNotWhitelistedException,
	"IdentityTypeNotSupportedException":        newErrorIdentityTypeNotSupportedException,
	"InternalFailureException":                 newErrorInternalFailureException,
	"InvalidNextTokenException":                newErrorInvalidNextTokenException,
	"InvalidParameterValueException":           newErrorInvalidParameterValueException,
	"LimitExceededException":                   newErrorLimitExceededException,
	"PreconditionNotMetException":              newErrorPreconditionNotMetException,
	"ResourceExistsException":                  newErrorResourceExistsException,
	"ResourceNotFoundException":                newErrorResourceNotFoundException,
	"ResourceUnavailableException":             newErrorResourceUnavailableException,
	"SessionLifetimeInMinutesInvalidException": newErrorSessionLifetimeInMinutesInvalidException,
	"ThrottlingException":                      newErrorThrottlingException,
	"UnsupportedPricingPlanException":          newErrorUnsupportedPricingPlanException,
	"UnsupportedUserEditionException":          newErrorUnsupportedUserEditionException,
	"QuickSightUserNotFoundException":          newErrorUserNotFoundException,
}
