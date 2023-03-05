package handlers

import (
	"fmt"
	"os"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/gofiber/uuid"

)


func RoomCreate(c *fibre.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s", guuid.New().String()))

}