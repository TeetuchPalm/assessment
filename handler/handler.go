package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type Expense struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Amount float64 `json:"amount"`
	Note string `json:"note"`
	Tags []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}

var db *sql.DB

func GetDB(dbsever *sql.DB) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
			db = dbsever
            return next(c)
            
        }
    }
}

func CreateExpensesHandler(c echo.Context) error {

	ex := Expense{}
	err := c.Bind(&ex) //เเปลงให้เป็น byte
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4)  RETURNING id, title, amount, note, tags", ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	err = row.Scan(&ex.ID,&ex.Title,&ex.Amount,&ex.Note,pq.Array(&ex.Tags))
	//err = row.Scan(&ex.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, ex)
}

func GetExpensesHandler(c echo.Context) error {
	
	id := c.Param("id")

	idint,err := strconv.ParseInt(id,10,64)
	intid := int(idint)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statment:" + err.Error()})
	}

	row := stmt.QueryRow(intid)
	ex := Expense{}
	err = row.Scan(&ex.ID,&ex.Title,&ex.Amount,&ex.Note,pq.Array(&ex.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "user not found"})
	case nil:
		return c.JSON(http.StatusOK, ex)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan user:" + err.Error()})
	}
}

func UpdateExpensesHandler(c echo.Context) error {
	ex := Expense{}
	err := c.Bind(&ex)
	id := c.Param("id")
	idint,err := strconv.ParseInt(id,10,64)
	if err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1;")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare statment update" + err.Error()})
	}
	ex.ID = int(idint)
	_, err = stmt.Exec(idint, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	
	switch err {
	case nil:
		return c.JSON(http.StatusOK, ex)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "error execute update" + err.Error()})
	}
}



