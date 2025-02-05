package Payloadschemas

import "context"



type UpdateEmailByIDPayload struct {
	NewEmail string
	ID       uint16
}

type contextstruct struct {
	ctx context.Context
	cancel context.CancelFunc
}