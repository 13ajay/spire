package bootstrap

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	bootstrapv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/bootstrap/v1"
	"github.com/spiffe/spire/pkg/agent/api/rpccontext"
	"github.com/spiffe/spire/pkg/server/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Config defines the bundle service configuration.
type Config struct {
	TrustDomain       spiffeid.TrustDomain
}

// Service defines the v1 bundle service properties.
type Service struct {
	bootstrapv1.UnsafeBootstrapServer

	td spiffeid.TrustDomain
}

// New creates a new bundle service.
func New(config Config) *Service {
	return &Service{
		td: config.TrustDomain,
	}
}

// RegisterService registers the bundle service on the gRPC server.
func RegisterService(s grpc.ServiceRegistrar, service *Service) {
	bootstrapv1.RegisterBootstrapServer(s, service)
}

// CountBundles returns the total number of bundles.
func (s *Service) GetTrustAnchorARN(ctx context.Context, _ *bootstrapv1.GetTrustAnchorARNRequest) (*bootstrapv1.GetTrustAnchorARNResponse, error) {
    log := rpccontext.Logger(ctx)
    api.MakeStatus(log, codes.OK, "Begin GetTrustAnchorARN", nil)
    return &bootstrapv1.GetTrustAnchorARNResponse{TrustAnchorArn: "TEST TRUST ANCHOR ARN"}, nil
}
