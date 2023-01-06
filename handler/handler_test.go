package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateExpensesHandler(t *testing.T) {
	e := echo.New()
	tagarray := []string{"food", "beverage"}
	stringForQuery := "{\"food\",\"beverage\"}"
	ex := Expense{Title: "babo", Amount: 27,Note: "asd",Tags : tagarray}
	exJSON := `{"Title":"babo","Amount":27,"Note":"asd","Tags": ["food", "beverage"]}`
	expectJSON := "{\"ID\":0,\"Title\":\"babo\",\"Amount\":27,\"Note\":\"asd\",\"Tags\":[\"food\",\"beverage\"]}\n"
	
	//err := c.Bind(&ex) //เเปลงให้เป็น byte
	/*if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}*/
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	e.Use(GetDB(db))
	rows := sqlmock.NewRows([]string{"ID","Title","Amount","Note","Tags"}).AddRow(0, "babo", float64(27), "asd", stringForQuery)
	//row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id", ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	//err = rows.Scan(&ex.ID)
	mock.ExpectQuery("INSERT INTO expenses").WithArgs(ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags)).WillReturnRows(rows)
	/*if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}*/
	req := httptest.NewRequest(http.MethodPost, uri("expenses"), strings.NewReader(exJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	


	if assert.NoError(t, CreateExpensesHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectJSON, rec.Body.String())
	}

}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}
