package models

type KafkaEvent struct {
	Name       string
	PolicyData *Policy
}

type AnniversaryStatement struct {
	Policy *Policy
}

type Token struct {
	Token string `json:"token"`
}
