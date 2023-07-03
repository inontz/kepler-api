package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	qrcode "github.com/skip2/go-qrcode"
)

var Bot *linebot.Client
var Err error

func init() {
	Bot, Err = linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)

}
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "Welcome to Kepler. 🚀",
	})
}

func QRGen(c *gin.Context) {
	currentTime := time.Now()
	tokenID, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
	}
	url := "https://kepler.inontz.xyz/profile/" + tokenID.String()
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("Error encoding")
	}
	data := map[string]interface{}{
		"name":      "KEPLER TESTING",
		"date":      currentTime.Format("2006.01.02 15:04:05"),
		"timestamp": time.Unix(1e9, 0).UTC(),
		"token":     tokenID.String(),
		"qrcode":    png,
	}
	// filename := "./resource/img/qrcode/" + tokenID.String() + ".png"

	// err = qrcode.WriteFile(url, qrcode.Medium, 256, filename)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "🚀 Welcome to Kepler QR. 🚀",
		"data":    data,
	})
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
