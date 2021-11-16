package types

import "fmt"

type EnumCustomersStatus string

const (
	EnumCustomersStatusActive  EnumCustomersStatus = "Active"
	EnumCustomersStatusPending EnumCustomersStatus = "Pending"
	EnumCustomersStatusBlocked EnumCustomersStatus = "Blocked"
)

func (e *EnumCustomersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumCustomersStatus(s)
	case string:
		*e = EnumCustomersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumCustomersStatus: %T", src)
	}
	return nil
}

type EnumCustomersType string

const (
	EnumCustomersTypeUser        EnumCustomersType = "User"
	EnumCustomersTypeApplication EnumCustomersType = "Application"
)

func (e *EnumCustomersType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumCustomersType(s)
	case string:
		*e = EnumCustomersType(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumCustomersType: %T", src)
	}
	return nil
}
