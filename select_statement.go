// Package pgsql helps build SQL queries.
package pgsql

import (
	"strings"
)

type SelectStatement struct {
	Args          Args
	SelectClause  SelectClause
	FromClause    FromClause
	WhereClause   WhereClause
	OrderByClause OrderByClause
	LimitClause   LimitClause
	OffsetClause  OffsetClause
}

func (s *SelectStatement) String() string {
	sb := &strings.Builder{}

	writeCount := 0
	f := func(s string) {
		if s == "" {
			return
		}

		if writeCount > 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(s)
		writeCount += 1
	}

	f(s.SelectClause.String())
	f(s.FromClause.String())
	f(s.WhereClause.String())
	f(s.OrderByClause.String())
	f(s.LimitClause.String())
	f(s.OffsetClause.String())

	return sb.String()
}

func (s *SelectStatement) Distinct() *SelectStatement {
	s.SelectClause.Distinct()
	return s
}

func (s *SelectStatement) DistinctOn(sql string, values ...interface{}) *SelectStatement {
	s.SelectClause.DistinctOn(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) Select(sql string, values ...interface{}) *SelectStatement {
	s.SelectClause.Select(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) From(sql string, values ...interface{}) *SelectStatement {
	s.FromClause.From(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) Where(sql string, values ...interface{}) *SelectStatement {
	s.WhereClause.Where(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) WhereOr(sql string, values ...interface{}) *SelectStatement {
	s.WhereClause.Or(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) OrderBy(sql string, values ...interface{}) *SelectStatement {
	s.OrderByClause.OrderBy(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) Limit(sql string, values ...interface{}) *SelectStatement {
	s.LimitClause.Limit(sql, &s.Args, values...)
	return s
}

func (s *SelectStatement) Offset(sql string, values ...interface{}) *SelectStatement {
	s.OffsetClause.Offset(sql, &s.Args, values...)
	return s
}
