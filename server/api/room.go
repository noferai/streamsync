package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"server/controller"
	"server/model"
)

var err error

// GetRoom godoc
// @Summary Retrieves room based on given ID
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} model.Room
// @Router /rooms/{id} [get]
// @Security ApiKeyAuth
func (base *Controller) GetRoom(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	room, err := controller.GetRoom(base.DB, id)
	if err != nil {
		ctx.AbortWithStatus(404)
	}

	ctx.JSON(200, room)
}

// GetRooms godoc
// @Summary Retrieves all rooms
// @Produce json
// @Success 200 {array} model.Room
// @Router /rooms/ [get]
// @Security ApiKeyAuth
func (base *Controller) GetRooms(ctx *gin.Context) {
	var args model.Args

	args.Sort = ctx.DefaultQuery("Sort", "ID")
	args.Order = ctx.DefaultQuery("Order", "DESC")
	args.Limit = ctx.DefaultQuery("Limit", "25")
	args.Search = ctx.DefaultQuery("Search", "")

	rooms, filteredData, totalData, err := controller.GetRooms(ctx, base.DB, args)
	if err != nil {
		ctx.AbortWithStatus(404)
	}

	data := model.Response{
		TotalData:    totalData,
		FilteredData: filteredData,
		Data:         rooms,
	}

	ctx.JSON(200, data)
}

// CreateRoom godoc
// @Summary Creates room
// @Produce json
// @Param room body model.Room true "Add Room"
// @Success 200 {object} model.Room
// @Router /rooms/create [post]
// @Security ApiKeyAuth
func (base *Controller) CreateRoom(ctx *gin.Context) {
	room := new(model.Room)
	ctx.MustBindWith(&room, binding.Form)

	room, err := controller.SaveRoom(base.DB, room)
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
// @Success 200 {object} model.Room
// @Router /rooms/update{id} [put]
// @Security ApiKeyAuth
func (base *Controller) UpdateRoom(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	room, err := controller.GetRoom(base.DB, id)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.MustBindWith(&room, binding.Form)

	room, err = controller.SaveRoom(base.DB, room)
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
// @Success 200 {object} model.Room
// @Router /rooms/delete{id} [delete]
// @Security ApiKeyAuth
func (base *Controller) DeleteRoom(ctx *gin.Context) {
	id := ctx.PostForm("id")

	err = controller.DeleteRoom(base.DB, id)
	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}

	ctx.JSON(200, gin.H{"id#" + id: "deleted"})
}
