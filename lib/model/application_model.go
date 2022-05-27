package model

import "time"

type Application struct {
	QrcodeScanning struct {
		Redirect bool `json:"redirect"`
		Interval int  `json:"interval"`
	} `json:"qrcodeScanning"`
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserPoolId  string    `json:"userPoolId"`
	Protocol    string    `json:"protocol"`
	IsOfficial  bool      `json:"isOfficial"`
	IsDeleted   bool      `json:"isDeleted"`
	IsDefault   bool      `json:"isDefault"`
	IsDemo      bool      `json:"isDemo"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Secret      string    `json:"secret"`
	Identifier  string    `json:"identifier"`
	Jwks        struct {
		Keys []struct {
			E   string `json:"e"`
			N   string `json:"n"`
			D   string `json:"d"`
			P   string `json:"p"`
			Q   string `json:"q"`
			Dp  string `json:"dp"`
			Dq  string `json:"dq"`
			Qi  string `json:"qi"`
			Kty string `json:"kty"`
			Kid string `json:"kid"`
			Alg string `json:"alg"`
			Use string `json:"use"`
		} `json:"keys"`
	} `json:"jwks"`
	SsoPageCustomizationSettings interface{}   `json:"ssoPageCustomizationSettings"`
	Logo                         string        `json:"logo"`
	RedirectUris                 []string      `json:"redirectUris"`
	LogoutRedirectUris           []interface{} `json:"logoutRedirectUris"`
	OidcProviderEnabled          bool          `json:"oidcProviderEnabled"`
	OauthProviderEnabled         bool          `json:"oauthProviderEnabled"`
	SamlProviderEnabled          bool          `json:"samlProviderEnabled"`
	CasProviderEnabled           bool          `json:"casProviderEnabled"`
	RegisterDisabled             bool          `json:"registerDisabled"`
	LoginTabs                    []string      `json:"loginTabs"`
	PasswordTabConfig            struct {
		EnabledLoginMethods []string `json:"enabledLoginMethods"`
	} `json:"passwordTabConfig"`
	DefaultLoginTab            string        `json:"defaultLoginTab"`
	RegisterTabs               []string      `json:"registerTabs"`
	DefaultRegisterTab         string        `json:"defaultRegisterTab"`
	LdapConnections            interface{}   `json:"ldapConnections"`
	AdConnections              []interface{} `json:"adConnections"`
	DisabledSocialConnections  interface{}   `json:"disabledSocialConnections"`
	DisabledOidcConnections    []interface{} `json:"disabledOidcConnections"`
	DisabledSamlConnections    []interface{} `json:"disabledSamlConnections"`
	DisabledOauth2Connections  []interface{} `json:"disabledOauth2Connections"`
	DisabledCasConnections     []interface{} `json:"disabledCasConnections"`
	DisabledAzureAdConnections []interface{} `json:"disabledAzureAdConnections"`
	ExtendsFieldsEnabled       bool          `json:"extendsFieldsEnabled"`
	ExtendsFields              []interface{} `json:"extendsFields"`
	Ext                        struct {
		DontFinishNotYet bool   `json:"_dontFinishNotYet"`
		AppName          string `json:"_appName"`
		AliyunDomain     string `json:"AliyunDomain"`
		AliyunAccountId  string `json:"AliyunAccountId"`
		SamlConfig       struct {
		} `json:"samlConfig"`
	} `json:"ext"`
	Css        interface{} `json:"css"`
	OidcConfig struct {
		GrantTypes               []string      `json:"grant_types"`
		ResponseTypes            []string      `json:"response_types"`
		IdTokenSignedResponseAlg string        `json:"id_token_signed_response_alg"`
		TokenEndpointAuthMethod  string        `json:"token_endpoint_auth_method"`
		AuthorizationCodeExpire  int           `json:"authorization_code_expire"`
		IdTokenExpire            int           `json:"id_token_expire"`
		AccessTokenExpire        int           `json:"access_token_expire"`
		RefreshTokenExpire       int           `json:"refresh_token_expire"`
		CasExpire                int           `json:"cas_expire"`
		SkipConsent              bool          `json:"skip_consent"`
		RedirectUris             []string      `json:"redirect_uris"`
		PostLogoutRedirectUris   []interface{} `json:"post_logout_redirect_uris"`
		ClientId                 string        `json:"client_id"`
		ClientSecret             string        `json:"client_secret"`
	} `json:"oidcConfig"`
	OidcJWEConfig interface{} `json:"oidcJWEConfig"`
	SamlConfig    struct {
		Acs                                string      `json:"acs"`
		Audience                           string      `json:"audience"`
		Recipient                          string      `json:"recipient"`
		Destination                        string      `json:"destination"`
		Mappings                           interface{} `json:"mappings"`
		DigestAlgorithm                    string      `json:"digestAlgorithm"`
		SignatureAlgorithm                 string      `json:"signatureAlgorithm"`
		AuthnContextClassRef               string      `json:"authnContextClassRef"`
		LifetimeInSeconds                  int         `json:"lifetimeInSeconds"`
		SignResponse                       bool        `json:"signResponse"`
		NameIdentifierFormat               string      `json:"nameIdentifierFormat"`
		SamlRequestSigningCert             string      `json:"samlRequestSigningCert"`
		SamlResponseSigningCert            string      `json:"samlResponseSigningCert"`
		SamlResponseSigningKey             string      `json:"samlResponseSigningKey"`
		SamlResponseSigningCertFingerprint string      `json:"samlResponseSigningCertFingerprint"`
		EmailDomainSubstitution            string      `json:"emailDomainSubstitution"`
	} `json:"samlConfig"`
	OauthConfig                 interface{} `json:"oauthConfig"`
	CasConfig                   interface{} `json:"casConfig"`
	ShowAuthorizationPage       bool        `json:"showAuthorizationPage"`
	EnableSubAccount            bool        `json:"enableSubAccount"`
	EnableDeviceMutualExclusion bool        `json:"enableDeviceMutualExclusion"`
	LoginRequireEmailVerified   bool        `json:"loginRequireEmailVerified"`
	AgreementEnabled            bool        `json:"agreementEnabled"`
	IsIntegrate                 bool        `json:"isIntegrate"`
	SsoEnabled                  bool        `json:"ssoEnabled"`
	Template                    string      `json:"template"`
	SkipMfa                     bool        `json:"skipMfa"`
	CasExpireBaseBrowser        bool        `json:"casExpireBaseBrowser"`
	PermissionStrategy          struct {
		Enabled         bool        `json:"enabled"`
		DefaultStrategy string      `json:"defaultStrategy"`
		AllowPolicyId   interface{} `json:"allowPolicyId"`
		DenyPolicyId    interface{} `json:"denyPolicyId"`
	} `json:"permissionStrategy"`
}

type ApplicationActiveUsers struct {
	ThirdPartyIdentity struct {
		Provider     string `json:"provider"`
		RefreshToken string `json:"refreshToken"`
		AccessToken  string `json:"accessToken"`
		Scope        string `json:"scope"`
		ExpiresIn    string `json:"expiresIn"`
		UpdatedAt    string `json:"updatedAt"`
	} `json:"thirdPartyIdentity"`
	Id                        string      `json:"id"`
	CreatedAt                 time.Time   `json:"createdAt"`
	UpdatedAt                 time.Time   `json:"updatedAt"`
	UserPoolId                string      `json:"userPoolId"`
	IsRoot                    bool        `json:"isRoot"`
	Status                    string      `json:"status"`
	Oauth                     string      `json:"oauth"`
	Email                     string      `json:"email"`
	Phone                     string      `json:"phone"`
	Username                  string      `json:"username"`
	Unionid                   string      `json:"unionid"`
	Openid                    string      `json:"openid"`
	Nickname                  string      `json:"nickname"`
	Company                   string      `json:"company"`
	Photo                     string      `json:"photo"`
	Browser                   string      `json:"browser"`
	Device                    string      `json:"device"`
	Password                  string      `json:"password"`
	Salt                      string      `json:"salt"`
	LoginsCount               int         `json:"loginsCount"`
	LastIp                    string      `json:"lastIp"`
	Name                      string      `json:"name"`
	GivenName                 string      `json:"givenName"`
	FamilyName                string      `json:"familyName"`
	MiddleName                string      `json:"middleName"`
	Profile                   string      `json:"profile"`
	PreferredUsername         string      `json:"preferredUsername"`
	Website                   string      `json:"website"`
	Gender                    string      `json:"gender"`
	Birthdate                 string      `json:"birthdate"`
	Zoneinfo                  string      `json:"zoneinfo"`
	Locale                    string      `json:"locale"`
	Address                   string      `json:"address"`
	Formatted                 string      `json:"formatted"`
	StreetAddress             string      `json:"streetAddress"`
	Locality                  string      `json:"locality"`
	Region                    string      `json:"region"`
	PostalCode                string      `json:"postalCode"`
	City                      string      `json:"city"`
	Province                  string      `json:"province"`
	Country                   string      `json:"country"`
	RegisterSource            []string    `json:"registerSource"`
	SecretInfo                interface{} `json:"secretInfo"`
	EmailVerified             bool        `json:"emailVerified"`
	PhoneVerified             bool        `json:"phoneVerified"`
	LastLogin                 time.Time   `json:"lastLogin"`
	Blocked                   bool        `json:"blocked"`
	IsDeleted                 bool        `json:"isDeleted"`
	SendSmsCount              int         `json:"sendSmsCount"`
	SendSmsLimitCount         int         `json:"sendSmsLimitCount"`
	DataVersion               string      `json:"dataVersion"`
	EncryptedPassword         string      `json:"encryptedPassword"`
	SignedUp                  time.Time   `json:"signedUp"`
	ExternalId                string      `json:"externalId"`
	MainDepartmentId          string      `json:"mainDepartmentId"`
	MainDepartmentCode        string      `json:"mainDepartmentCode"`
	LastMfaTime               string      `json:"lastMfaTime"`
	PasswordSecurityLevel     int         `json:"passwordSecurityLevel"`
	ResetPasswordOnFirstLogin bool        `json:"resetPasswordOnFirstLogin"`
	SyncExtInfo               interface{} `json:"syncExtInfo"`
	PhoneCountryCode          string      `json:"phoneCountryCode"`
	Source                    interface{} `json:"source"`
	LastIP                    string      `json:"lastIP"`
	Token                     string      `json:"token"`
	TokenExpiredAt            time.Time   `json:"tokenExpiredAt"`
}

type ApplicationAgreement struct {
	UserPoolId string `json:"userPoolId"`
	AppId      string `json:"appId"`
	Title      string `json:"title"`
	Lang       string `json:"lang"`
	Required   bool   `json:"required"`
	Order      int    `json:"order"`
	Id         int    `json:"id"`
}

type ApplicationTenantDetails struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Logo        string    `json:"logo"`
	Domain      string    `json:"domain"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Protocol    string    `json:"protocol"`
	IsIntegrate bool      `json:"isIntegrate"`
	Tenants     []Tenant  `json:"tenants"`
}
