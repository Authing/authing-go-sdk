package model

import "time"

type Tenant struct {
	ID                           string      `json:"id"`
	CreatedAt                    time.Time   `json:"createdAt"`
	UpdatedAt                    time.Time   `json:"updatedAt"`
	UserPoolID                   string      `json:"userPoolId"`
	Name                         string      `json:"name"`
	Logo                         string      `json:"logo"`
	Description                  interface{} `json:"description"`
	CSS                          interface{} `json:"css"`
	SsoPageCustomizationSettings interface{} `json:"ssoPageCustomizationSettings"`
	DefaultLoginTab              string      `json:"defaultLoginTab"`
	DefaultRegisterTab           string      `json:"defaultRegisterTab"`
	PasswordTabConfig            struct {
		EnabledLoginMethods []string `json:"enabledLoginMethods"`
	} `json:"passwordTabConfig"`
	LoginTabs     []string    `json:"loginTabs"`
	RegisterTabs  []string    `json:"registerTabs"`
	ExtendsFields interface{} `json:"extendsFields"`
}

type GetTenantListResponse struct {
	TotalCount int64    `json:"totalCount"`
	List       []Tenant `json:"list"`
}

type TenantDetails struct {
	Tenant
	Apps []struct {
		QrcodeScanning struct {
			Redirect bool `json:"redirect"`
			Interval int  `json:"interval"`
		} `json:"qrcodeScanning"`
		ID          string      `json:"id"`
		CreatedAt   time.Time   `json:"createdAt"`
		UpdatedAt   time.Time   `json:"updatedAt"`
		UserPoolID  string      `json:"userPoolId"`
		Protocol    string      `json:"protocol"`
		IsOfficial  bool        `json:"isOfficial"`
		IsDeleted   bool        `json:"isDeleted"`
		IsDefault   bool        `json:"isDefault"`
		IsDemo      bool        `json:"isDemo"`
		Name        string      `json:"name"`
		Description interface{} `json:"description"`
		Secret      string      `json:"secret"`
		Identifier  string      `json:"identifier"`
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
		InitLoginURL                 interface{}   `json:"initLoginUrl"`
		OidcProviderEnabled          bool          `json:"oidcProviderEnabled"`
		OauthProviderEnabled         bool          `json:"oauthProviderEnabled"`
		SamlProviderEnabled          bool          `json:"samlProviderEnabled"`
		CasProviderEnabled           bool          `json:"casProviderEnabled"`
		RegisterDisabled             bool          `json:"registerDisabled"`
		LoginTabs                    []string      `json:"loginTabs"`
		PasswordTabConfig            struct {
			EnabledLoginMethods []string `json:"enabledLoginMethods"`
		} `json:"passwordTabConfig"`
		DefaultLoginTab      string        `json:"defaultLoginTab"`
		RegisterTabs         []string      `json:"registerTabs"`
		DefaultRegisterTab   string        `json:"defaultRegisterTab"`
		ExtendsFieldsEnabled bool          `json:"extendsFieldsEnabled"`
		ExtendsFields        []interface{} `json:"extendsFields"`
		ComplateFiledsPlace  []interface{} `json:"complateFiledsPlace"`
		SkipComplateFileds   bool          `json:"skipComplateFileds"`
		Ext                  interface{}   `json:"ext"`
		CSS                  interface{}   `json:"css"`
		OidcConfig           struct {
			GrantTypes               []string      `json:"grant_types"`
			ResponseTypes            []string      `json:"response_types"`
			IDTokenSignedResponseAlg string        `json:"id_token_signed_response_alg"`
			TokenEndpointAuthMethod  string        `json:"token_endpoint_auth_method"`
			AuthorizationCodeExpire  int           `json:"authorization_code_expire"`
			IDTokenExpire            int           `json:"id_token_expire"`
			AccessTokenExpire        int           `json:"access_token_expire"`
			RefreshTokenExpire       int           `json:"refresh_token_expire"`
			CasExpire                int           `json:"cas_expire"`
			SkipConsent              bool          `json:"skip_consent"`
			RedirectUris             []string      `json:"redirect_uris"`
			PostLogoutRedirectUris   []interface{} `json:"post_logout_redirect_uris"`
			ClientID                 string        `json:"client_id"`
			ClientSecret             string        `json:"client_secret"`
		} `json:"oidcConfig"`
		OidcJWEConfig interface{} `json:"oidcJWEConfig"`
		SamlConfig    interface{} `json:"samlConfig"`
		OauthConfig   struct {
			ID                              string   `json:"id"`
			ClientSecret                    string   `json:"client_secret"`
			RedirectUris                    []string `json:"redirect_uris"`
			Grants                          []string `json:"grants"`
			AccessTokenLifetime             int      `json:"access_token_lifetime"`
			RefreshTokenLifetime            int      `json:"refresh_token_lifetime"`
			IntrospectionEndpointAuthMethod string   `json:"introspection_endpoint_auth_method"`
			RevocationEndpointAuthMethod    string   `json:"revocation_endpoint_auth_method"`
		} `json:"oauthConfig"`
		CasConfig                   interface{} `json:"casConfig"`
		ShowAuthorizationPage       bool        `json:"showAuthorizationPage"`
		EnableSubAccount            bool        `json:"enableSubAccount"`
		EnableDeviceMutualExclusion bool        `json:"enableDeviceMutualExclusion"`
		LoginRequireEmailVerified   bool        `json:"loginRequireEmailVerified"`
		AgreementEnabled            bool        `json:"agreementEnabled"`
		IsIntegrate                 bool        `json:"isIntegrate"`
		SsoEnabled                  bool        `json:"ssoEnabled"`
		Template                    interface{} `json:"template"`
		SkipMfa                     bool        `json:"skipMfa"`
		CasExpireBaseBrowser        bool        `json:"casExpireBaseBrowser"`
		AppType                     string      `json:"appType"`
		PermissionStrategy          struct {
			Enabled         bool        `json:"enabled"`
			DefaultStrategy string      `json:"defaultStrategy"`
			AllowPolicyID   interface{} `json:"allowPolicyId"`
			DenyPolicyID    interface{} `json:"denyPolicyId"`
		} `json:"permissionStrategy"`
	} `json:"apps"`
}

