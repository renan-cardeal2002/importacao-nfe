package domain

import "importa-nfe/internal/controllers/mapper"

type NfeParser struct {
	Data []byte
	Nfe  mapper.NfeProc
}
