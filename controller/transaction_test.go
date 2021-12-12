package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"test-project-1/constants"
	. "test-project-1/controller"
	"test-project-1/helper"
	"test-project-1/lib/database"
	"test-project-1/model"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TransactionSetup(db *sql.DB) {
	_, err := db.Exec(
		"truncate transaction_details",
	)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		"delete from transactions",
	)
	if err != nil {
		panic(err)
	}

}

func insertItem(db *sql.DB) {
	_, err := db.Exec(
		"delete from items",
	)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(
		"insert into items(id,name,price,cost,created_at,updated_at) values(1, 'pulsa 10rb', 10000, 2000,curdate(),curdate()), (2, 'pulsa 20rb', 20000, 2000,curdate(),curdate())",
	)
	if err != nil {
		panic(err)
	}
}

func login(e *echo.Echo, db *sql.DB) (token string) {
	reqBody := M{
		"email":    "darienkentanu@gmail.com",
		"password": "password",
	}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/login")
	userSetup(db)
	insertDataUser(db)
	um := database.NewUserModel(db)
	uc := NewUserController(um)
	err = uc.LoginUser(c)
	if err != nil {
		panic(err)
	}
	body := rec.Body.String()
	var response = struct {
		Status string     `json:"status"`
		Data   model.User `json:"data"`
	}{}
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		panic(err)
	}
	token = response.Data.Token
	return token
}

func TestNewTransaction(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		response   interface{}
		reqBody    M
		condition  func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool
	}{
		{
			name:       "newTransaction",
			path:       "/newtransaction",
			expectCode: 201,
			response:   "success",
			reqBody: M{
				"item_input": []M{
					{
						"id":       1,
						"quantity": 10,
					},
					{
						"id":       2,
						"quantity": 25,
					},
				},
			},
			condition: assert.NoError,
		},
		{
			name:       "newTransaction",
			path:       "/newtransaction",
			expectCode: 200,
			response:   "",
			reqBody: M{
				"item_input": []M{
					{
						"id":       10,
						"quantity": 10,
					},
					{
						"id":       20,
						"quantity": 25,
					},
				},
			},
			condition: assert.Error,
		},
	}

	e, db := helper.InitEcho()
	token := login(e, db)

	for _, testCase := range testCases {
		TransactionSetup(db)
		insertItem(db)
		tm := database.NewTransactionModel(db)
		im := database.NewItemModel(db)
		tc := NewTransactionController(tm, im)
		payload, err := json.Marshal(testCase.reqBody)
		if err != nil {
			t.Error(err)
		}
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		// t.Run(testCase.name, func(t *testing.T) {
		if testCase.condition(t, middleware.JWT([]byte(constants.JWT_SECRET))(tc.NewTransaction)(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			body := rec.Body.String()

			var response = struct {
				Status string `json:"status"`
				Data   M      `json:"data"`
			}{}
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, testCase.response, response.Status)
		}
		// })

	}
}
