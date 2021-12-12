package controller_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	. "test-project-1/controller"
	"test-project-1/helper"
	"test-project-1/lib/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func userSetup(db *sql.DB) error {
	_, err := db.Exec("truncate users")
	if err != nil {
		return err
	}
	return nil
}

func insertDataUser(db *sql.DB) {
	_, err := db.Exec("insert into users values (?,?,?,?,?,?,?,?)",
		1,
		"darien kentanu",
		"darienkentanu@gmail.com",
		"$2a$14$T5/nZjF5gKE6dWy/0QwkneKSB82cEknPfz86RrJrt.S70wvuXVtIO",
		nil,
		"2021-12-12 05:21:50.118",
		"2021-12-12 05:21:50.118",
		nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestRegisterUser(t *testing.T) {
	var testCases = []struct {
		name       string
		path       string
		expectCode int
		response   string
		reqBody    M
	}{
		{
			name:       "RegisterUser",
			path:       "/register",
			expectCode: http.StatusCreated,
			response:   "success",
			reqBody: M{"name": "darien kentanu",
				"email":    "darienkentanu@gmail.com",
				"password": "password",
			},
		},
	}

	e, db := helper.InitEcho()

	userSetup(db)
	um := database.NewUserModel(db)
	uc := NewUserController(um)

	for _, testCase := range testCases {
		register, err := json.Marshal(testCase.reqBody)
		if err != nil {
			t.Error(err)
		}

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(register))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			if assert.NoError(t, uc.RegisterUser(c)) {
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
		})
	}
}
