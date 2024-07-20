package config

func NewConfiguration(configFile *string) *Configuration {
	return NewJSONConfigurator(configFile).GetConfiguration()
}