type CreateTenantRequest struct {
	Name        string `json:"name"`
	AppIds      string `json:"appIds"`
	Logo        string `json:"logo,omitempty"`
	Description string `json:"description,omitempty"`
}

type TenantSsoPageCustomizationSettings struct {
	AutoRegisterThenLogin bool `json:"autoRegisterThenLogin,omitempty"`
	HideForgetPassword    bool `json:"hideForgetPassword,omitempty"`
	HideIdp               bool `json:"hideIdp,omitempty"`
	HideSocialLogin       bool `json:"hideSocialLogin,omitempty"`
}

type ConfigTenantRequest struct {
	CSS                          string                              `json:"css,omitempty"`
	SsoPageCustomizationSettings *TenantSsoPageCustomizationSettings `json:"ssoPageCustomizationSettings,omitempty"`
}

type TenantMembersResponse struct {
	ListTotal int64 `json:"listTotal"`
	List      []struct {
		ID       string `json:"id"`
		TenantID string `json:"tenantId"`
		User     *User  `json:"user"`
	} `json:"list"`
}

type AddTenantMembersResponse struct {
	Tenant
	Users *[]User `json:"users"`
}

type ListExtIdpResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	TenantID    string `json:"tenantId"`
	Connections []struct {
		ID          string `json:"id"`
		Type        string `json:"type"`
		Identifier  string `json:"identifier"`
		DisplayName string `json:"displayName"`
		Logo        string `json:"logo"`
		Enabled     bool   `json:"enabled"`
	} `json:"connections"`
}

type ExtIdpDetailResponse struct {
	ID          string                    `json:"id"`
	Name        string                    `json:"name"`
	Type        string                    `json:"type"`
	Connections []ExtIdpConnectionDetails `json:"connections"`
}

type ExtIdpConnection struct {
	Type            string      `json:"type"`
	Identifier      string      `json:"identifier"`
	DisplayName     string      `json:"displayName"`
	Fields          interface{} `json:"fields"`
	Logo            string      `json:"logo,omitempty"`
	UserMatchFields []string    `json:"userMatchFields,omitempty"`
}

type ExtIdpConnectionDetails struct {
	ID string `json:"id"`
	ExtIdpConnection
}

type CreateExtIdpRequest struct {
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	TenantUd    string             `json:"tenantUd"`
	Connections []ExtIdpConnection `json:"connections"`
}

type UpdateExtIdpRequest struct {
	Name string `json:"name"`
}

type CreateExtIdpConnectionRequest struct {
	ExtIdpId        string      `json:"extIdpId"`
	Type            string      `json:"type"`
	Identifier      string      `json:"identifier"`
	DisplayName     string      `json:"displayName"`
	Fields          interface{} `json:"fields"`
	Logo            string      `json:"logo,omitempty"`
	UserMatchFields []string    `json:"userMatchFields,omitempty"`
}

type UpdateExtIdpConnectionRequest struct {
	DisplayName     string      `json:"displayName"`
	Fields          interface{} `json:"fields"`
	Logo            string      `json:"logo,omitempty"`
	UserMatchFields []string    `json:"userMatchFields,omitempty"`
}

type ChangeExtIdpConnectionStateRequest struct {
	AppID    string `json:"appId,omitempty"`
	TenantID string `json:"tenantId,omitempty"`
	Enabled  bool   `json:"enabled"`
}
