package mytype

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type NullTime struct {
	sql.NullTime
}

func (n *NullTime) MarshalJSON() ([]byte, error) {
	fmt.Println(n)
	if n == nil || !n.Valid {
		return nil, nil
	}

	return json.Marshal(n.Time)
}
