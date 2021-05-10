package services

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/jackc/pgx/v4"
	"log"
	"strings"
)

type Logger struct {
}

func (pl *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	switch level {
	case pgx.LogLevelTrace:
		log.Println(fmt.Sprintf(`[PGX_LOG] %s : %s`, strings.ToUpper(level.String()), msg), data)
	case pgx.LogLevelError:
		log.Println(fmt.Sprintf(`[PGX_LOG] %s : %s`, strings.ToUpper(level.String()), data["err"]))
	}
}

type GoPgLogger struct {

}

func (d GoPgLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}
func (d GoPgLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}