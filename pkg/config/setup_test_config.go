package config

func SetupTestConfig() *Configuration {
	configFile := GetConfigFullFileName("config.test/settings.json")
	config := NewJSONConfigurator(&configFile)

	return config.GetConfiguration()
}
