package constant

const (
	HttpMethodGet  = "GET"
	HttpMethodPost = "POST"
)

const (
	CoreAuthingDefaultUrl  = "https://core.authing.cn"
	CoreAuthingGraphqlPath = "/graphql/v2"
	PublicKey              = "MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC4xKeUgQ+Aoz7TLfAfs9+paePb5KIofVthEopwrXFkp8OCeocaTHt9ICjTT2QeJh6cZaDaArfZ873GPUn00eOIZ7Ae+TiA2BKHbCvloW3w5Lnqm70iSsUi5Fmu9/2+68GZRH9L7Mlh8cFksCicW2Y2W2uMGKl64GDcIq3au+aqJQIDAQAB"

	/**
	 * token 过期时间
	 */
	AccessTokenExpiresAt int64 = 0

	/**
	 * 应用 Id
	 */
	AppId = ""

	//应用密钥
	Secret = ""
	//应用身份协议
	Protocol = "oidc"
	//获取 token 端点认证方式
	TokenEndPointAuthMethod = ClientSecretPost
	//检查 token 端点认证方式
	IntrospectionEndPointAuthMethod = ClientSecretPost
	//撤回 token 端点认证方式
	RevocationEndPointAuthMethod = ClientSecretPost

	//应用回调地址
	RedirectUri = ""
	//Websocket 服务器域名
	WebsocketHost = ""

	SdkType    = "SDK"
	SdkVersion = "go:2.0.0"

	// TokenCacheKeyPrefix token缓存key前缀
	TokenCacheKeyPrefix = "token_"
)

type ProtocolEnum string

const (
	OAUTH ProtocolEnum = "oauth"
	OIDC  ProtocolEnum = "oidc"
	CAS   ProtocolEnum = "cas"
	SAML  ProtocolEnum = "saml"
)

type AuthMethodEnum string

const (
	ClientSecretPost  = "client_secret_post"
	ClientSecretBasic = "client_secret_basic"
	None              = "none"
)
