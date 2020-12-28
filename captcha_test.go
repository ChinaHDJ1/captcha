package captcha

import (
	"log"
	"os"
	"testing"
)

func TestCaptcha(t *testing.T) {
	builder := NewBuilder(DefaultConfig)
	gif, letters, err := builder.Build()
	if err != nil {
		log.Fatal("[Build Captcha]", err)
	}

	file, err := os.Create("./captcha.gif")
	if err != nil {
		log.Fatal("[Create File]", err)
	}

	defer file.Close()

	log.Println("[Letters]", letters)

	log.Println(file.Write(gif))
}
