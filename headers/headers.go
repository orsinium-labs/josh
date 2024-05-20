package headers

// An HTTP request/response header name.
//
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers
type Header string

const (
	// Defines the authentication method that should be used to access a resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate
	WWWAuthenticate Header = "WWW-Authenticate"

	// Contains the credentials to authenticate a user-agent with a server.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization
	Authorization Header = "Authorization"

	// Defines the authentication method that should be used to access a resource behind a proxy server.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authenticate
	ProxyAuthenticate Header = "Proxy-Authenticate"

	// Contains the credentials to authenticate a user agent with a proxy server.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Proxy-Authorization
	ProxyAuthorization Header = "Proxy-Authorization"

	// The time, in seconds, that the object has been in a proxy cache.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Age
	Age Header = "Age"

	// Directives for caching mechanisms in both requests and responses.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control
	CacheControl Header = "Cache-Control"

	// Clears browsing data (e.g. cookies, storage, cache) associated with the requesting website.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Clear-Site-Data
	ClearSiteData Header = "Clear-Site-Data"

	// The date/time after which the response is considered stale.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expires
	Expires Header = "Expires"

	// Specifies a set of rules that define how a URL's query parameters will affect cache matching. These rules dictate whether the same URL with different URL parameters should be saved as separate browser cache entries.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/No-Vary-Search
	NoVarySearch Header = "No-Vary-Search"

	// The last modification date of the resource, used to compare several versions of the same resource. It is less accurate than {{HTTPHeader("ETag")}}, but easier to calculate in some environments. Conditional requests using {{HTTPHeader("If-Modified-Since")}} and {{HTTPHeader("If-Unmodified-Since")}} use this value to change the behavior of the request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
	LastModified Header = "Last-Modified"

	// A unique string identifying the version of the resource. Conditional requests using {{HTTPHeader("If-Match")}} and {{HTTPHeader("If-None-Match")}} use this value to change the behavior of the request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ETag
	ETag Header = "ETag"

	// Makes the request conditional, and applies the method only if the stored resource matches one of the given ETags.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match
	IfMatch Header = "If-Match"

	// Makes the request conditional, and applies the method only if the stored resource _doesn't_ match any of the given ETags. This is used to update caches (for safe requests), or to prevent uploading a new resource when one already exists.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-None-Match
	IfNoneMatch Header = "If-None-Match"

	// Makes the request conditional, and expects the resource to be transmitted only if it has been modified after the given date. This is used to transmit data only when the cache is out of date.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Modified-Since
	IfModifiedSince Header = "If-Modified-Since"

	// Makes the request conditional, and expects the resource to be transmitted only if it has not been modified after the given date. This ensures the coherence of a new fragment of a specific range with previous ones, or to implement an optimistic concurrency control system when modifying existing documents.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Unmodified-Since
	IfUnmodifiedSince Header = "If-Unmodified-Since"

	// Determines how to match request headers to decide whether a cached response can be used rather than requesting a fresh one from the origin server.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Vary
	Vary Header = "Vary"

	// Controls whether the network connection stays open after the current transaction finishes.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Connection
	Connection Header = "Connection"

	// Controls how long a persistent connection should stay open.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Keep-Alive
	KeepAlive Header = "Keep-Alive"

	// Informs the server about the {{Glossary("MIME_type", "types")}} of data that can be sent back.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept
	Accept Header = "Accept"

	// The encoding algorithm, usually a [compression algorithm](/en-US/docs/Web/HTTP/Compression), that can be used on the resource sent back.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding
	AcceptEncoding Header = "Accept-Encoding"

	// Informs the server about the human language the server is expected to send back. This is a hint and is not necessarily under the full control of the user: the server should always pay attention not to override an explicit user choice (like selecting a language from a dropdown).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Language
	AcceptLanguage Header = "Accept-Language"

	// Indicates expectations that need to be fulfilled by the server to properly handle the request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect
	Expect Header = "Expect"

	// When using [`TRACE`](/en-US/docs/Web/HTTP/Methods/TRACE), indicates the maximum number of hops the request can do before being reflected to the sender.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Max-Forwards
	MaxForwards Header = "Max-Forwards"

	// Contains stored [HTTP cookies](/en-US/docs/Web/HTTP/Cookies) previously sent by the server with the {{HTTPHeader("Set-Cookie")}} header.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cookie
	Cookie Header = "Cookie"

	// Send cookies from the server to the user-agent.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Set-Cookie
	SetCookie Header = "Set-Cookie"

	// Indicates whether the response to the request can be exposed when the credentials flag is true.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Credentials
	AccessControlAllowCredentials Header = "Access-Control-Allow-Credentials"

	// Used in response to a {{Glossary("Preflight_request", "preflight request")}} to indicate which HTTP headers can be used when making the actual request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers
	AccessControlAllowHeaders Header = "Access-Control-Allow-Headers"

	// Specifies the methods allowed when accessing the resource in response to a preflight request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods
	AccessControlAllowMethods Header = "Access-Control-Allow-Methods"

	// Indicates whether the response can be shared.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin
	AccessControlAllowOrigin Header = "Access-Control-Allow-Origin"

	// Indicates which headers can be exposed as part of the response by listing their names.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers
	AccessControlExposeHeaders Header = "Access-Control-Expose-Headers"

	// Indicates how long the results of a preflight request can be cached.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Max-Age
	AccessControlMaxAge Header = "Access-Control-Max-Age"

	// Used when issuing a preflight request to let the server know which HTTP headers will be used when the actual request is made.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Headers
	AccessControlRequestHeaders Header = "Access-Control-Request-Headers"

	// Used when issuing a preflight request to let the server know which [HTTP method](/en-US/docs/Web/HTTP/Methods) will be used when the actual request is made.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Request-Method
	AccessControlRequestMethod Header = "Access-Control-Request-Method"

	// Indicates where a fetch originates from.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Origin
	Origin Header = "Origin"

	// Specifies origins that are allowed to see values of attributes retrieved via features of the [Resource Timing API](/en-US/docs/Web/API/Performance_API/Resource_timing), which would otherwise be reported as zero due to cross-origin restrictions.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Timing-Allow-Origin
	TimingAllowOrigin Header = "Timing-Allow-Origin"

	// Indicates if the resource transmitted should be displayed inline (default behavior without the header), or if it should be handled like a download and the browser should present a "Save As" dialog.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	ContentDisposition Header = "Content-Disposition"

	// The size of the resource, in decimal number of bytes.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length
	ContentLength Header = "Content-Length"

	// Indicates the media type of the resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type
	ContentType Header = "Content-Type"

	// Used to specify the compression algorithm.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Encoding
	ContentEncoding Header = "Content-Encoding"

	// Describes the human language(s) intended for the audience, so that it allows a user to differentiate according to the users' own preferred language.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Language
	ContentLanguage Header = "Content-Language"

	// Indicates an alternate location for the returned data.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Location
	ContentLocation Header = "Content-Location"

	// Contains information from the client-facing side of proxy servers that is altered or lost when a proxy is involved in the path of the request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Forwarded
	Forwarded Header = "Forwarded"

	// Added by proxies, both forward and reverse proxies, and can appear in the request headers and the response headers.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Via
	Via Header = "Via"

	// Indicates the URL to redirect a page to.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Location
	Location Header = "Location"

	// Directs the browser to reload the page or redirect to another. Takes the same value as the `meta` element with [`http-equiv="refresh"`](/en-US/docs/Web/HTML/Element/meta#http-equiv).
	Refresh Header = "Refresh"

	// Contains an Internet email address for a human user who controls the requesting user agent.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/From
	From Header = "From"

	// Specifies the domain name of the server (for virtual hosting), and (optionally) the TCP port number on which the server is listening.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host
	Host Header = "Host"

	// The address of the previous web page from which a link to the currently requested page was followed.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referer
	Referer Header = "Referer"

	// Governs which referrer information sent in the {{HTTPHeader("Referer")}} header should be included with requests made.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
	ReferrerPolicy Header = "Referrer-Policy"

	// Contains a characteristic string that allows the network protocol peers to identify the application type, operating system, software vendor or software version of the requesting software user agent.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/User-Agent
	UserAgent Header = "User-Agent"

	// Lists the set of HTTP request methods supported by a resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Allow
	Allow Header = "Allow"

	// Contains information about the software used by the origin server to handle the request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server
	Server Header = "Server"

	// Indicates if the server supports range requests, and if so in which unit the range can be expressed.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Ranges
	AcceptRanges Header = "Accept-Ranges"

	// Indicates the part of a document that the server should return.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Range
	Range Header = "Range"

	// Creates a conditional range request that is only fulfilled if the given etag or date matches the remote resource. Used to prevent downloading two ranges from incompatible version of the resource.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Range
	IfRange Header = "If-Range"

	// Indicates where in a full body message a partial message belongs.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range
	ContentRange Header = "Content-Range"

	// Allows a server to declare an embedder policy for a given document.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Embedder-Policy
	CrossOriginEmbedderPolicy Header = "Cross-Origin-Embedder-Policy"

	// Prevents other domains from opening/controlling a window.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Opener-Policy
	CrossOriginOpenerPolicy Header = "Cross-Origin-Opener-Policy"

	// Prevents other domains from reading the response of the resources to which this header is applied. See also [CORP explainer article](/en-US/docs/Web/HTTP/Cross-Origin_Resource_Policy).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cross-Origin-Resource-Policy
	CrossOriginResourcePolicy Header = "Cross-Origin-Resource-Policy"

	// Controls resources the user agent is allowed to load for a given page.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
	ContentSecurityPolicy Header = "Content-Security-Policy"

	// Allows web developers to experiment with policies by monitoring, but not enforcing, their effects. These violation reports consist of {{Glossary("JSON")}} documents sent via an HTTP `POST` request to the specified URI.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy-Report-Only
	ContentSecurityPolicyReportOnly Header = "Content-Security-Policy-Report-Only"

	// Provides a mechanism to allow and deny the use of browser features in a website's own frame, and in {{htmlelement("iframe")}}s that it embeds.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Permissions-Policy
	PermissionsPolicy Header = "Permissions-Policy"

	// Force communication using HTTPS instead of HTTP.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
	StrictTransportSecurity Header = "Strict-Transport-Security"

	// Sends a signal to the server expressing the client's preference for an encrypted and authenticated response, and that it can successfully handle the {{CSP("upgrade-insecure-requests")}} directive.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade-Insecure-Requests
	UpgradeInsecureRequests Header = "Upgrade-Insecure-Requests"

	// Disables MIME sniffing and forces browser to use the type given in {{HTTPHeader("Content-Type")}}.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
	XContentTypeOptions Header = "X-Content-Type-Options"

	// Indicates whether a browser should be allowed to render a page in a {{HTMLElement("frame")}}, {{HTMLElement("iframe")}}, {{HTMLElement("embed")}} or {{HTMLElement("object")}}.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
	XFrameOptions Header = "X-Frame-Options"

	// Specifies if a cross-domain policy file (`crossdomain.xml`) is allowed. The file may define a policy to grant clients, such as Adobe's Flash Player (now obsolete), Adobe Acrobat, Microsoft Silverlight (now obsolete), or Apache Flex, permission to handle data across domains that would otherwise be restricted due to the [Same-Origin Policy](/en-US/docs/Web/Security/Same-origin_policy). See the [Cross-domain Policy File Specification](https://www.adobe.com/devnet-docs/acrobatetk/tools/AppSec/CrossDomain_PolicyFile_Specification.pdf) for more information.
	XPermittedCrossDomainPolicies Header = "X-Permitted-Cross-Domain-Policies"

	// May be set by hosting environments or other frameworks and contains information about them while not providing any usefulness to the application or its visitors. Unset this header to avoid exposing potential vulnerabilities.
	XPoweredBy Header = "X-Powered-By"

	// Enables cross-site scripting filtering.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection
	XXSSProtection Header = "X-XSS-Protection"

	// Indicates the relationship between a request initiator's origin and its target's origin. It is a Structured Header whose value is a token with possible values `cross-site`, `same-origin`, `same-site`, and `none`.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Site
	SecFetchSite Header = "Sec-Fetch-Site"

	// Indicates the request's mode to a server. It is a Structured Header whose value is a token with possible values `cors`, `navigate`, `no-cors`, `same-origin`, and `websocket`.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Mode
	SecFetchMode Header = "Sec-Fetch-Mode"

	// Indicates whether or not a navigation request was triggered by user activation. It is a Structured Header whose value is a boolean so possible values are `?0` for false and `?1` for true.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-User
	SecFetchUser Header = "Sec-Fetch-User"

	// Indicates the request's destination. It is a Structured Header whose value is a token with possible values `audio`, `audioworklet`, `document`, `embed`, `empty`, `font`, `image`, `manifest`, `object`, `paintworklet`, `report`, `script`, `serviceworker`, `sharedworker`, `style`, `track`, `video`, `worker`, and `xslt`.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Fetch-Dest
	SecFetchDest Header = "Sec-Fetch-Dest"

	// Indicates the purpose of the request, when the purpose is something other than immediate use by the user-agent. The header currently has one possible value, `prefetch`, which indicates that the resource is being fetched preemptively for a possible future navigation.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-Purpose
	SecPurpose Header = "Sec-Purpose"

	// A request header sent in preemptive request to {{domxref("fetch()")}} a resource during service worker boot. The value, which is set with {{domxref("NavigationPreloadManager.setHeaderValue()")}}, can be used to inform a server that a different resource should be returned than in a normal `fetch()` operation.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Service-Worker-Navigation-Preload
	ServiceWorkerNavigationPreload Header = "Service-Worker-Navigation-Preload"

	// Used to specify a server endpoint for the browser to send warning and error reports to.
	ReportTo Header = "Report-To"

	// Specifies the form of encoding used to safely transfer the resource to the user.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Transfer-Encoding
	TransferEncoding Header = "Transfer-Encoding"

	// Specifies the transfer encodings the user agent is willing to accept.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/TE
	TE Header = "TE"

	// Allows the sender to include additional fields at the end of chunked message.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Trailer
	Trailer Header = "Trailer"

	// Used to list alternate ways to reach this service.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Svc
	AltSvc Header = "Alt-Svc"

	// Used to identify the alternative service in use.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Used
	AltUsed Header = "Alt-Used"

	// Contains the date and time at which the message was originated.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Date
	Date Header = "Date"

	// This entity-header field provides a means for serializing one or more links in HTTP headers. It is semantically equivalent to the HTML {{HTMLElement("link")}} element.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Link
	Link Header = "Link"

	// Indicates how long the user agent should wait before making a follow-up request.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
	RetryAfter Header = "Retry-After"

	// Communicates one or more metrics and descriptions for the given request-response cycle.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Server-Timing
	ServerTiming Header = "Server-Timing"

	// Links generated code to a [source map](https://firefox-source-docs.mozilla.org/devtools-user/debugger/how_to/use_a_source_map/index.html).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/SourceMap
	SourceMap Header = "SourceMap"

	// This HTTP/1.1 (only) header can be used to upgrade an already established client/server connection to a different protocol (over the same transport protocol). For example, it can be used by a client to upgrade a connection from HTTP 1.1 to HTTP 2.0, or an HTTP or HTTPS connection into a WebSocket.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Upgrade
	Upgrade Header = "Upgrade"

	// Servers can advertise support for Client Hints using the `Accept-CH` header field or an equivalent HTML `<meta>` element with [`http-equiv`](/en-US/docs/Web/HTML/Element/meta#http-equiv) attribute.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-CH
	AcceptCH Header = "Accept-CH"

	// Servers use `Critical-CH` along with {{HttpHeader("Accept-CH")}} to specify that accepted client hints are also [critical client hints](/en-US/docs/Web/HTTP/Client_hints#critical_client_hints).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Critical-CH
	CriticalCH Header = "Critical-CH"

	// User agent's branding and version.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA
	SecCHUA Header = "Sec-CH-UA"

	// User agent's underlying platform architecture.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Arch
	SecCHUAArch Header = "Sec-CH-UA-Arch"

	// User agent's underlying CPU architecture bitness (for example "64" bit).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Bitness
	SecCHUABitness Header = "Sec-CH-UA-Bitness"

	// Full version for each brand in the user agent's brand list.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Full-Version-List
	SecCHUAFullVersionList Header = "Sec-CH-UA-Full-Version-List"

	// User agent is running on a mobile device or, more generally, prefers a "mobile" user experience.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Mobile
	SecCHUAMobile Header = "Sec-CH-UA-Mobile"

	// User agent's device model.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Model
	SecCHUAModel Header = "Sec-CH-UA-Model"

	// User agent's underlying operation system/platform.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Platform
	SecCHUAPlatform Header = "Sec-CH-UA-Platform"

	// User agent's underlying operation system version.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-CH-UA-Platform-Version
	SecCHUAPlatformVersion Header = "Sec-CH-UA-Platform-Version"

	// User's preference of dark or light color scheme.
	SecCHUAPrefersColorScheme Header = "Sec-CH-UA-Prefers-Color-Scheme"

	// User's preference to see fewer animations and content layout shifts.
	SecCHUAPrefersReducedMotion Header = "Sec-CH-UA-Prefers-Reduced-Motion"

	// Approximate amount of available client RAM memory. This is part of the [Device Memory API](/en-US/docs/Web/API/Device_Memory_API).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Device-Memory
	DeviceMemory Header = "Device-Memory"

	// Approximate bandwidth of the client's connection to the server, in Mbps. This is part of the [Network Information API](/en-US/docs/Web/API/Network_Information_API).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Downlink
	Downlink Header = "Downlink"

	// The {{Glossary("effective connection type")}} ("network profile") that best matches the connection's latency and bandwidth. This is part of the [Network Information API](/en-US/docs/Web/API/Network_Information_API).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/ECT
	ECT Header = "ECT"

	// Application layer round trip time (RTT) in milliseconds, which includes the server processing time. This is part of the [Network Information API](/en-US/docs/Web/API/Network_Information_API).
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/RTT
	RTT Header = "RTT"

	// A string `on` that indicates the user agent's preference for reduced data usage.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Save-Data
	SaveData Header = "Save-Data"

	// Indicates whether the user consents to a website or service selling or sharing their personal information with third parties.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Sec-GPC
	SecGPC Header = "Sec-GPC"

	// Provides a mechanism to allow web applications to isolate their origins.
	OriginIsolation Header = "Origin-Isolation"

	// Defines a mechanism that enables developers to declare a network error reporting policy.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/NEL
	NEL Header = "NEL"

	// A client can express the desired push policy for a request by sending an [`Accept-Push-Policy`](https://datatracker.ietf.org/doc/html/draft-ruellan-http-accept-push-policy-00#section-3.1) header field in the request.
	AcceptPushPolicy Header = "Accept-Push-Policy"

	// A client can send the [`Accept-Signature`](https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#name-the-accept-signature-header) header field to indicate intention to take advantage of any available signatures and to indicate what kinds of signatures it supports.
	AcceptSignature Header = "Accept-Signature"

	// Indicates that the request has been conveyed in TLS early data.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Early-Data
	EarlyData Header = "Early-Data"

	// A [`Push-Policy`](https://datatracker.ietf.org/doc/html/draft-ruellan-http-accept-push-policy-00#section-3.2) defines the server behavior regarding push when processing a request.
	PushPolicy Header = "Push-Policy"

	// The [`Signature`](https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#name-the-signature-header) header field conveys a list of signatures for an exchange, each one accompanied by information about how to determine the authority of and refresh that signature.
	Signature Header = "Signature"

	// The [`Signed-Headers`](https://wicg.github.io/webpackage/draft-yasskin-http-origin-signed-responses.html#name-the-signed-headers-header) header field identifies an ordered list of response header fields to include in a signature.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Signed-Headers
	SignedHeaders Header = "Signed-Headers"

	// Provides a list of URLs pointing to text resources containing [speculation rule](/en-US/docs/Web/API/Speculation_Rules_API) JSON definitions. When the response is an HTML document, these rules will be added to the document's speculation rule set.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Speculation-Rules
	SpeculationRules Header = "Speculation-Rules"

	// Set by a navigation target to opt-in to using various higher-risk loading modes. For example, cross-origin, same-site [prerendering](/en-US/docs/Web/API/Speculation_Rules_API#using_prerendering) requires a `Supports-Loading-Mode` value of `credentialed-prerender`.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Supports-Loading-Mode
	SupportsLoadingMode Header = "Supports-Loading-Mode"

	// Identifies the originating IP addresses of a client connecting to a web server through an HTTP proxy or a load balancer.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-For
	XForwardedFor Header = "X-Forwarded-For"

	// Identifies the original host requested that a client used to connect to your proxy or load balancer.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host
	XForwardedHost Header = "X-Forwarded-Host"

	// Identifies the protocol (HTTP or HTTPS) that a client used to connect to your proxy or load balancer.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Forwarded-Proto
	XForwardedProto Header = "X-Forwarded-Proto"

	// Controls DNS prefetching, a feature by which browsers proactively perform domain name resolution on both links that the user may choose to follow as well as URLs for items referenced by the document, including images, CSS, JavaScript, and so forth.
	XDNSPrefetchControl Header = "X-DNS-Prefetch-Control"

	// The [`X-Robots-Tag`](https://developers.google.com/search/docs/advanced/robots/robots_meta_tag) HTTP header is used to indicate how a web page is to be indexed within public search engine results. The header is effectively equivalent to `<meta name="robots" content="…">`.
	XRobotsTag Header = "X-Robots-Tag"

	// Implementation-specific header that may have various effects anywhere along the request-response chain. Used for backwards compatibility with HTTP/1.0 caches where the `Cache-Control` header is not yet present.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Pragma
	Pragma Header = "Pragma"

	// General warning information about possible problems.
	//
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Warning
	Warning Header = "Warning"
)
