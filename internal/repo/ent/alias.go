// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/drrrMikado/shorten/internal/repo/ent/alias"
)

// Alias is the model entity for the Alias schema.
type Alias struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// URL holds the value of the "url" field.
	URL string `json:"url,omitempty"`
	// Pv holds the value of the "pv" field.
	Pv uint64 `json:"pv,omitempty"`
	// Expire holds the value of the "expire" field.
	Expire time.Time `json:"expire,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Alias) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case alias.FieldID, alias.FieldPv:
			values[i] = new(sql.NullInt64)
		case alias.FieldKey, alias.FieldURL:
			values[i] = new(sql.NullString)
		case alias.FieldCreateTime, alias.FieldUpdateTime, alias.FieldExpire:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Alias", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Alias fields.
func (a *Alias) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case alias.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case alias.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				a.CreateTime = value.Time
			}
		case alias.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				a.UpdateTime = value.Time
			}
		case alias.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				a.Key = value.String
			}
		case alias.FieldURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field url", values[i])
			} else if value.Valid {
				a.URL = value.String
			}
		case alias.FieldPv:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pv", values[i])
			} else if value.Valid {
				a.Pv = uint64(value.Int64)
			}
		case alias.FieldExpire:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire", values[i])
			} else if value.Valid {
				a.Expire = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Alias.
// Note that you need to call Alias.Unwrap() before calling this method if this Alias
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Alias) Update() *AliasUpdateOne {
	return (&AliasClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Alias entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Alias) Unwrap() *Alias {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Alias is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Alias) String() string {
	var builder strings.Builder
	builder.WriteString("Alias(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(a.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(a.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", key=")
	builder.WriteString(a.Key)
	builder.WriteString(", url=")
	builder.WriteString(a.URL)
	builder.WriteString(", pv=")
	builder.WriteString(fmt.Sprintf("%v", a.Pv))
	builder.WriteString(", expire=")
	builder.WriteString(a.Expire.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// AliasSlice is a parsable slice of Alias.
type AliasSlice []*Alias

func (a AliasSlice) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}