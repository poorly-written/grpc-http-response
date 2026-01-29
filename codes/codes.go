package codes

import (
	"sync"

	grpcCodes "google.golang.org/grpc/codes"
)

type Code interface {
	IsError() bool
	HttpCode() int
	GrpcCode() grpcCodes.Code
}

type code struct {
	http int
	grpc grpcCodes.Code
}

func (c code) IsError() bool {
	return c.grpc != grpcCodes.OK
}

func (c code) HttpCode() int {
	return c.http
}

func (c code) GrpcCode() grpcCodes.Code {
	return c.grpc
}

var Ok = code{http: 200, grpc: grpcCodes.OK}
var Created = code{http: 201, grpc: grpcCodes.OK}
var Accepted = code{http: 202, grpc: grpcCodes.OK}
var NoContent = code{http: 204, grpc: grpcCodes.OK}
var PartialContent = code{http: 206, grpc: grpcCodes.OK}

var BadRequest = code{http: 400, grpc: grpcCodes.InvalidArgument}
var Unauthorized = code{http: 401, grpc: grpcCodes.Unauthenticated}
var PaymentRequired = code{http: 402, grpc: grpcCodes.FailedPrecondition}
var Forbidden = code{http: 403, grpc: grpcCodes.PermissionDenied}
var NotFound = code{http: 404, grpc: grpcCodes.NotFound}
var MethodNotAllowed = code{http: 405, grpc: grpcCodes.Unimplemented}
var NotAcceptable = code{http: 406, grpc: grpcCodes.InvalidArgument}
var ProxyAuthenticationRequired = code{http: 407, grpc: grpcCodes.FailedPrecondition}
var RequestTimeout = code{http: 408, grpc: grpcCodes.DeadlineExceeded}
var Conflict = code{http: 409, grpc: grpcCodes.AlreadyExists}
var Gone = code{http: 410, grpc: grpcCodes.NotFound}
var LengthRequired = code{http: 411, grpc: grpcCodes.FailedPrecondition}
var PreconditionFailed = code{http: 412, grpc: grpcCodes.FailedPrecondition}
var RequestEntityTooLarge = code{http: 413, grpc: grpcCodes.FailedPrecondition}
var RequestUriTooLong = code{http: 414, grpc: grpcCodes.FailedPrecondition}
var UnsupportedMediaType = code{http: 415, grpc: grpcCodes.FailedPrecondition}
var RequestedRangeNotSatisfiable = code{http: 416, grpc: grpcCodes.FailedPrecondition}
var ExpectationFailed = code{http: 417, grpc: grpcCodes.FailedPrecondition}
var IAmATeapot = code{http: 418, grpc: grpcCodes.FailedPrecondition}
var MisdirectedRequest = code{http: 421, grpc: grpcCodes.FailedPrecondition}
var UnprocessableEntity = code{http: 422, grpc: grpcCodes.InvalidArgument}
var Locked = code{http: 423, grpc: grpcCodes.FailedPrecondition}
var FailedDependency = code{http: 424, grpc: grpcCodes.FailedPrecondition}
var TooEarly = code{http: 425, grpc: grpcCodes.FailedPrecondition}
var UpgradeRequired = code{http: 426, grpc: grpcCodes.FailedPrecondition}
var PreconditionRequired = code{http: 428, grpc: grpcCodes.FailedPrecondition}
var TooManyRequests = code{http: 429, grpc: grpcCodes.ResourceExhausted}
var RequestHeaderFieldsTooLarge = code{http: 431, grpc: grpcCodes.FailedPrecondition}
var UnavailableForLegalReasons = code{http: 451, grpc: grpcCodes.FailedPrecondition}
var ClientClosedRequest = code{http: 499, grpc: grpcCodes.Canceled}

