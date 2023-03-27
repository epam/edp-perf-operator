package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	mocks "github.com/epam/edp-perf-operator/v2/mocks/controllers/perfserver/chain/handler"
	clientMock "github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
)

func TestCreateDefChain(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	scheme.AddKnownTypes(appsv1.SchemeGroupVersion, &perfApi.PerfServer{})

	perfClient := new(clientMock.MockPerfClient)
	client := fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(&perfApi.PerfServer{}).Build()

	want := CheckConnectionToPerf{
		next: PutPerfProject{
			next: PutEdpComponent{
				client: client,
				scheme: scheme,
			},
			perfClient: perfClient,
		},
		client:     client,
		perfClient: perfClient,
	}

	require.Equal(t, want, CreateDefChain(client, scheme, perfClient))
}

func TestNextServeOrNil(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		next    *mocks.PerfServerHandler
		nextErr error
		wantErr require.ErrorAssertionFunc
	}{
		{
			name:    "should finish handling of PerfServer",
			next:    nil,
			wantErr: require.NoError,
		},
		{
			name:    "should fail due to ServeRequest error",
			next:    mocks.NewPerfServerHandler(t),
			nextErr: fmt.Errorf("failed"),
			wantErr: require.Error,
		},
		{
			name:    "should run ServeRequest without errors",
			next:    mocks.NewPerfServerHandler(t),
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			perfServer := &perfApi.PerfServer{}

			if tt.next == nil {
				tt.wantErr(t, nextServeOrNil(nil, perfServer))
				return
			}

			tt.next.On("ServeRequest", perfServer).Return(tt.nextErr)
			tt.wantErr(t, nextServeOrNil(tt.next, perfServer))
		})
	}
}
