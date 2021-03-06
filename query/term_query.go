package query

import (
	"fmt"
	"github.com/MintegralTech/juno/datastruct"
	"github.com/MintegralTech/juno/debug"
	"github.com/MintegralTech/juno/document"
	"github.com/MintegralTech/juno/helpers"
	"github.com/MintegralTech/juno/index"
	"github.com/MintegralTech/juno/marshal"
	"strings"
)

type TermQuery struct {
	iterator datastruct.Iterator
	label    string
	debugs   *debug.Debug
}

func NewTermQuery(iter datastruct.Iterator) (tq *TermQuery) {
	if iter == nil {
		return nil
	}
	return &TermQuery{
		iterator: iter,
	}
}

func (tq *TermQuery) SetLabel(label string) {
	if tq != nil {
		tq.label = label
	}
}

func (tq *TermQuery) Next() {
	if tq == nil || tq.iterator == nil {
		return
	}
	if tq.debugs != nil && tq.debugs.Level == 2 {
		tq.debugs.NextCounter++
	}
	tq.iterator.Next()
}

func (tq *TermQuery) GetGE(id document.DocId) (document.DocId, error) {
	if tq == nil || tq.iterator == nil {
		return 0, helpers.DocumentError
	}

	element := tq.iterator.GetGE(id)
	if tq.debugs != nil && tq.debugs.Level == 2 {
		tq.debugs.LECounter++
	}
	if element == nil {
		if tq.debugs != nil {
			tq.debugs.AddDebugMsg(fmt.Sprintf("docId: %d, reason: %v", id, helpers.ElementNotfound))
		}
		return 0, helpers.ElementNotfound
	}
	return element.Key(), nil
}

func (tq *TermQuery) Current() (document.DocId, error) {
	if tq == nil || tq.iterator == nil {
		return 0, helpers.DocumentError
	}
	element := tq.iterator.Current()
	if element == nil {
		return 0, helpers.NoMoreData
	}
	return element.Key(), nil
}

func (tq *TermQuery) DebugInfo() *debug.Debug {
	if tq != nil && tq.debugs != nil {
		tq.debugs.FieldName = tq.iterator.GetFieldName()
		return tq.debugs
	}
	return nil
}

func (tq *TermQuery) Marshal() map[string]interface{} {
	if tq == nil {
		return map[string]interface{}{}
	}
	res := make(map[string]interface{}, 1)
	fields := strings.Split(tq.iterator.GetFieldName(), index.SEP)
	res["="] = []string{fields[0], fields[1]}
	return res
}

func (tq *TermQuery) MarshalV2() *marshal.MarshalInfo {
	if tq == nil {
		return nil
	}
	fields := strings.Split(tq.iterator.GetFieldName(), index.SEP)
	info := &marshal.MarshalInfo{
		Name:       fields[0],
		QueryValue: fields[1],
		Operation:  "=",
		Nodes:      nil,
	}
	if tq.label != "" {
		info.Label = tq.label
	}
	return info
}
func (tq *TermQuery) UnmarshalV2(idx index.Index, marshalInfo *marshal.MarshalInfo) Query {
	if marshalInfo == nil || marshalInfo.Operation != "=" {
		return nil
	}
	return NewTermQuery(idx.GetInvertedIndex().Iterator(marshalInfo.Name, fmt.Sprintf("%s", marshalInfo.QueryValue)))
}

func (tq *TermQuery) Unmarshal(idx index.Index, res map[string]interface{}) Query {
	v, ok := res["="]
	if !ok {
		return nil
	}
	return NewTermQuery(idx.GetInvertedIndex().Iterator(fmt.Sprint(v.([]string)[0]), fmt.Sprint(v.([]string)[1])))
}

func (tq *TermQuery) SetDebug(level int) {
	if tq == nil {
		return
	}
	if tq.debugs == nil {
		tq.debugs = debug.NewDebug(level, "TermQuery")
	}
}
