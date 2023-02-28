package models

type EnvVariables struct {
	MongoDbUrl string `env:"MONGO_DB_URL,required"`
}
