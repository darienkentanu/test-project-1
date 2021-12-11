package controller

import (
	"net/http"
	"strconv"
	"test-project-1/lib/database"
	"test-project-1/model"
	"time"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	ItemModel database.ItemModel
}

func NewItemController(ItemModel database.ItemModel) *ItemController {
	return &ItemController{ItemModel: ItemModel}
}

func (ic *ItemController) AddItem(c echo.Context) error {
	var item = model.Item{}
	c.Bind(&item)
	item.Created_at = time.Now()
	item.Updated_at = time.Now()
	insertedItem, err := ic.ItemModel.AddItem(item)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "an error occured")
	}
	return c.JSON(http.StatusAccepted, M{
		"status": "success",
		"data":   insertedItem,
	})
}

func (ic *ItemController) GetItems(c echo.Context) error {
	items, err := ic.ItemModel.GetAllItems()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, M{
		"status": "success",
		"data":   items,
	})
}

func (ic *ItemController) DeleteItem(c echo.Context) error {
	ids := c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please input a valid item id")
	}
	status, err := ic.ItemModel.DeleteItem(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, status)
}

func (ic *ItemController) EditItem(c echo.Context) error {
	var ids = c.Param("id")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please input a valid item id")
	}
	var item = model.Item{}
	c.Bind(&item)
	item.Updated_at = time.Now()
	editedItem, err := ic.ItemModel.EditItem(item, id)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, "an error occured")
	}
	return c.JSON(http.StatusAccepted, M{
		"status": "success",
		"data":   editedItem,
	})
}
