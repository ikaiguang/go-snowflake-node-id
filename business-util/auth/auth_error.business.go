package authutil

import (
	"github.com/go-kratos/kratos/v2/errors"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"

	errorutil "github.com/ikaiguang/go-srv-kit/error"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "Can not get key while signing token")

	ErrInvalidAuthInfo    = errors.BadRequest(commonv1.ERROR_BAD_REQUEST.String(), "ValidateFunc : invalid auth info")
	ErrInvalidAuthUser    = errors.BadRequest(commonv1.ERROR_UNAUTHORIZED.String(), "ValidateFunc : invalid auth user")
	ErrLoginLimit         = errors.Unauthorized(commonv1.ERROR_UNAUTHORIZED.String(), "ValidateFunc : Token is invalid")
	ErrGetRedisData       = errors.BadRequest(commonv1.ERROR_EXPECTATION_FAILED.String(), "RedisCacheData : get redis data failed")
	ErrSetRedisData       = errors.BadRequest(commonv1.ERROR_EXPECTATION_FAILED.String(), "RedisCacheData : set redis data failed")
	ErrUnmarshalRedisData = errors.BadRequest(commonv1.ERROR_EXPECTATION_FAILED.String(), "RedisCacheData : unmarshal redis data failed")
	ErrMarshalRedisData   = errors.BadRequest(commonv1.ERROR_EXPECTATION_FAILED.String(), "RedisCacheData : marshal redis data failed")
)

// Is ...
func Is(err, target error) bool {
	return errorutil.Is(err, target)
}
