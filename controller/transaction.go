package controller

import (
	"fmt"
	"log"
	"net/http"
	"test-project-1/helper"
	"test-project-1/lib/database"
	"test-project-1/model"
	"time"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	TransactionModel database.TransactionModel
	ItemModel        database.ItemModel
}

func NewTransactionController(TransactionModel database.TransactionModel, ItemModel database.ItemModel) *TransactionController {
	return &TransactionController{
		TransactionModel: TransactionModel,
		ItemModel:        ItemModel,
	}
}

func (tc *TransactionController) NewTransaction(c echo.Context) error {
	var inputJson model.Transaction_Input
	c.Bind(&inputJson)
	var detailItems []model.Transaction_Detail
	var totalPrice float32
	var totalCost float32
	for _, item := range inputJson.Items {
		var tDetail model.Transaction_Detail
		id := item.Id
		i, err := tc.ItemModel.GetItemById(id)
		if err != nil {
			log.Println(err)
			echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
		}

		exist, err := tc.ItemModel.CheckItemID(item.Id)
		if err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "item id not exist")
		}
		if !exist {
			log.Println("item id not exist")
			return echo.NewHTTPError(http.StatusBadRequest, "item id not exist")
		}
		tDetail.ItemID = item.Id
		tDetail.ItemPrice = i.Price
		totalPrice += i.Price * float32(item.Quantity)
		tDetail.ItemCost = i.Cost
		totalCost += i.Cost * float32(item.Quantity)
		tDetail.ItemQuantity = item.Quantity
		tDetail.Created_at = time.Now()
		tDetail.Updated_at = time.Now()
		detailItems = append(detailItems, tDetail)
	}
	t := model.Transaction{}
	t.Number = helper.CreateTransactionNumber()
	fmt.Println(t.Number)
	t.PriceTotal = totalPrice
	t.CostTotal = totalCost
	t.Created_at = time.Now()
	t.Updated_at = time.Now()
	id, err := tc.TransactionModel.CreateTransaction(t, detailItems)
	t.Id = id
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	if id == -1 {
		return echo.NewHTTPError(http.StatusInternalServerError, id)
	}
	err = tc.TransactionModel.InsertTransactionID(id)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	items, err := tc.TransactionModel.GetTransactionDetailByTransactionID(id)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "an error has been occured")
	}
	return c.JSON(http.StatusCreated, M{
		"status":      "success",
		"Transaction": t,
		"Detail":      items,
	})
}
