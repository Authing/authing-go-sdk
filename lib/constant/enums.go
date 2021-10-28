package constant

const (
	HttpMethodGet  = "GET"
	HttpMethodPost = "POST"
)

const (
	CoreAuthingDefaultUrl  = "https://core.authing.cn"
	CoreAuthingGraphqlPath = "/graphql/v2"

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
	UserCacheKeyPrefix  = "user_"
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

type ResourceTargetTypeEnum string

const (
	USER  ResourceTargetTypeEnum = "USER"
	ROLE  ResourceTargetTypeEnum = "ROLE"
	GROUP ResourceTargetTypeEnum = "GROUP"
	ORG   ResourceTargetTypeEnum = "ORG"
)

type ApplicationDefaultAccessPolicies string

const (
	AllowAll ApplicationDefaultAccessPolicies = "ALLOW_ALL"
	DenyAll  ApplicationDefaultAccessPolicies = "DENY_ALL"
)

type GetAuthorizedTargetsOpt string

const (
	AND GetAuthorizedTargetsOpt = "AND"
	OR  GetAuthorizedTargetsOpt = "OR"
)

type ProviderTypeEnum string

const (
	DingTalk   ProviderTypeEnum = "dingtalk"
	WechatWork ProviderTypeEnum = "wechatwork"
	AD         ProviderTypeEnum = "ad"
)

type PrincipalAuthenticateType string

const (
	P PrincipalAuthenticateType = "P"
	E PrincipalAuthenticateType = "E"
)

type MfaSource string

const (
	Self        MfaSource = "SELF"
	Application MfaSource = "APPLICATION"
)

type SocialProviderType string

const (
	WECHATPC            SocialProviderType = "wechat:pc"
	GITHUB              SocialProviderType = "github"
	GOOGLE              SocialProviderType = "google"
	QQ                  SocialProviderType = "qq"
	APPLE               SocialProviderType = "apple"
	BAIDU               SocialProviderType = "baidu"
	ALIPAY              SocialProviderType = "alipay"
	LARK_APP_STORE      SocialProviderType = "lark:app-store"
	LARK_CUSTOM_APP     SocialProviderType = "lark:custom-app"
	WEIBO               SocialProviderType = "weibo"
	DINGTALK            SocialProviderType = "dingtalk"
	WECHAT_WEB          SocialProviderType = "wechat:webpage-authorization"
	ALIPAY_MOBILE       SocialProviderType = "alipay"
	WECHAT_MQ_DEFAULT   SocialProviderType = "wechat:miniprogram:default"
	WECHAT_MOBILE       SocialProviderType = "wechat:mobile"
	WECHATWORK_SP_AUTHZ SocialProviderType = "wechatwork:service-provider:authorization"
	WECHATWORK_SP_QR    SocialProviderType = "wechatwork:service-provider:qrconnect"
	WECHATWORK_CORP_QR  SocialProviderType = "wechatwork:corp:qrconnect"
	WECHAT_MP_AL        SocialProviderType = "wechat:miniprogram:app-launch"
	WECHAT_MP_QR        SocialProviderType = "wechat:miniprogram:qrconnect"
)

type GenerateCodeChallengeMethod string

const (
	PLAIN GenerateCodeChallengeMethod = "plain"
	S256  GenerateCodeChallengeMethod = "S256"
)

type TicketFormat string

const (
	XML  TicketFormat = "XML"
	JSON TicketFormat = "JSON"
)
