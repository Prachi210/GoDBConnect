package Config

import "github.com/spf13/viper"

func LoadConfig() (config ConfigModel, err error) {
	path := "C:/Users/practhak/go/src/github.com/Prachi210/GoDBConnect"
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	//viper.AutomaticEnv()

	err = viper.ReadInConfig()
	err = viper.Unmarshal(&config)
	config.CountPerRoutine = 3
	if err != nil {
		return
	}
	return
}
