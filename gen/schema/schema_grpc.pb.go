// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: schema/schema.proto

package schema

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SchemaService_CreateSchema_FullMethodName            = "/schema.SchemaService/CreateSchema"
	SchemaService_CreateSchemaVersion_FullMethodName     = "/schema.SchemaService/CreateSchemaVersion"
	SchemaService_CreateSchemaWithVersion_FullMethodName = "/schema.SchemaService/CreateSchemaWithVersion"
	SchemaService_CreateVersions_FullMethodName          = "/schema.SchemaService/CreateVersions"
	SchemaService_ListSchemas_FullMethodName             = "/schema.SchemaService/ListSchemas"
	SchemaService_ListVersions_FullMethodName            = "/schema.SchemaService/ListVersions"
	SchemaService_ListSchemasWithVersions_FullMethodName = "/schema.SchemaService/ListSchemasWithVersions"
	SchemaService_ListTrashSchemas_FullMethodName        = "/schema.SchemaService/ListTrashSchemas"
	SchemaService_GetSchema_FullMethodName               = "/schema.SchemaService/GetSchema"
	SchemaService_GetVersion_FullMethodName              = "/schema.SchemaService/GetVersion"
	SchemaService_GetSchemaWithVersions_FullMethodName   = "/schema.SchemaService/GetSchemaWithVersions"
	SchemaService_GetSchemaWithVersion_FullMethodName    = "/schema.SchemaService/GetSchemaWithVersion"
	SchemaService_GetVersionIdList_FullMethodName        = "/schema.SchemaService/GetVersionIdList"
	SchemaService_RemoveSchema_FullMethodName            = "/schema.SchemaService/RemoveSchema"
	SchemaService_RemoveSchemaVersion_FullMethodName     = "/schema.SchemaService/RemoveSchemaVersion"
	SchemaService_CloneSchema_FullMethodName             = "/schema.SchemaService/CloneSchema"
	SchemaService_CloneSchemaVersion_FullMethodName      = "/schema.SchemaService/CloneSchemaVersion"
	SchemaService_SaveSchema_FullMethodName              = "/schema.SchemaService/SaveSchema"
	SchemaService_SaveSchemaVersion_FullMethodName       = "/schema.SchemaService/SaveSchemaVersion"
	SchemaService_GetSchemasCount_FullMethodName         = "/schema.SchemaService/GetSchemasCount"
	SchemaService_GetFieldSchema_FullMethodName          = "/schema.SchemaService/GetFieldSchema"
	SchemaService_MoveSchemaToGroup_FullMethodName       = "/schema.SchemaService/MoveSchemaToGroup"
	SchemaService_MarkSchemaAsTrash_FullMethodName       = "/schema.SchemaService/MarkSchemaAsTrash"
)

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchemaServiceClient interface {
	// CREATE
	CreateSchema(ctx context.Context, in *CreateSchemaReq, opts ...grpc.CallOption) (*Schema, error)
	CreateSchemaVersion(ctx context.Context, in *CreateSchemaVersionReq, opts ...grpc.CallOption) (*SchemaVersion, error)
	CreateSchemaWithVersion(ctx context.Context, in *SchemaWithVersion, opts ...grpc.CallOption) (*SchemaWithVersion, error)
	CreateVersions(ctx context.Context, in *Versions, opts ...grpc.CallOption) (*Versions, error)
	// LIST
	ListSchemas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Schemas, error)
	ListVersions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Versions, error)
	ListSchemasWithVersions(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SchemasWithVersions, error)
	ListTrashSchemas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SchemasWithVersions, error)
	// GET
	GetSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Schema, error)
	GetVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaVersion, error)
	GetSchemaWithVersions(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaWithVersions, error)
	GetSchemaWithVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaWithVersion, error)
	GetVersionIdList(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*IdsRequest, error)
	// REMOVE
	RemoveSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveSchemaVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error)
	// CLONE
	CloneSchema(ctx context.Context, in *CloneSchemaReq, opts ...grpc.CallOption) (*SchemaWithVersions, error)
	CloneSchemaVersion(ctx context.Context, in *CloneSchemaVersionReq, opts ...grpc.CallOption) (*SchemaVersion, error)
	// SAVE/UPDATE
	SaveSchema(ctx context.Context, in *Schema, opts ...grpc.CallOption) (*Schema, error)
	SaveSchemaVersion(ctx context.Context, in *SchemaVersion, opts ...grpc.CallOption) (*SchemaVersion, error)
	// Miscellaneous
	GetSchemasCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SchemaCount, error)
	GetFieldSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*FieldSchema, error)
	MoveSchemaToGroup(ctx context.Context, in *MoveSchema, opts ...grpc.CallOption) (*Empty, error)
	MarkSchemaAsTrash(ctx context.Context, in *TrashSchema, opts ...grpc.CallOption) (*Empty, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) CreateSchema(ctx context.Context, in *CreateSchemaReq, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, SchemaService_CreateSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CreateSchemaVersion(ctx context.Context, in *CreateSchemaVersionReq, opts ...grpc.CallOption) (*SchemaVersion, error) {
	out := new(SchemaVersion)
	err := c.cc.Invoke(ctx, SchemaService_CreateSchemaVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CreateSchemaWithVersion(ctx context.Context, in *SchemaWithVersion, opts ...grpc.CallOption) (*SchemaWithVersion, error) {
	out := new(SchemaWithVersion)
	err := c.cc.Invoke(ctx, SchemaService_CreateSchemaWithVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CreateVersions(ctx context.Context, in *Versions, opts ...grpc.CallOption) (*Versions, error) {
	out := new(Versions)
	err := c.cc.Invoke(ctx, SchemaService_CreateVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListSchemas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Schemas, error) {
	out := new(Schemas)
	err := c.cc.Invoke(ctx, SchemaService_ListSchemas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListVersions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Versions, error) {
	out := new(Versions)
	err := c.cc.Invoke(ctx, SchemaService_ListVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListSchemasWithVersions(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SchemasWithVersions, error) {
	out := new(SchemasWithVersions)
	err := c.cc.Invoke(ctx, SchemaService_ListSchemasWithVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListTrashSchemas(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SchemasWithVersions, error) {
	out := new(SchemasWithVersions)
	err := c.cc.Invoke(ctx, SchemaService_ListTrashSchemas_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, SchemaService_GetSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaVersion, error) {
	out := new(SchemaVersion)
	err := c.cc.Invoke(ctx, SchemaService_GetVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetSchemaWithVersions(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaWithVersions, error) {
	out := new(SchemaWithVersions)
	err := c.cc.Invoke(ctx, SchemaService_GetSchemaWithVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetSchemaWithVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*SchemaWithVersion, error) {
	out := new(SchemaWithVersion)
	err := c.cc.Invoke(ctx, SchemaService_GetSchemaWithVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetVersionIdList(ctx context.Context, in *IdsRequest, opts ...grpc.CallOption) (*IdsRequest, error) {
	out := new(IdsRequest)
	err := c.cc.Invoke(ctx, SchemaService_GetVersionIdList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) RemoveSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, SchemaService_RemoveSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) RemoveSchemaVersion(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, SchemaService_RemoveSchemaVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CloneSchema(ctx context.Context, in *CloneSchemaReq, opts ...grpc.CallOption) (*SchemaWithVersions, error) {
	out := new(SchemaWithVersions)
	err := c.cc.Invoke(ctx, SchemaService_CloneSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CloneSchemaVersion(ctx context.Context, in *CloneSchemaVersionReq, opts ...grpc.CallOption) (*SchemaVersion, error) {
	out := new(SchemaVersion)
	err := c.cc.Invoke(ctx, SchemaService_CloneSchemaVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) SaveSchema(ctx context.Context, in *Schema, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, SchemaService_SaveSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) SaveSchemaVersion(ctx context.Context, in *SchemaVersion, opts ...grpc.CallOption) (*SchemaVersion, error) {
	out := new(SchemaVersion)
	err := c.cc.Invoke(ctx, SchemaService_SaveSchemaVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetSchemasCount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SchemaCount, error) {
	out := new(SchemaCount)
	err := c.cc.Invoke(ctx, SchemaService_GetSchemasCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) GetFieldSchema(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*FieldSchema, error) {
	out := new(FieldSchema)
	err := c.cc.Invoke(ctx, SchemaService_GetFieldSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) MoveSchemaToGroup(ctx context.Context, in *MoveSchema, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, SchemaService_MoveSchemaToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) MarkSchemaAsTrash(ctx context.Context, in *TrashSchema, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, SchemaService_MarkSchemaAsTrash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
// All implementations must embed UnimplementedSchemaServiceServer
// for forward compatibility
type SchemaServiceServer interface {
	// CREATE
	CreateSchema(context.Context, *CreateSchemaReq) (*Schema, error)
	CreateSchemaVersion(context.Context, *CreateSchemaVersionReq) (*SchemaVersion, error)
	CreateSchemaWithVersion(context.Context, *SchemaWithVersion) (*SchemaWithVersion, error)
	CreateVersions(context.Context, *Versions) (*Versions, error)
	// LIST
	ListSchemas(context.Context, *Empty) (*Schemas, error)
	ListVersions(context.Context, *Empty) (*Versions, error)
	ListSchemasWithVersions(context.Context, *Filter) (*SchemasWithVersions, error)
	ListTrashSchemas(context.Context, *Empty) (*SchemasWithVersions, error)
	// GET
	GetSchema(context.Context, *IdRequest) (*Schema, error)
	GetVersion(context.Context, *IdRequest) (*SchemaVersion, error)
	GetSchemaWithVersions(context.Context, *IdRequest) (*SchemaWithVersions, error)
	GetSchemaWithVersion(context.Context, *IdRequest) (*SchemaWithVersion, error)
	GetVersionIdList(context.Context, *IdsRequest) (*IdsRequest, error)
	// REMOVE
	RemoveSchema(context.Context, *IdRequest) (*Empty, error)
	RemoveSchemaVersion(context.Context, *IdRequest) (*Empty, error)
	// CLONE
	CloneSchema(context.Context, *CloneSchemaReq) (*SchemaWithVersions, error)
	CloneSchemaVersion(context.Context, *CloneSchemaVersionReq) (*SchemaVersion, error)
	// SAVE/UPDATE
	SaveSchema(context.Context, *Schema) (*Schema, error)
	SaveSchemaVersion(context.Context, *SchemaVersion) (*SchemaVersion, error)
	// Miscellaneous
	GetSchemasCount(context.Context, *Empty) (*SchemaCount, error)
	GetFieldSchema(context.Context, *IdRequest) (*FieldSchema, error)
	MoveSchemaToGroup(context.Context, *MoveSchema) (*Empty, error)
	MarkSchemaAsTrash(context.Context, *TrashSchema) (*Empty, error)
	mustEmbedUnimplementedSchemaServiceServer()
}

// UnimplementedSchemaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchemaServiceServer struct {
}

func (UnimplementedSchemaServiceServer) CreateSchema(context.Context, *CreateSchemaReq) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}
func (UnimplementedSchemaServiceServer) CreateSchemaVersion(context.Context, *CreateSchemaVersionReq) (*SchemaVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchemaVersion not implemented")
}
func (UnimplementedSchemaServiceServer) CreateSchemaWithVersion(context.Context, *SchemaWithVersion) (*SchemaWithVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchemaWithVersion not implemented")
}
func (UnimplementedSchemaServiceServer) CreateVersions(context.Context, *Versions) (*Versions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVersions not implemented")
}
func (UnimplementedSchemaServiceServer) ListSchemas(context.Context, *Empty) (*Schemas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}
func (UnimplementedSchemaServiceServer) ListVersions(context.Context, *Empty) (*Versions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}
func (UnimplementedSchemaServiceServer) ListSchemasWithVersions(context.Context, *Filter) (*SchemasWithVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemasWithVersions not implemented")
}
func (UnimplementedSchemaServiceServer) ListTrashSchemas(context.Context, *Empty) (*SchemasWithVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrashSchemas not implemented")
}
func (UnimplementedSchemaServiceServer) GetSchema(context.Context, *IdRequest) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedSchemaServiceServer) GetVersion(context.Context, *IdRequest) (*SchemaVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedSchemaServiceServer) GetSchemaWithVersions(context.Context, *IdRequest) (*SchemaWithVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaWithVersions not implemented")
}
func (UnimplementedSchemaServiceServer) GetSchemaWithVersion(context.Context, *IdRequest) (*SchemaWithVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemaWithVersion not implemented")
}
func (UnimplementedSchemaServiceServer) GetVersionIdList(context.Context, *IdsRequest) (*IdsRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersionIdList not implemented")
}
func (UnimplementedSchemaServiceServer) RemoveSchema(context.Context, *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSchema not implemented")
}
func (UnimplementedSchemaServiceServer) RemoveSchemaVersion(context.Context, *IdRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSchemaVersion not implemented")
}
func (UnimplementedSchemaServiceServer) CloneSchema(context.Context, *CloneSchemaReq) (*SchemaWithVersions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneSchema not implemented")
}
func (UnimplementedSchemaServiceServer) CloneSchemaVersion(context.Context, *CloneSchemaVersionReq) (*SchemaVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CloneSchemaVersion not implemented")
}
func (UnimplementedSchemaServiceServer) SaveSchema(context.Context, *Schema) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSchema not implemented")
}
func (UnimplementedSchemaServiceServer) SaveSchemaVersion(context.Context, *SchemaVersion) (*SchemaVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSchemaVersion not implemented")
}
func (UnimplementedSchemaServiceServer) GetSchemasCount(context.Context, *Empty) (*SchemaCount, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchemasCount not implemented")
}
func (UnimplementedSchemaServiceServer) GetFieldSchema(context.Context, *IdRequest) (*FieldSchema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFieldSchema not implemented")
}
func (UnimplementedSchemaServiceServer) MoveSchemaToGroup(context.Context, *MoveSchema) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MoveSchemaToGroup not implemented")
}
func (UnimplementedSchemaServiceServer) MarkSchemaAsTrash(context.Context, *TrashSchema) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkSchemaAsTrash not implemented")
}
func (UnimplementedSchemaServiceServer) mustEmbedUnimplementedSchemaServiceServer() {}

// UnsafeSchemaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchemaServiceServer will
// result in compilation errors.
type UnsafeSchemaServiceServer interface {
	mustEmbedUnimplementedSchemaServiceServer()
}

func RegisterSchemaServiceServer(s grpc.ServiceRegistrar, srv SchemaServiceServer) {
	s.RegisterService(&SchemaService_ServiceDesc, srv)
}

func _SchemaService_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CreateSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CreateSchema(ctx, req.(*CreateSchemaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CreateSchemaVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CreateSchemaVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CreateSchemaVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CreateSchemaVersion(ctx, req.(*CreateSchemaVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CreateSchemaWithVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaWithVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CreateSchemaWithVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CreateSchemaWithVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CreateSchemaWithVersion(ctx, req.(*SchemaWithVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CreateVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Versions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CreateVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CreateVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CreateVersions(ctx, req.(*Versions))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ListSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListSchemas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ListVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListVersions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListSchemasWithVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListSchemasWithVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ListSchemasWithVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListSchemasWithVersions(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListTrashSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListTrashSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_ListTrashSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListTrashSchemas(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchema(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetVersion(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetSchemaWithVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchemaWithVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetSchemaWithVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchemaWithVersions(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetSchemaWithVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchemaWithVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetSchemaWithVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchemaWithVersion(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetVersionIdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetVersionIdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetVersionIdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetVersionIdList(ctx, req.(*IdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_RemoveSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).RemoveSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_RemoveSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).RemoveSchema(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_RemoveSchemaVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).RemoveSchemaVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_RemoveSchemaVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).RemoveSchemaVersion(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CloneSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneSchemaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CloneSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CloneSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CloneSchema(ctx, req.(*CloneSchemaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CloneSchemaVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloneSchemaVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CloneSchemaVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_CloneSchemaVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CloneSchemaVersion(ctx, req.(*CloneSchemaVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_SaveSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Schema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).SaveSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_SaveSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).SaveSchema(ctx, req.(*Schema))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_SaveSchemaVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchemaVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).SaveSchemaVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_SaveSchemaVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).SaveSchemaVersion(ctx, req.(*SchemaVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetSchemasCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchemasCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetSchemasCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchemasCount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_GetFieldSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetFieldSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_GetFieldSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetFieldSchema(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_MoveSchemaToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).MoveSchemaToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_MoveSchemaToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).MoveSchemaToGroup(ctx, req.(*MoveSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_MarkSchemaAsTrash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrashSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).MarkSchemaAsTrash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SchemaService_MarkSchemaAsTrash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).MarkSchemaAsTrash(ctx, req.(*TrashSchema))
	}
	return interceptor(ctx, in, info, handler)
}

// SchemaService_ServiceDesc is the grpc.ServiceDesc for SchemaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchemaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "schema.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSchema",
			Handler:    _SchemaService_CreateSchema_Handler,
		},
		{
			MethodName: "CreateSchemaVersion",
			Handler:    _SchemaService_CreateSchemaVersion_Handler,
		},
		{
			MethodName: "CreateSchemaWithVersion",
			Handler:    _SchemaService_CreateSchemaWithVersion_Handler,
		},
		{
			MethodName: "CreateVersions",
			Handler:    _SchemaService_CreateVersions_Handler,
		},
		{
			MethodName: "ListSchemas",
			Handler:    _SchemaService_ListSchemas_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _SchemaService_ListVersions_Handler,
		},
		{
			MethodName: "ListSchemasWithVersions",
			Handler:    _SchemaService_ListSchemasWithVersions_Handler,
		},
		{
			MethodName: "ListTrashSchemas",
			Handler:    _SchemaService_ListTrashSchemas_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _SchemaService_GetSchema_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _SchemaService_GetVersion_Handler,
		},
		{
			MethodName: "GetSchemaWithVersions",
			Handler:    _SchemaService_GetSchemaWithVersions_Handler,
		},
		{
			MethodName: "GetSchemaWithVersion",
			Handler:    _SchemaService_GetSchemaWithVersion_Handler,
		},
		{
			MethodName: "GetVersionIdList",
			Handler:    _SchemaService_GetVersionIdList_Handler,
		},
		{
			MethodName: "RemoveSchema",
			Handler:    _SchemaService_RemoveSchema_Handler,
		},
		{
			MethodName: "RemoveSchemaVersion",
			Handler:    _SchemaService_RemoveSchemaVersion_Handler,
		},
		{
			MethodName: "CloneSchema",
			Handler:    _SchemaService_CloneSchema_Handler,
		},
		{
			MethodName: "CloneSchemaVersion",
			Handler:    _SchemaService_CloneSchemaVersion_Handler,
		},
		{
			MethodName: "SaveSchema",
			Handler:    _SchemaService_SaveSchema_Handler,
		},
		{
			MethodName: "SaveSchemaVersion",
			Handler:    _SchemaService_SaveSchemaVersion_Handler,
		},
		{
			MethodName: "GetSchemasCount",
			Handler:    _SchemaService_GetSchemasCount_Handler,
		},
		{
			MethodName: "GetFieldSchema",
			Handler:    _SchemaService_GetFieldSchema_Handler,
		},
		{
			MethodName: "MoveSchemaToGroup",
			Handler:    _SchemaService_MoveSchemaToGroup_Handler,
		},
		{
			MethodName: "MarkSchemaAsTrash",
			Handler:    _SchemaService_MarkSchemaAsTrash_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schema/schema.proto",
}
