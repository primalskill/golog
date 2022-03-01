package golog

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// ProdLog holds the production logger dependencies.
type ProdLog struct {
	z *zerolog.Logger
}

type payload struct {
	Msg string `json:"msg"`
	Meta []Meta `json:"meta,omitempty"`
}

// NewProductionLog returns a new production logger.
func NewProductionLog() *ProdLog {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = "t"
	zl := zerolog.New(os.Stdout)

	return &ProdLog{
		z: &zl,
	}
}

// Info
func (p *ProdLog) Info(msg string, meta ...Meta) {
	pl, isErr := p.genJSON(msg, meta...)
	if isErr {	
		p.z.Error().Timestamp().RawJSON("payload", pl).Send()
	}
	
	p.z.Info().Timestamp().RawJSON("payload", pl).Send()	
}

// Warn
func (p *ProdLog) Warn(msg string, meta ...Meta) {
	pl, isErr := p.genJSON(msg, meta...)
	if isErr {
		p.z.Error().Timestamp().RawJSON("payload", pl).Send()
	}
	
	p.z.Warn().Timestamp().RawJSON("payload", pl).Send()	
}

// Error
func (p *ProdLog) Error(err error, meta ...Meta) {
	pl, isErr := p.genJSON(err.Error(), meta...)
	if isErr {
		p.z.Error().Timestamp().RawJSON("payload", pl).Send()
	}
	
	p.z.Error().Timestamp().RawJSON("payload", pl).Send()
}


func (p *ProdLog) genJSON(msg string, meta ...Meta) (ret []byte, isErr bool) {
	pret := payload{
		Msg: msg,
		Meta: meta,
	}

	ret, err := json.Marshal(pret)
	if err != nil {
		ret = []byte(fmt.Sprintf(`{"msg":"Can't marshal JSON log meta: %s"}`, err.Error()))
		isErr = true
		return
	}

	return
}

