package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
	"github.com/epam/edp-perf-operator/v2/pkg/client/perf/mock"
)

func TestCheckConnectionToPerfServeRequest(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	require.NoError(t, perfApi.AddToScheme(scheme))

	tests := []struct {
		name             string
		objects          []runtime.Object
		wantErr          require.ErrorAssertionFunc
		connectedSuccess bool
		connectedErr     error
	}{
		{
			name:             "should be executed successfully",
			objects:          []runtime.Object{&perfApi.PerfServer{}},
			wantErr:          require.NoError,
			connectedSuccess: true,
		},
		{
			name:             "should be executed with error",
			objects:          []runtime.Object{},
			wantErr:          require.Error,
			connectedSuccess: false,
			connectedErr:     fmt.Errorf("failed"),
		},
		{
			name:             "should not be updated",
			objects:          []runtime.Object{},
			wantErr:          require.NoError,
			connectedSuccess: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockPerfClient := new(mock.MockPerfClient)
			mockPerfClient.On("Connected").Return(tt.connectedSuccess, tt.connectedErr)

			checkConnectionToPerf := CheckConnectionToPerf{
				client:     fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build(),
				perfClient: mockPerfClient,
			}

			server := &perfApi.PerfServer{}
			tt.wantErr(t, checkConnectionToPerf.ServeRequest(server))
		})
	}
}
