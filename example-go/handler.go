package handler

import (
	"log"

	"github.com/Kong/go-pdk"
)

type Config struct {
	HeaderName string
}

func New() interface{} {
	return &Config{}
}

func (conf *Config) Access(kong *pdk.PDK) {
	headerValue, err := kong.Request.GetHeader("X-Example")
	if err != nil {
		log.Println("Error getting header:", err)
		return
	}

	kong.Response.SetHeader("X-Example-Response", headerValue)
}
