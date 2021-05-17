package domain

import (
	//"context"
	"database/sql"
	"fmt"
	"github.com/issengi/goboot/app/config"
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
	db := config.DBEngine
	var errorQuery error
	var result int64
	query := fmt.Sprintf(`SELECT COUNT(*) FROM %s`, i.GetName())
	if condition != "" {
		query = fmt.Sprintf(`%s WHERE %s`, query, condition)
	}
	errorQuery = db.QueryRow(query, args...).Scan(&result)
	return result, errorQuery
}
