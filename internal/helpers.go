package internal

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"strconv"
)

func ConvertToInt4(value interface{}) (pgtype.Int4, error) {
	var intVal pgtype.Int4
	switch v := value.(type) {
	case string:
		int64val, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return pgtype.Int4{}, err
		}
		if err := intVal.Scan(int32(int64val)); err != nil {
			return pgtype.Int4{}, err
		}
	default:
		return pgtype.Int4{}, fmt.Errorf("unsupported text for pgtype int4 conversion: %T", v)
	}
	return intVal, nil
}

func ConvertToBool(value interface{}) (pgtype.Bool, error) {
	var boolVal pgtype.Bool
	switch v := value.(type) {
	case bool:
		if err := boolVal.Scan(v); err != nil {
			return pgtype.Bool{}, err
		}
	default:
		return pgtype.Bool{}, fmt.Errorf("unsupported text for pgtype bool conversion: %T", v)
	}
	return boolVal, nil
}
