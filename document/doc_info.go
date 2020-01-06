package document

type DocId uint64
type FieldType int64
type IndexType int64

const (
	InvertedIndexType = iota
	StorageIndexType
	BothIndexType
)

const (
	INVERT_INDEX_STRING_TYPE = iota
	INVERT_INDEX_INT_TYPE
	INVERT_INDEX_FLOAT_TYPE
	INVERT_INDEX_STRING_SLICE_TYPE
	INVERT_INDEX_INT_SLICE_TYPE
	INVERT_INDEX_FLOAT_SLICE_TYPE
	INVERT_INDEX_STRING_MAP_TYPE
	INVERT_INDEX_INT_MAP_TYPE
	INVERT_INDEX_FLOAT_MAP_TYPE
	STORAGE_INDEX_STRING_TYPE
	STORAGE_INDEX_INT_TYPE
	STORAGE_INDEX_FLOAT_TYPE
	STORAGE_INDEX_STRING_SLICE_TYPE
	STORAGE_INDEX_INT_SLICE_TYPE
	STORAGE_INDEX_FLOAT_SLICE_TYPE
	STORAGE_INDEX_STRING_MAP_TYPE
	STORAGE_INDEX_INT_MAP_TYPE
	STORAGE_INDEX_FLOAT_MAP_TYPE
)

const (
	BoolFieldType FieldType = iota
	IntFieldType
	FloatFieldType
	StringFieldType
	SliceFieldType
	MapFieldType
	SelfDefinedFieldType
	DefaultFieldType = StringFieldType
)

type Field struct {
	Name      string
	IndexType IndexType
	Value     interface{}
	ValueType FieldType
}

type DocInfo struct {
	Id     DocId
	Fields []*Field
}
