package statuses

type Status int

const (
	Continue           Status = 100
	SwitchingProtocols Status = 101
	Processing         Status = 102
	EarlyHints         Status = 103

	OK                   Status = 200
	Created              Status = 201
	Accepted             Status = 202
	NonAuthoritativeInfo Status = 203
	NoContent            Status = 204
	ResetContent         Status = 205
	PartialContent       Status = 206
	MultiStatus          Status = 207
	AlreadyReported      Status = 208
	IMUsed               Status = 226

	MultipleChoices   Status = 300
	MovedPermanently  Status = 301
	Found             Status = 302
	SeeOther          Status = 303
	NotModified       Status = 304
	UseProxy          Status = 305
	_                 Status = 306
	TemporaryRedirect Status = 307
	PermanentRedirect Status = 308

	BadRequest                   Status = 400
	Unauthorized                 Status = 401
	PaymentRequired              Status = 402
	Forbidden                    Status = 403
	NotFound                     Status = 404
	MethodNotAllowed             Status = 405
	NotAcceptable                Status = 406
	ProxyAuthRequired            Status = 407
	RequestTimeout               Status = 408
	Conflict                     Status = 409
	Gone                         Status = 410
	LengthRequired               Status = 411
	PreconditionFailed           Status = 412
	RequestEntityTooLarge        Status = 413
	RequestURITooLong            Status = 414
	UnsupportedMediaType         Status = 415
	RequestedRangeNotSatisfiable Status = 416
	ExpectationFailed            Status = 417
	Teapot                       Status = 418
	MisdirectedRequest           Status = 421
	UnprocessableEntity          Status = 422
	Locked                       Status = 423
	FailedDependency             Status = 424
	TooEarly                     Status = 425
	UpgradeRequired              Status = 426
	PreconditionRequired         Status = 428
	TooManyRequests              Status = 429
	RequestHeaderFieldsTooLarge  Status = 431
	UnavailableForLegalReasons   Status = 451

	InternalServerError           Status = 500
	NotImplemented                Status = 501
	BadGateway                    Status = 502
	ServiceUnavailable            Status = 503
	GatewayTimeout                Status = 504
	HTTPVersionNotSupported       Status = 505
	VariantAlsoNegotiates         Status = 506
	InsufficientStorage           Status = 507
	LoopDetected                  Status = 508
	NotExtended                   Status = 510
	NetworkAuthenticationRequired Status = 511
)
