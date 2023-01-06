// go:build integration
package handler

import (
	"encoding/json"
	_ "fmt"
	"io"
	_ "log"
	"net/http"
	_"net/http/httptest"
	_"strings"
	"testing"
	"bytes"

	_"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	body := bytes.NewBufferString(`{
		"Title": "strawberry smoothie",
		"Amount": 79,
		"Note": "night market promotion discount 10 bath", 
		"Tags": ["food", "beverage"]
}`)
var ex Expense
	//e := echo.New()
	/*e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	//e.Use(handler.GetDB(db))

	e.POST("/expenses", CreateExpensesHandler)

	log.Fatal(e.Start(":2565"))*/
	
	/*for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("localhost:%d", serverPort), 30*time.Second)
		if err != nil {
			log.Println(err)
		}
		if conn != nil {
			conn.Close()
			break
		}
	}*/
	// Arrange

	//exJSON := `{"Title":"nut","Amount":27,"Note":"ass","Tags": ["zzzz", "ssss"]}`
	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&ex)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, ex.ID)
	assert.Equal(t, "strawberry smoothie", ex.Title)
	assert.Equal(t, float64(79), ex.Amount)
	assert.Equal(t, "night market promotion discount 10 bath", ex.Note)
	assert.Equal(t, []string{"food", "beverage"} , ex.Tags)
	

	//expectJSON := "{\"ID\":7,\"Title\":\"nut\",\"Amount\":27,\"Note\":\"ass\",\"Tags\":[\"zzzz\",\"ssss\"]}\n"


	/*if assert.NoError(t, CreateExpensesHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectJSON, rec.Body.String())
	}*/


	
	// Assertions
	/*if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, expectJSON, string(byteBody))
	}*/

	/*ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = eh.Shutdown(ctx)
	assert.NoError(t, err)*/

	
}
func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	//AuthToken := "Basic cmVzaXN0ZWR6OjY5Njk="
	//req.Header.Add("Authorization", AuthToken)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}
