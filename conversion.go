package wrapperspb_conversion

import "google.golang.org/protobuf/types/known/wrapperspb"

func FromDoubleValue(val *float64) *wrapperspb.DoubleValue {
	if val == nil {
		return nil
	}
	return wrapperspb.Double(*val)
}

func ToDoubleValue(val *wrapperspb.DoubleValue) *float64 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromFloatValue(val *float32) *wrapperspb.FloatValue {
	if val == nil {
		return nil
	}
	return wrapperspb.Float(*val)
}

func ToFloatValue(val *wrapperspb.FloatValue) *float32 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromInt64Value(val *int64) *wrapperspb.Int64Value {
	if val == nil {
		return nil
	}
	return wrapperspb.Int64(*val)
}

func ToInt64Value(val *wrapperspb.Int64Value) *int64 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromInt32Value(val *int32) *wrapperspb.Int32Value {
	if val == nil {
		return nil
	}
	return wrapperspb.Int32(*val)
}

func ToInt32Value(val *wrapperspb.Int32Value) *int32 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromUInt64Value(val *uint64) *wrapperspb.UInt64Value {
	if val == nil {
		return nil
	}
	return wrapperspb.UInt64(*val)
}

func ToUInt64Value(val *wrapperspb.UInt64Value) *uint64 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromUInt32Value(val *uint32) *wrapperspb.UInt32Value {
	if val == nil {
		return nil
	}
	return wrapperspb.UInt32(*val)
}

func ToUInt32Value(val *wrapperspb.UInt32Value) *uint32 {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromBoolValue(val *bool) *wrapperspb.BoolValue {
	if val == nil {
		return nil
	}
	return wrapperspb.Bool(*val)
}

func ToBoolValue(val *wrapperspb.BoolValue) *bool {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromStringValue(val *string) *wrapperspb.StringValue {
	if val == nil {
		return nil
	}
	return wrapperspb.String(*val)
}

func ToStringValue(val *wrapperspb.StringValue) *string {
	if val == nil {
		return nil
	}
	return &val.Value
}

func FromBytesValue(val *[]byte) *wrapperspb.BytesValue {
	if val == nil {
		return nil
	}
	return wrapperspb.Bytes(*val)
}

func ToBytesValue(val *wrapperspb.BytesValue) *[]byte {
	if val == nil {
		return nil
	}
	return &val.Value
}