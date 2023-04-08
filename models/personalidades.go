package models

type Personalidade struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
	Ativo    string `json:"ativo"`
}

var Personalidades []Personalidade
