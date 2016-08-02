package dialect

import (
	"fmt"
	"strings"
	"time"
)

const (
	//With timezone support
	postgresTimeFormat = "2006-01-02 15:04:05.999999999-07:00"
)

type postgreSQL struct{}

func (d postgreSQL) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d postgreSQL) EncodeString(s string) string {
	// http://www.postgresql.org/docs/9.2/static/sql-syntax-lexical.html
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d postgreSQL) EncodeBool(b bool) string {
	if b {
		return "TRUE"
	}
	return "FALSE"
}

func (d postgreSQL) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(postgresTimeFormat) + `'`
}

func (d postgreSQL) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`E'\\x%x'`, b)
}

func (d postgreSQL) Placeholder() string {
	return "?"
}
