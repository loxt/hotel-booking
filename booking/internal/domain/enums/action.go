package enums

type Action string

const (
	Pay    Action = "pay"
	Finish Action = "finish"
	Cancel Action = "cancel"
	Refund Action = "refund"
	Reopen Action = "reopen"
)
