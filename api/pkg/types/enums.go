package types

import "fmt"

type EnumTicketsPriority string

const (
	EnumTicketsPriorityDraft    EnumTicketsPriority = "Draft"
	EnumTicketsPriorityRegular  EnumTicketsPriority = "Regular"
	EnumTicketsPriorityPremium  EnumTicketsPriority = "Premium"
	EnumTicketsPriorityPromoted EnumTicketsPriority = "Promoted"
)

func (e *EnumTicketsPriority) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = EnumTicketsPriority(s)
	case string:
		*e = EnumTicketsPriority(s)
	default:
		return fmt.Errorf("unsupported scan type for EnumTicketsPriority: %T", src)
	}
	return nil
}
