package constants

const (
	ApiName                       string = "eventstore"
	ApiEventSourceName            string = "eventstore"
	ApiEventStoreAPIHealthUrl            = "http://localhost:2000/api/v1/eventstore/health"
	ApiAppName                    string = "app"
	ApiEventName                  string = "event"
	ApiVersionedRouteTemplate     string = "/api/v%v/%v/"
	ApiBaseURLLocal               string = "https://duxtevfd.duxte.net:2433/api/v1"
	ApiBaseURLRemote              string = "https://duxtevfd.duxte.net:2433/api/v1"
	ApiBaseURLUat                 string = "https://duxtevfd.duxte.net:2433/api/v1"
	ApiBaseURL                    string = "https://duxtevfd.duxte.net:2433/api/v1"
	ApiSecret                     string = "jdnfksdmfksd"        //ACCESS_SECRET
	ApiRefreshSecret              string = "mcmvmkmsdnfsdmfdsjf" // put o  env variable  //HAKIKIGANJANI_REFRESH_TOKEN
	ApiConfigFolderPath           string = "/etc/integrations/eventstore/conf/"
	ApiTargetHostRemote           string = "remote"
	ApiTargetHostLocal            string = "local"
	ApiAPIAllowCredentials        bool   = false
	ApiMaxLoginCount              int    = 5
	ApiUsernameSeparatorCharacter string = ","
	ApiDefaulPort                        = 1025
	ApiEventStorePort                    = ApiDefaulPort
	ApiAppDefaultPort                    = 1026
	ApiWebAppDefaultPort                 = 1027
	ApiEventDefaultPort                  = 1028
	ApiPayloadDefaultPort                = 1029
	ApiSubscriptionDefaultPort           = 1030
	ApiUserDefaultPort                   = 1031
	ApiProfileDefaultPort                = 1032
	ApiSettingDefaultPort                = 1033
)
