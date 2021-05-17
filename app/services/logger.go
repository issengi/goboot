package services

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type GoPgLogger struct {

}

func (d GoPgLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}
func (d GoPgLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.Query)
	return nil
}