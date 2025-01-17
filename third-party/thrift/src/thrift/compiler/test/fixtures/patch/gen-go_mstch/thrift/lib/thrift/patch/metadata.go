// @generated by Thrift for [[[ program path ]]]
// This file is probably not the place you want to edit!

package patch // [[[ program thrift source path ]]]

import (
    standard "thrift/lib/thrift/standard"
    id "thrift/lib/thrift/id"
    thrift "github.com/facebook/fbthrift/thrift/lib/go/thrift"
    metadata "github.com/facebook/fbthrift/thrift/lib/thrift/metadata"
)

var _ = standard.GoUnusedProtection__
var _ = id.GoUnusedProtection__
// (needed to ensure safety because of naive import list construction)
var _ = thrift.ZERO
var _ = metadata.GoUnusedProtection__

// Premade Thrift types
var (
    premadeThriftType_bool = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BOOL_TYPE.Ptr(),
            )
    premadeThriftType_byte = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BYTE_TYPE.Ptr(),
            )
    premadeThriftType_i16 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I16_TYPE.Ptr(),
            )
    premadeThriftType_i32 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I32_TYPE.Ptr(),
            )
    premadeThriftType_i64 = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_I64_TYPE.Ptr(),
            )
    premadeThriftType_float = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_FLOAT_TYPE.Ptr(),
            )
    premadeThriftType_double = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_DOUBLE_TYPE.Ptr(),
            )
    premadeThriftType_string = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_STRING_TYPE.Ptr(),
            )
    premadeThriftType_binary = metadata.NewThriftType().SetTPrimitive(
        metadata.ThriftPrimitiveType_THRIFT_BINARY_TYPE.Ptr(),
            )
    premadeThriftType_standard_ByteBuffer = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("standard.ByteBuffer").
            SetUnderlyingType(premadeThriftType_binary),
            )
    premadeThriftType_id_FieldId = metadata.NewThriftType().SetTTypedef(
        metadata.NewThriftTypedefType().
            SetName("id.FieldId").
            SetUnderlyingType(premadeThriftType_i16),
            )
    premadeThriftType_list_i16 = metadata.NewThriftType().SetTList(
        metadata.NewThriftListType().
            SetValueType(premadeThriftType_i16),
            )
)

var structMetadatas = []*metadata.ThriftStruct{
    metadata.NewThriftStruct().
    SetName("patch.GeneratePatch").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("patch.AssignOnlyPatch").
    SetIsUnion(false),
    metadata.NewThriftStruct().
    SetName("patch.BoolPatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(9).
    SetName("invert").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.BytePatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_byte),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_byte),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.I16Patch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_i16),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_i16),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.I32Patch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_i32),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_i32),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.I64Patch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_i64),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_i64),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.FloatPatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_float),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_float),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.DoublePatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_double),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("add").
    SetIsOptional(false).
    SetType(premadeThriftType_double),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.StringPatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("prepend").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
            metadata.NewThriftField().
    SetId(9).
    SetName("append").
    SetIsOptional(false).
    SetType(premadeThriftType_string),
        },
    ),
    metadata.NewThriftStruct().
    SetName("patch.BinaryPatch").
    SetIsUnion(false).
    SetFields(
        []*metadata.ThriftField{
            metadata.NewThriftField().
    SetId(1).
    SetName("assign").
    SetIsOptional(true).
    SetType(premadeThriftType_standard_ByteBuffer),
            metadata.NewThriftField().
    SetId(2).
    SetName("clear").
    SetIsOptional(false).
    SetType(premadeThriftType_bool),
            metadata.NewThriftField().
    SetId(8).
    SetName("prepend").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_ByteBuffer),
            metadata.NewThriftField().
    SetId(9).
    SetName("append").
    SetIsOptional(false).
    SetType(premadeThriftType_standard_ByteBuffer),
        },
    ),
}

var exceptionMetadatas = []*metadata.ThriftException{
}

var enumMetadatas = []*metadata.ThriftEnum{
}

var serviceMetadatas = []*metadata.ThriftService{
}

// GetThriftMetadata returns complete Thrift metadata for current and imported packages.
func GetThriftMetadata() *metadata.ThriftMetadata {
    includedEnumsMetadatas := []map[string]*metadata.ThriftEnum{
        GetEnumsMetadata(),
        standard.GetEnumsMetadata(),
        id.GetEnumsMetadata(),
    }
    includedStructsMetadatas := []map[string]*metadata.ThriftStruct{
        GetStructsMetadata(),
        standard.GetStructsMetadata(),
        id.GetStructsMetadata(),
    }
    includedExceptionsMetadatas := []map[string]*metadata.ThriftException{
        GetExceptionsMetadata(),
        standard.GetExceptionsMetadata(),
        id.GetExceptionsMetadata(),
    }
    includedServicesMetadatas := []map[string]*metadata.ThriftService{
        GetServicesMetadata(),
        standard.GetServicesMetadata(),
        id.GetServicesMetadata(),
    }

	allEnums := make(map[string]*metadata.ThriftEnum)
	allStructs := make(map[string]*metadata.ThriftStruct)
	allExceptions := make(map[string]*metadata.ThriftException)
    allServices := make(map[string]*metadata.ThriftService)

    for _, includedEnumsMetadata := range includedEnumsMetadatas {
        for enumName, thriftEnum := range includedEnumsMetadata {
            allEnums[enumName] = thriftEnum
        }
    }
    for _, includedStructsMetadata := range includedStructsMetadatas {
        for structName, thriftStruct := range includedStructsMetadata {
            allStructs[structName] = thriftStruct
        }
    }
    for _, includedExceptionsMetadata := range includedExceptionsMetadatas {
        for exceptionName, thriftException := range includedExceptionsMetadata {
            allExceptions[exceptionName] = thriftException
        }
    }
    for _, includedServicesMetadata := range includedServicesMetadatas {
        for serviceName, thriftService := range includedServicesMetadata {
            allServices[serviceName] = thriftService
        }
    }

    return metadata.NewThriftMetadata().
        SetEnums(allEnums).
        SetStructs(allStructs).
        SetExceptions(allExceptions).
        SetServices(allServices)
}

// GetStructsMetadata returns Thrift metadata for enums in the current package.
func GetEnumsMetadata() map[string]*metadata.ThriftEnum {
    result := make(map[string]*metadata.ThriftEnum)
    for _, enumMetadata := range enumMetadatas {
        result[enumMetadata.GetName()] = enumMetadata
    }
    return result
}

// GetStructsMetadata returns Thrift metadata for structs in the current package.
func GetStructsMetadata() map[string]*metadata.ThriftStruct {
    result := make(map[string]*metadata.ThriftStruct)
    for _, structMetadata := range structMetadatas {
        result[structMetadata.GetName()] = structMetadata
    }
    return result
}

// GetStructsMetadata returns Thrift metadata for exceptions in the current package.
func GetExceptionsMetadata() map[string]*metadata.ThriftException {
    result := make(map[string]*metadata.ThriftException)
    for _, exceptionMetadata := range exceptionMetadatas {
        result[exceptionMetadata.GetName()] = exceptionMetadata
    }
    return result
}

// GetStructsMetadata returns Thrift metadata for services in the current package.
func GetServicesMetadata() map[string]*metadata.ThriftService {
    result := make(map[string]*metadata.ThriftService)
    for _, serviceMetadata := range serviceMetadatas {
        result[serviceMetadata.GetName()] = serviceMetadata
    }
    return result
}
