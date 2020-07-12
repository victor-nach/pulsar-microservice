package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Res - Response struct
type Res struct {
	Ctx    *gin.Context `json:"-"`
	Err    error        `json:"-"`
	ErrMsg string		`json:"error,omitempty"`
	Msg    string       `json:"message,omitempty"`
	Status int          `json:"status,omitempty"`
	Data   interface{}  `json:"data,omitempty"`
}

// ResErr - sends a json response if there is an error
func ResErr(r Res) {
	if r.ErrMsg == "" {
		r.ErrMsg = "Error occured"
	}
	if r.Status == 0 {
		r.Status = 400
	}
	log.Println("errory: ", r.Err)
	log.Println("stru: ", r)
	r.Ctx.JSON(r.Status, r)
}

// ResSuccess - sends a json respnse on success
func ResSuccess(r Res) {
	if r.Msg == "" {
		r.Msg = "Successful!"
	}
	if r.Status == 0 {
		r.Status = 200
	}
	r.Ctx.JSON(r.Status, r)
}
