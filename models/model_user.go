package models

//type companyType int

/*
const (
	Corporations companyType = iota
	NonProfit
	Cooperative
	SoleProprietorship
)
*/

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type KafkaConfig struct {
	KAFKA_TOPIC  string `env:"KAFKA_TOPIC"`
	KAFKA_BROKER string `env:"KAFKA_BROKER"`
}
