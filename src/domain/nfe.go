package domain

import "importa-nfe/src/controllers/mapper"

type NfeParser struct {
	Data []byte
	Nfe  mapper.NfeProc
}