var InternalServerError = code{http: 500, grpc: grpcCodes.Internal}
var NotImplemented = code{http: 501, grpc: grpcCodes.Unimplemented}
var BadGateway = code{http: 502, grpc: grpcCodes.Unavailable}
var ServiceUnavailable = code{http: 503, grpc: grpcCodes.Unavailable}
var GatewayTimeout = code{http: 504, grpc: grpcCodes.DeadlineExceeded}
var VersionNotSupported = code{http: 505, grpc: grpcCodes.Unimplemented}
var VariantAlsoNegotiatesExperimental = code{http: 506, grpc: grpcCodes.FailedPrecondition}
var InsufficientStorage = code{http: 507, grpc: grpcCodes.ResourceExhausted}
var LoopDetected = code{http: 508, grpc: grpcCodes.Internal}
var NotExtended = code{http: 510, grpc: grpcCodes.FailedPrecondition}
var NetworkAuthenticationRequired = code{http: 511, grpc: grpcCodes.FailedPrecondition}
var WebServerReturnedAnUnknownError = code{http: 520, grpc: grpcCodes.Unknown}

var defaultErrorCodeIfNotInMapper Code = BadRequest
var defaultCodeIfNotInMapperSetOnce sync.Once

func SetDefaultCodeIfNotInMapper(code Code) {
	defaultCodeIfNotInMapperSetOnce.Do(func() {
		defaultErrorCodeIfNotInMapper = code
	})
}

var codeMapperSetOnce sync.Once
var codeMapperLock sync.Mutex
var codeMapper = map[int]Code{
	int(grpcCodes.OK):                 Ok,
	int(grpcCodes.Canceled):           ClientClosedRequest,
	int(grpcCodes.Unknown):            BadRequest,
	int(grpcCodes.InvalidArgument):    UnprocessableEntity,
	int(grpcCodes.DeadlineExceeded):   GatewayTimeout,
	int(grpcCodes.NotFound):           NotFound,
	int(grpcCodes.AlreadyExists):      BadRequest,
	int(grpcCodes.PermissionDenied):   Forbidden,
	int(grpcCodes.ResourceExhausted):  InternalServerError,
	int(grpcCodes.FailedPrecondition): UnprocessableEntity,
	int(grpcCodes.Aborted):            ClientClosedRequest,
	int(grpcCodes.OutOfRange):         InternalServerError,
	int(grpcCodes.Unimplemented):      InternalServerError,
	int(grpcCodes.Internal):           InternalServerError,
	int(grpcCodes.Unavailable):        ServiceUnavailable,
	int(grpcCodes.DataLoss):           InternalServerError,
	int(grpcCodes.Unauthenticated):    Unauthorized,
	200:                               Ok,
	201:                               Created,
	202:                               Accepted,
	204:                               NoContent,
	206:                               PartialContent,
	400:                               BadRequest,
	401:                               Unauthorized,
	402:                               PaymentRequired,
	403:                               Forbidden,
	404:                               NotFound,
	405:                               MethodNotAllowed,
	406:                               NotAcceptable,
	407:                               ProxyAuthenticationRequired,
	408:                               RequestTimeout,
	409:                               Conflict,
	410:                               Gone,
	411:                               LengthRequired,
	412:                               PreconditionFailed,
	413:                               RequestEntityTooLarge,
	414:                               RequestUriTooLong,
	415:                               UnsupportedMediaType,
	416:                               RequestedRangeNotSatisfiable,
	417:                               ExpectationFailed,
	418:                               IAmATeapot,
	421:                               MisdirectedRequest,
	422:                               UnprocessableEntity,
	423:                               Locked,
	424:                               FailedDependency,
	425:                               TooEarly,
	426:                               UpgradeRequired,
	428:                               PreconditionRequired,
	429:                               TooManyRequests,
	431:                               RequestHeaderFieldsTooLarge,
	451:                               UnavailableForLegalReasons,
	500:                               InternalServerError,
	501:                               NotImplemented,
	502:                               BadGateway,
	503:                               ServiceUnavailable,
	505:                               VersionNotSupported,
	506:                               VariantAlsoNegotiatesExperimental,
	507:                               InsufficientStorage,
	508:                               LoopDetected,
	510:                               NotExtended,
	511:                               NetworkAuthenticationRequired,
	520:                               WebServerReturnedAnUnknownError,
}

func SetCodeMapper(mapper map[int]Code) {
	codeMapperSetOnce.Do(func() {
		codeMapper = mapper
	})
}

func AddToCodeMapper(when int, code Code) {
	codeMapperLock.Lock()
	defer codeMapperLock.Unlock()
	codeMapper[when] = code
}

func Find(v int) Code {
	if v, ok := codeMapper[v]; ok {
		return v
	}

	return defaultErrorCodeIfNotInMapper
}
