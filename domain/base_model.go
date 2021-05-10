package domain

import (
	//"context"
	"database/sql"
	"fmt"
	"time"
)

type BaseModel struct {
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt 	sql.NullTime
}

type BaseModelInterface interface {
	GetName() string
}

func (b BaseModel)TotalRow(i BaseModelInterface, condition string, args ...interface{}) (int64, error) {
	//db := config.DBEngine.Conn
	var result int64
	var errorQuery error
	queryString := fmt.Sprintf("SELECT COUNT(*) AS total_row FROM %s", i.GetName())
	if condition != "" {
		queryString += fmt.Sprintf(" WHERE %s", condition)
		//errorQuery = db.QueryRow(context.Background(), queryString, args...).Scan(&result)
	}else{
		//errorQuery = db.QueryRow(context.Background(), queryString, args...).Scan(&result)
	}
	return result, errorQuery
}
