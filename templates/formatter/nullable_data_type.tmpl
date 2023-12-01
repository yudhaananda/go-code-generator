package formatter

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"strconv"
)

var nullBytes = []byte("null")

type NullableDataType[T comparable] struct {
	Data  T
	Valid bool
}

func (n *NullableDataType[T]) Scan(value any) error {
	var data T
	if value == nil {
		n.Data, n.Valid = data, false
		return nil
	}
	switch s := value.(type) {
	case []uint8:
		value := string(s)
		switch d := any(n.Data).(type) {
		case string:
			d = value
			n.Data = any(d).(T)

		case int:
			d, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)

		case uint8:
			temp, err := strconv.ParseUint(value, 10, 8)
			if err != nil {
				return err
			}
			d = uint8(temp)
			n.Data = any(d).(T)

		case uint16:
			temp, err := strconv.ParseUint(value, 10, 16)
			if err != nil {
				return err
			}
			d = uint16(temp)
			n.Data = any(d).(T)

		case uint32:
			temp, err := strconv.ParseUint(value, 10, 32)
			if err != nil {
				return err
			}
			d = uint32(temp)
			n.Data = any(d).(T)

		case uint64:
			d, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)

		case int32:
			temp, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				return err
			}
			d = int32(temp)
			n.Data = any(d).(T)

		case int64:
			d, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)

		case bool:
			d, err := strconv.ParseBool(value)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)

		case float64:
			d, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)

		case float32:
			temp, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return err
			}
			d = float32(temp)
			n.Data = any(d).(T)

		case complex64:
			temp, err := strconv.ParseComplex(value, 64)
			if err != nil {
				return err
			}
			d = complex64(temp)
			n.Data = any(d).(T)

		case complex128:
			d, err := strconv.ParseComplex(value, 64)
			if err != nil {
				return err
			}
			n.Data = any(d).(T)
		}

	default:
		n.Data = value.(T)
	}

	n.Valid = true
	return nil
}

func (n NullableDataType[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Data, nil
}

func (i *NullableDataType[T]) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return nullBytes, nil
	}
	return json.Marshal(i.Data)
}

func (i *NullableDataType[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		return nil
	}
	err := json.Unmarshal(b, &i.Data)
	i.Valid = (err == nil)
	return err
}
