package fetch

type Config struct {
	Url     string
	Method  string
	Headers ConfigHeaders
	Query   ConfigQuery
	Params  ConfigParams
	Body    ConfigBody
}

type ConfigHeaders map[string]string
type ConfigQuery map[string]string
type ConfigParams map[string]string
type ConfigBody interface{}
