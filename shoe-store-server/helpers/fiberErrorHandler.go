package helpers

import (
	"github.com/gofiber/fiber/v2/log"
)

func HandleError(msg string, err error, quit bool) {
	if err != nil {
		if quit {
			log.Fatalf("%s: %s", msg, err)
		} else {
			log.Errorf("%s: %s", msg, err)
		}
	}
}
