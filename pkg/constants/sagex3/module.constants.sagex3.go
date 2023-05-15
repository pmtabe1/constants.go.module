package sagex3

// var (
// 	SoapContextKey = "publicName"
// )

const (
	SOAP_EMPTY_PAYLOAD_LENGTH_SIZE = 50
	SOAP_EMPTY_PAYLOAD_DATA        = `<?xml version="1.0" encoding="UTF-8"?><RESULT/>`

	//Headers
	HeaderIntentKey     = "intent"
	HeaderUniqueKey     = "unique"
	HeaderActionKey     = "action"
	HeaderDateFormatKey = "dateformat"
	HeaderDateLookupKey = "datekey"
	HeaderUserKey       = "userkey"
	HeaderWebhookUrl    = "webhookurl"
	Sagex3DateFormat    = "DD/MM/YYYY"

	// Empty Payload
	EmptyPayload       = ""
	EmptyPayloadObject = `{}`
	EmptyPayloadArray  = `[]`

	// Integration Error Constans
	SageX3HostPort                         = "localhost:8124"
	IntegrationHostPort                    = "localhost:9000"
	IntegrationAppName                     = "sagex3"
	IntegrationPortNumber                  = "9000"
	ERPPortNumber                          = "8124"
	ERPProdHostOrIP                        = "172.16.100.12"
	ERPDevHostOrIP                         = "172.16.100.12"
	IntegrationHostOrIP                    = "localhost"
	IntegrationDefaultLocalAddress         = "localhost:9000"
	IntegrationAPISageX3Origins            = "localhost:9000;"
	IntegrationName                        = "SageX3Integration"
	IntegrationAPIBaseEndpoint             = "/api/v1/integration"
	IntegrationAPISageX3Endpoint           = "/sagex3/"
	IntegrationAPISageX3EndpointIdentifier = "sagex3"
	IntegrationAPIEndpointIdentifier       = "sagex3"

	IntegrationAPIurlTemplate             = "%v://%v:%v/%v"    // protocol://hostname:port/baseendpoint/
	IntegrationAPIurlWithEndpointTemplate = "%v://%v:%v/%v/%v" // protocol://hostname:port/baseendpoint/resource

	IntegrationAPIUrl             = "http://" + IntegrationHostOrIP + ":" + IntegrationPortNumber + IntegrationAPIBaseEndpoint + IntegrationAPISageX3Endpoint
	IntegrationAPISecureUrl       = "https://" + IntegrationHostOrIP + ":" + IntegrationPortNumber + IntegrationAPIBaseEndpoint + IntegrationAPISageX3Endpoint
	IntegrationServerHostHttpUrl  = IntegrationAPIUrl
	IntegrationServerHostHttpsUrl = IntegrationAPISecureUrl

	ErrorFailureNilDueToNilRequest                        = "Nil Request Provided Please Fix and Try AGAIN"
	ErrorFailureNilDueToPayloadValidationRequest          = "Invalid Payload  Request Provided Please Fix and Try AGAIN"
	ErrorFailureNilDueToEmptyPayloadInARequest            = "Empty Payload in  Request object Provided Please Fix and Try AGAIN"
	ErrorFailureNilIntegrationObject                      = "Integration object is nil Please provide it"
	ErrorFailureNilIntegrationObjectTemplate              = "Integration  %v Object is nil Please provide it"
	ErrorFailureNilIntegrationObjectInstanceTemplate      = "Integration  %v Object Instance is nil Please provide it"
	ErrorFailureNilIntegrationObjectKeyValueParamTemplate = "Integration keyvalue  %v of Object  is nil Please provide it"
	ErrorFailureNoSoapTargetActionAvailable               = "No Target Instent Soap Action Provided | The one provided is not supported yet | Specify Supported ones wss:save,wss:query,wss:read,wss:delete"

	// Integration Deloyment Constants
	IntegrationRootCertsPathTemplate = "certs/%v/integrations"
	// IntegrationRootCertsRelativePathTemplate  = "../../../../certs/%v/integrations"
	// IntegrationRootCertsCanonicalPathTemplate = "../../../../certs/%v/integrations"
	IntegrationClientCertPathTemplate = "certs/%v/integrations/%v/client.crt"
	IntegrationClientKeyPathTemplate  = "certs/%v/integrations/%v/client.key"
	IntegrationServerCertPathTemplate = "certs/%v/integrations/%v/server.crt"
	IntegrationServerKeyPathTemplate  = "certs/%v/integrations/%v/server.key"
	IntegrationCAPathTemplate         = "certs/%v/integrations/%v/ca.key"

	IntegrationCADevPath          = "certs/dev/integrations/sagex3/ca.crt"
	IntegrationCAProdPath         = "certs/dev/integrations/sagex3/ca.crt"
	IntegrationClientCAProdPath   = "certs/dev/integrations/sagex3/ca.crt"
	IntegrationServerCADevPath    = "certs/dev/integrations/sagex3/ca.crt"
	IntegrationServerCAProdPath   = "certs/dev/integrations/sagex3/ca.crt"
	IntegrationClientCertDevPath  = "certs/dev/integrations/sagex3/client.crt"
	IntegrationClientCertProdPath = "certs/prod/integrations/sagex3/client.crt"
	IntegrationClientKeyDevPath   = "certs/dev/integrations/sagex3/client.key"
	IntegrationClientKeyProdPath  = "certs/prod/integrations/sagex3/client.key"
	IntegrationServerCertProdPath = "certs/prod/integrations/sagex3/server.crt"
	IntegrationServerCertDevPath  = "certs/dev/integrations/sagex3/server.crt"
	IntegrationServerKeyProdPath  = "certs/prod/integrations/sagex3/server.key"
	IntegrationServerKeyDevPath   = "certs/dev/integrations/sagex3/server.key"
	IntegrationServerName         = "SageX3 Integration API"

	TimeoutEnvKey           = "SAGEX3TIMEOUT"
	CaCertPathEnvKey        = "SAGEX3CACERTPATH"
	ServerCertPathEnvKey    = "SAGEX3SERVERCERTPATH"
	ClientCertPathEnvKey    = "SAGEX3CLIENTCERTPATH"
	ClientCertKeyPathEnvKey = "SAGEX3CLIENTCERKEYTPATH"

	DevEnvKey                       = "sagex3dev"
	DevEnvValue                     = "dev"
	ProdEnvKey                      = "sagex3prod"
	ProdEnvValue                    = "prod"
	SageX3DefaultPaymentTerm        = "GBCOD"
	SageX3SalesOrderIntentItem      = "item"
	SageX3CustomerItemIntentCompany = "company"
	SageX3CustomerItemIntentAddress = "address"
	SageX3CustomerItemIntentContact = "contact"
	SOAP_CLIENT_TIMEOUT             = 6000000

	Wsdl = "http://172.16.100.12:8124/soap-wsdl/syracuse/collaboration/syracuse/CAdxWebServiceXmlCC?wsdl"

	Timeout = 6000000
	// Username = ""
	// password = ""

	//Services Soap WS
	// SalesOrderService    = ""
	// PaymentService       = ""
	// ShipmentService      = ""
	// LocationService      = ""
	// BankAccountService   = ""
	// BankSortCodesService = ""
	// InvoiceService       = ""
	// ReceiptService       = ""
	// CustomerService      = ""

	// Systems
	SageX3CustomerCategoryCodeBCGCOD = "SARAF"
	SageX3CustomerCategoryNameBCGCOD = "Sarafu uza"
	DefaultListSize                  = 2000
	SystemSarafu                     = "sarafu"
	SystemSageX3                     = "sagex3"
	SystemLoginext                   = "loginext"
	SystemSCB                        = "scb"
	SystemNMB                        = "nmb"
	SystemDTB                        = "dtb"
	SystemCRDB                       = "crdb"
	SystemPOS                        = "pos"
	SystemBarcode                    = "barcode"
	SoapContextKey                   = "publicName"
	SoapObjectContextKey             = "objectKey"
	QueryActionList                  = "list"
	QueryActionListWithKey           = "listWithKey"
	QueryActionDetail                = "detail"
	QueryActionRead                  = "read"
	SoapAction                       = "soapAction"
	QueryActionGetDescription        = "getDescription"
	CommandActionSave                = "save"
	CommandActionModify              = "modify"
	CommandActionDelete              = "delete"

	DirectionTo   = "To"
	DirectionFrom = "from"

	GetMethod             = "GET"
	PostMethod            = "POST"
	PatchMethod           = "PATCH"
	PutMethod             = "PUT"
	DeleteMethod          = "DELETE"
	ProtocolTypeRest      = "Rest"
	ProtocolTypeSoap      = "Soap"
	ProtocolTypeGraphQL   = "GraphQL"
	ProtocolTypeWebsocket = "Websocket"
	ProtocolTypeGrpc      = "Grpc"
	SupportedStandands    = "soap|rest|graphql"
	IntegratedBanks       = "scb|crdb|nmb|dtb|banks"
	DefaultStandard       = ProtocolTypeSoap
	CodeLang              = "ENG"
	PoolAlias             = "PILOT" // "WSP"//"PILOT"
	PoolId                = ""
	RequestConfig         = "adxwss.optreturn=XML&adxwss.beutify=true"

	//RequestConfig         = "adxwss.optreturn=JSON&adxwss.beutify=true"

	RequestJSONConfig           = "adxwss.optreturn=JSON&adxwss.beutify=true"
	RequestXMLConfig            = "adxwss.optreturn=XML&adxwss.beutify=true"
	ServiceIsLikelyOfflineError = "Request Failed |Likely your client can not reach the API -You could be behind Proxy or VPN or offline |Please check or contact the API Service Administrator"

	//RequestConfig         = "adxwss.trace.on=off&adxwss.trace.size=16384&adonix.trace.on=off&adonix.trace.level=3&adonix.trace.size=8&adxwss.optreturn=JSON&adxwss.beutify=true"
	//Operation allowd

	ParamKeyActionCode = "actioncode"
	ParamKeyBlocKey    = "blockkey"
	ParamKeyLineKeys   = "linekeys"
	ParamKeyLineXml    = "linexml"
	ParamKeyObjectXml  = "objectxml"
	ParamKeyInputXml   = "inputxml"
	ParamKeyLineSize   = "linesize"
	ParamKeyRange      = "range"
	ParamKeyPagination = "pagination"

	//product API Constants

	//ProductStatus
	productStatusActive        = 1
	productStatusInDevelopment = 2
	productStatusOnShortage    = 3
	productStatusNotRenewed    = 4
	productStatusObsolete      = 5
	productStatusNotUsable     = 6

	//
	productJsonDecodingError         = "productJsonDecodingError"
	productCategoryJsonDecodingError = "productCategoryJsonDecodingError"
	sagex3TypeTable                  = "Table"
	sagex3TypeList                   = "List"
	sagex3TypeModDisplay             = "Display"
	sagex3TypeDecimal                = "Decimal"
	sagex3TypeChar                   = "Char"
	sagex3TypeInteger                = "Integer"
	Sagex3PayloadKey                 = "SageX3Payload"

	//CONTENT TYPES

	JSONContentType = "application/json"
	XMLContentType  = "application/xml"
)
