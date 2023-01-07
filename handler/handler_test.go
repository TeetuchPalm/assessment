package handler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
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
	expectJSON := "{\"id\":0,\"title\":\"babo\",\"amount\":27,\"note\":\"asd\",\"tags\":[\"food\",\"beverage\"]}\n"
	
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

func TestGetexpensesHandler(t *testing.T) {
	e := echo.New()
	stringForQuery := "{\"food\",\"beverage\"}"
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	expectJSON := "{\"id\":10,\"title\":\"babo\",\"amount\":27,\"note\":\"asd\",\"tags\":[\"food\",\"beverage\"]}\n"
	
	
	row := sqlmock.NewRows([]string{"ID","Title","Amount","Note","Tags"}).AddRow(10, "babo", float64(27), "asd", stringForQuery)
	
	// ถ้าเรา Prepare มา เราต้องใช้ ExpectPrepare ก่อนจะเป็น ExpectQuery 
	mock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").ExpectQuery().WithArgs(10).WillReturnRows(row)

	req := httptest.NewRequest(http.MethodPost, uri("expenses", strconv.Itoa(10)), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
    c.SetParamValues("10")
   
	if assert.NoError(t, GetExpensesHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectJSON, rec.Body.String())
	}
   
   }

func TestUpdateexpensesHandler(t *testing.T) {
	e := echo.New()
	tagarray := []string{"food", "beverage"}
	ex := Expense{Title: "babo", Amount: 27,Note: "asd",Tags : tagarray}
	exJSON := `{"Title":"babo","Amount":27,"Note":"asd","Tags": ["food", "beverage"]}`
	//stringForQuery := "{\"food\",\"beverage\"}"
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	expectJSON := "{\"id\":10,\"title\":\"babo\",\"amount\":27,\"note\":\"asd\",\"tags\":[\"food\",\"beverage\"]}\n"
	
	//row := sqlmock.NewRows([]string{"ID","Title","Amount","Note","Tags"}).AddRow(10, "babo", float64(27), "asd", stringForQuery)
	mockrow :=  sqlmock.NewResult(1,1)
	
	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1;")).ExpectExec().WithArgs(10, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags)).WillReturnResult(mockrow)

	req := httptest.NewRequest(http.MethodPut, uri("expenses", strconv.Itoa(10)), strings.NewReader(exJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
    c.SetParamValues("10")
   
	if assert.NoError(t, UpdateExpensesHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expectJSON, rec.Body.String())
	}
   
   }
   func TestGetAllexpensesHandler(t *testing.T) {
	e := echo.New()
	stringForQuery := "{\"food\",\"beverage\"}"
	stringForQuery2 := "{\"food2\",\"beverage2\"}"
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	expectJSON := "[{\"id\":1,\"title\":\"babo\",\"amount\":27,\"note\":\"asd\",\"tags\":[\"food\",\"beverage\"]},{\"id\":2,\"title\":\"babo2\",\"amount\":28,\"note\":\"asd2\",\"tags\":[\"food2\",\"beverage2\"]}]\n"
	
	
	row := sqlmock.NewRows([]string{"ID","Title","Amount","Note","Tags"}).AddRow(1, "babo", float64(27), "asd", stringForQuery).AddRow(2, "babo2", float64(28), "asd2", stringForQuery2)
	
	// ถ้าเรา Prepare มา เราต้องใช้ ExpectPrepare ก่อนจะเป็น ExpectQuery 
	mock.ExpectPrepare("SELECT id, title, amount, note, tags FROM expenses").ExpectQuery().WithArgs().WillReturnRows(row)

	req := httptest.NewRequest(http.MethodPost, uri("expenses"), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
   
	if assert.NoError(t, GetAllExpensesHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
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
