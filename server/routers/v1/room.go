package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/controllers/v1"
	"server/core/models"
)

var err error

// GetRoom godoc
// @Summary Retrieves room based on given ID
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} models.Room
// @Router /rooms/{id} [get]
// @Security ApiKeyAuth
func (server *Server) GetRoom(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	room, err := v1.GetRoom(server.db, id)
	if err != nil {
		ctx.AbortWithStatus(404)
	}

	ctx.JSON(200, room)
}

// GetRooms godoc
// @Summary Retrieves all rooms
// @Produce json
// @Success 200 {array} models.Room
// @Router /rooms/ [get]
// @Security ApiKeyAuth
func (server *Server) GetRooms(ctx *gin.Context) {
	var args models.Args

	args.Sort = ctx.DefaultQuery("Sort", "ID")
	args.Order = ctx.DefaultQuery("Order", "DESC")
	args.Limit = ctx.DefaultQuery("Limit", "25")
	args.Search = ctx.DefaultQuery("Search", "")

	rooms, filteredData, totalData, err := v1.GetRooms(server.db, args)
	if err != nil {
		ctx.AbortWithStatus(404)
	}

	data := models.Response{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         rooms,
	}

	ctx.JSON(200, data)
}

// CreateRoom godoc
// @Summary Creates room
// @Produce json
// @Param room body models.Room true "Add Room"
// @Success 200 {object} models.Room
// @Router /rooms/create [post]
// @Security ApiKeyAuth
func (server *Server) CreateRoom(ctx *gin.Context) {
	room := new(models.Room)
	ctx.MustBindWith(&room, binding.Form)

	room, err := v1.SaveRoom(server.db, room)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, room)
}

// UpdateRoom godoc
// @Summary Updates room
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} models.Room
// @Router /rooms/update{id} [put]
// @Security ApiKeyAuth
func (server *Server) UpdateRoom(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	room, err := v1.GetRoom(server.db, id)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.MustBindWith(&room, binding.Form)

	room, err = v1.SaveRoom(server.db, room)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, room)
}

// DeleteRoom godoc
// @Summary Deletes room
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} models.Room
// @Router /rooms/delete{id} [delete]
// @Security ApiKeyAuth
func (server *Server) DeleteRoom(ctx *gin.Context) {
	id := ctx.PostForm("id")

	err = v1.DeleteRoom(server.db, id)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, gin.H{"id#" + id: "deleted"})
}
