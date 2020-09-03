// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Contains an array.
type ArrayValue interface {
	isArrayValue()
}

// An array of arrays.
type ArrayValueMemberArrayValues struct {
	Value []ArrayValue
}

func (*ArrayValueMemberArrayValues) isArrayValue() {}

// An array of Boolean values.
type ArrayValueMemberBooleanValues struct {
	Value []*bool
}

func (*ArrayValueMemberBooleanValues) isArrayValue() {}

// An array of integers.
type ArrayValueMemberDoubleValues struct {
	Value []*float64
}

func (*ArrayValueMemberDoubleValues) isArrayValue() {}

// An array of floating point numbers.
type ArrayValueMemberLongValues struct {
	Value []*int64
}

func (*ArrayValueMemberLongValues) isArrayValue() {}

// An array of strings.
type ArrayValueMemberStringValues struct {
	Value []*string
}

func (*ArrayValueMemberStringValues) isArrayValue() {}

// Contains the metadata for a column.
type ColumnMetadata struct {
	// The type of the column.
	ArrayBaseColumnType *int32
	// A value that indicates whether the column increments automatically.
	IsAutoIncrement *bool
	// A value that indicates whether the column is case-sensitive.
	IsCaseSensitive *bool
	// A value that indicates whether the column contains currency values.
	IsCurrency *bool
	// A value that indicates whether an integer column is signed.
	IsSigned *bool
	// The label for the column.
	Label *string
	// The name of the column.
	Name *string
	// A value that indicates whether the column is nullable.
	Nullable *int32
	// The precision value of a decimal number column.
	Precision *int32
	// The scale value of a decimal number column.
	Scale *int32
	// The name of the schema that owns the table that includes the column.
	SchemaName *string
	// The name of the table that includes the column.
	TableName *string
	// The type of the column.
	Type *int32
	// The database-specific data type of the column.
	TypeName *string
}

// Contains a value.
type Field interface {
	isField()
}

// An array of values.
type FieldMemberArrayValue struct {
	Value ArrayValue
}

func (*FieldMemberArrayValue) isField() {}

// A value of BLOB data type.
type FieldMemberBlobValue struct {
	Value []byte
}

func (*FieldMemberBlobValue) isField() {}

// A value of Boolean data type.
type FieldMemberBooleanValue struct {
	Value bool
}

func (*FieldMemberBooleanValue) isField() {}

// A value of double data type.
type FieldMemberDoubleValue struct {
	Value float64
}

func (*FieldMemberDoubleValue) isField() {}

// A NULL value.
type FieldMemberIsNull struct {
	Value bool
}

func (*FieldMemberIsNull) isField() {}

// A value of long data type.
type FieldMemberLongValue struct {
	Value int64
}

func (*FieldMemberLongValue) isField() {}

// A value of string data type.
type FieldMemberStringValue struct {
	Value string
}

func (*FieldMemberStringValue) isField() {}

// A record returned by a call.
type Record struct {
	// The values returned in the record.
	Values []Value
}

// The result set returned by a SQL statement.
type ResultFrame struct {
	// The records in the result set.
	Records []*Record
	// The result-set metadata in the result set.
	ResultSetMetadata *ResultSetMetadata
}

// The metadata of the result set returned by a SQL statement.
type ResultSetMetadata struct {
	// The number of columns in the result set.
	ColumnCount *int64
	// The metadata of the columns in the result set.
	ColumnMetadata []*ColumnMetadata
}

// Options that control how the result set is returned.
type ResultSetOptions struct {
	// A value that indicates how a field of DECIMAL type is represented in the
	// response. The value of STRING, the default, specifies that it is converted to a
	// String value. The value of DOUBLE_OR_LONG specifies that it is converted to a
	// Long value if its scale is 0, or to a Double value otherwise. Conversion to
	// Double or Long can result in roundoff errors due to precision loss. We recommend
	// converting to String, especially when working with currency values.
	DecimalReturnType DecimalReturnType
}

// A parameter used in a SQL statement.
type SqlParameter struct {
	// The name of the parameter.
	Name *string
	// A hint that specifies the correct object type for data type mapping. Values:
	//
	//
	// * DECIMAL - The corresponding String parameter value is sent as an object of
	// DECIMAL type to the database.
	//
	//     * TIMESTAMP - The corresponding String
	// parameter value is sent as an object of TIMESTAMP type to the database. The
	// accepted format is YYYY-MM-DD HH:MM:SS[.FFF].
	//
	//     * TIME - The corresponding
	// String parameter value is sent as an object of TIME type to the database. The
	// accepted format is HH:MM:SS[.FFF].
	//
	//     * DATE - The corresponding String
	// parameter value is sent as an object of DATE type to the database. The accepted
	// format is YYYY-MM-DD.
	TypeHint TypeHint
	// The value of the parameter.
	Value Field
}

// The result of a SQL statement.  <important> <p>This data type is deprecated.</p>
// </important>
type SqlStatementResult struct {
	// The number of records updated by a SQL statement.
	NumberOfRecordsUpdated *int64
	// The result set of the SQL statement.
	ResultFrame *ResultFrame
}

// A structure value returned by a call.
type StructValue struct {
	// The attributes returned in the record.
	Attributes []Value
}

// The response elements represent the results of an update.
type UpdateResult struct {
	// Values for fields generated during the request.
	GeneratedFields []Field
}

// Contains the value of a column.  <important> <p>This data type is
// deprecated.</p> </important>
type Value interface {
	isValue()
}

// An array of column values.
type ValueMemberArrayValues struct {
	Value []Value
}

func (*ValueMemberArrayValues) isValue() {}

// A value for a column of big integer data type.
type ValueMemberBigIntValue struct {
	Value int64
}

func (*ValueMemberBigIntValue) isValue() {}

// A value for a column of BIT data type.
type ValueMemberBitValue struct {
	Value bool
}

func (*ValueMemberBitValue) isValue() {}

// A value for a column of BLOB data type.
type ValueMemberBlobValue struct {
	Value []byte
}

func (*ValueMemberBlobValue) isValue() {}

// A value for a column of double data type.
type ValueMemberDoubleValue struct {
	Value float64
}

func (*ValueMemberDoubleValue) isValue() {}

// A value for a column of integer data type.
type ValueMemberIntValue struct {
	Value int32
}

func (*ValueMemberIntValue) isValue() {}

// A NULL value.
type ValueMemberIsNull struct {
	Value bool
}

func (*ValueMemberIsNull) isValue() {}

// A value for a column of real data type.
type ValueMemberRealValue struct {
	Value float32
}

func (*ValueMemberRealValue) isValue() {}

// A value for a column of string data type.
type ValueMemberStringValue struct {
	Value string
}

func (*ValueMemberStringValue) isValue() {}

// A value for a column of STRUCT data type.
type ValueMemberStructValue struct {
	Value *StructValue
}

func (*ValueMemberStructValue) isValue() {}

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte
}

func (*UnknownUnionMember) isArrayValue() {}
func (*UnknownUnionMember) isField()      {}
func (*UnknownUnionMember) isValue()      {}
