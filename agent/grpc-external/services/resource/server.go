package resource

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	"github.com/hashicorp/consul/acl"
	"github.com/hashicorp/consul/acl/resolver"
	"github.com/hashicorp/consul/internal/resource"
	"github.com/hashicorp/consul/internal/storage"
	"github.com/hashicorp/consul/proto-public/pbresource"
)

type Server struct {
	Config
}

type Config struct {
	Logger      hclog.Logger
	Registry    Registry
	Backend     Backend
	ACLResolver ACLResolver
}

//go:generate mockery --name Registry --inpackage
type Registry interface {
	resource.Registry
}

//go:generate mockery --name Backend --inpackage
type Backend interface {
	storage.Backend
}

//go:generate mockery --name ACLResolver --inpackage
type ACLResolver interface {
	ResolveTokenAndDefaultMeta(string, *acl.EnterpriseMeta, *acl.AuthorizerContext) (resolver.Result, error)
}

func NewServer(cfg Config) *Server {
	return &Server{cfg}
}

var _ pbresource.ResourceServiceServer = (*Server)(nil)

func (s *Server) Register(grpcServer *grpc.Server) {
	pbresource.RegisterResourceServiceServer(grpcServer, s)
}

func (s *Server) WriteStatus(ctx context.Context, req *pbresource.WriteStatusRequest) (*pbresource.WriteStatusResponse, error) {
	// TODO
	return &pbresource.WriteStatusResponse{}, nil
}

func (s *Server) Delete(ctx context.Context, req *pbresource.DeleteRequest) (*pbresource.DeleteResponse, error) {
	// TODO
	return &pbresource.DeleteResponse{}, nil
}

// Extracts ACL token ID from incoming context
func tokenFromContext(ctx context.Context) string {
	// TODO(spatel): Figure out where to get the token from correctly. This is just a stand in for the PR
	//               Should I be using QueryOptions like some of the other gRPC servers?
	token := ctx.Value("x-token")
	if token == nil {
		return acl.AnonymousTokenID
	}
	return token.(string)
}

func (s *Server) resolveType(typ *pbresource.Type) (*resource.Registration, error) {
	v, ok := s.Registry.Resolve(typ)
	if ok {
		return &v, nil
	}
	return nil, status.Errorf(
		codes.InvalidArgument,
		"resource type %s not registered", resource.ToGVK(typ),
	)
}

func readConsistencyFrom(ctx context.Context) storage.ReadConsistency {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return storage.EventualConsistency
	}

	vals := md.Get("x-consul-consistency-mode")
	if len(vals) == 0 {
		return storage.EventualConsistency
	}

	if vals[0] == "consistent" {
		return storage.StrongConsistency
	}
	return storage.EventualConsistency
}

func clone[T proto.Message](v T) T { return proto.Clone(v).(T) }
