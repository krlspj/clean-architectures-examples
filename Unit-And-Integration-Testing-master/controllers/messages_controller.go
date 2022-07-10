package controllers

import (
	"efficient-api/domain"
	"efficient-api/services"
	"efficient-api/utils/error_utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//Since we are going for the message id more than we, we extracted this functionality to a function so we can have a DRY code.
func getMessageId(msgIdParam string) (int64, error_utils.MessageErr) {
	msgId, msgErr := strconv.ParseInt(msgIdParam, 10, 64)
	if msgErr != nil {
		return 0, error_utils.NewBadRequestError("message id should be a number")
	}
	return msgId, nil
}

func GetMessage(c *gin.Context) {
	msgId, err := getMessageId(c.Param("message_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	message, getErr := services.MessagesService.GetMessage(msgId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, message)
}

func GetAllMessages(c *gin.Context) {
	messages, getErr := services.MessagesService.GetAllMessages()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, messages)
}

func CreateMessage(c *gin.Context) {
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	msg, err := services.MessagesService.CreateMessage(&message)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, msg)
}

func UpdateMessage(c *gin.Context) {
	msgId, err := getMessageId(c.Param("message_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	var message domain.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	message.Id = msgId
	msg, err := services.MessagesService.UpdateMessage(&message)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func DeleteMessage(c *gin.Context) {
	msgId, err := getMessageId(c.Param("message_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	if err := services.MessagesService.DeleteMessage(msgId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

