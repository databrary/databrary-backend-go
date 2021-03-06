package custom_types

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql/driver"
	"errors"
)

// DataType is the 'data_type' enum type from schema 'public'.
// Generated by https://github.com/knq/xo
type DataType uint16

const (
	// DataTypeText is the 'text' DataType.
	DataTypeText = DataType(1)

	// DataTypeNumeric is the 'numeric' DataType.
	DataTypeNumeric = DataType(2)

	// DataTypeDate is the 'date' DataType.
	DataTypeDate = DataType(3)

	// DataTypeVoid is the 'void' DataType.
	DataTypeVoid = DataType(4)
)

// String returns the string value of the DataType.
func (dt DataType) String() string {
	var enumVal string

	switch dt {
	case DataTypeText:
		enumVal = "text"

	case DataTypeNumeric:
		enumVal = "numeric"

	case DataTypeDate:
		enumVal = "date"

	case DataTypeVoid:
		enumVal = "void"
	}

	return enumVal
}

// MarshalText marshals DataType into text.
func (dt DataType) MarshalText() ([]byte, error) {
	return []byte(dt.String()), nil
}

// UnmarshalText unmarshals DataType from text.
func (dt *DataType) UnmarshalText(text []byte) error {
	switch string(text) {
	case "text":
		*dt = DataTypeText

	case "numeric":
		*dt = DataTypeNumeric

	case "date":
		*dt = DataTypeDate

	case "void":
		*dt = DataTypeVoid

	default:
		return errors.New("invalid DataType")
	}

	return nil
}

// Value satisfies the sql/driver.Valuer interface for DataType.
func (dt DataType) Value() (driver.Value, error) {
	return dt.String(), nil
}

// Scan satisfies the database/sql.Scanner interface for DataType.
func (dt *DataType) Scan(src interface{}) error {
	buf, ok := src.([]byte)
	if !ok {
		return errors.New("invalid DataType")
	}

	return dt.UnmarshalText(buf)
}

// Nullable DataType. Just a wrapper around DataType.
type NullDataType struct {
	DataType DataType
	Valid    bool
}

// Implements Scanner interface.
// This is what is used to convert a column of type action from a postgres query
// into this Go type. The argument has the []byte representation of the column.
// Null columns scan to nv.Valid == false.
func (nv *NullDataType) Scan(value interface{}) error {
	if value == nil {
		nv.DataType, nv.Valid = DataType(0), false
		return nil
	}
	err := nv.DataType.Scan(value)
	if err != nil {
		nv.Valid = false
		return err
	} else {
		nv.Valid = true
		return nil
	}
}

// Implements Valuer interface
// This is what is used to convert a  Go type action to a postgres type.
// Valid == false turns into a Null value.
func (nv NullDataType) Value() (driver.Value, error) {
	if !nv.Valid {
		return nil, nil
	}
	return nv.DataType.Value()
}

// This function is used for testing SQLBoiler models, i.e. randomization
// for models that have a DataType column.
// Obviously it's not random but it doesn't really need to be anyway.
func DataTypeRandom() DataType {
	return DataTypeDate
}

// This function is used for testing SQLBoiler models, i.e. randomization
// for models that have a NullDataType column.
// Obviously it's not random but it doesn't really need to be anyway.
func NullDataTypeRandom() NullDataType {
	return NullDataType{
		DataType: DataTypeDate,
		Valid:    true,
	}
}
