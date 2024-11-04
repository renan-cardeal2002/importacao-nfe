package domain

import (
	"importa-nfe/internal/infrastructure/adapter/inbound/controllers/mapper"
)

type NfeParser struct {
	Data []byte
	Nfe  mapper.NfeProc
}
