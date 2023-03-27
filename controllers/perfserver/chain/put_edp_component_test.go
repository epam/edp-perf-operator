package chain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	componentApi "github.com/epam/edp-component-operator/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

const (
	fakeNamespace = "fake-namespace"
	fakeIconPath  = "icon path"
)

func TestServeRequest(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	require.NoError(t, componentApi.AddToScheme(scheme))
	require.NoError(t, perfApi.AddToScheme(scheme))
	require.NoError(t, metav1.AddMetaToScheme(scheme))

	tests := []struct {
		name    string
		scheme  *runtime.Scheme
		objects []runtime.Object
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "should skip creating EDP component since it already exists",
			objects: []runtime.Object{&componentApi.EDPComponent{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
			}},
			scheme:  scheme,
			wantErr: require.NoError,
		},
		{
			name:    "should fail to get client, because EDP component is not in the scheme",
			scheme:  runtime.NewScheme(),
			wantErr: require.Error,
		},
		{
			name: "should fail, because icon doesn't exist in hardcoded path",
			objects: []runtime.Object{&componentApi.EDPComponent{}, &perfApi.PerfServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
			}},
			scheme:  scheme,
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			putEdpComponent := PutEdpComponent{
				client: fake.NewClientBuilder().WithScheme(tt.scheme).WithRuntimeObjects(tt.objects...).Build(),
				scheme: tt.scheme,
			}

			server := &perfApi.PerfServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
			}

			tt.wantErr(t, putEdpComponent.ServeRequest(server))
		})
	}
}

func TestPutEdpComponentCreateEdpComponent(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	require.NoError(t, componentApi.AddToScheme(scheme))
	require.NoError(t, perfApi.AddToScheme(scheme))

	tests := []struct {
		name     string
		iconPath string
		wantErr  require.ErrorAssertionFunc
	}{
		{
			name:     "should create EdpComponent",
			iconPath: fakeIconPath,
			wantErr:  require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			perfServer := &perfApi.PerfServer{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
			}

			putEdpComponent := PutEdpComponent{
				client: fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(perfServer).Build(),
				scheme: scheme,
			}
			tt.wantErr(t, putEdpComponent.createEdpComponent(perfServer, tt.iconPath, func(s string) ([]byte, error) {
				return []byte{}, nil
			}))
		})
	}
}

func TestGetIcon(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		iconPath              string
		fileReaderFuncContent []byte
		fileReaderFuncErr     error
		want                  *string
		wantErr               require.ErrorAssertionFunc
	}{
		{
			name:              "should return error, because we cannot open the file",
			iconPath:          fakeIconPath,
			fileReaderFuncErr: fmt.Errorf("failed"),
			want:              nil,
			wantErr:           require.Error,
		},
		{
			name:                  "should return empty string, because file is empty",
			iconPath:              fakeIconPath,
			fileReaderFuncContent: []byte{},
			want:                  pointer.String(""),
			wantErr:               require.NoError,
		},
		{
			name:                  "should return value corresponding to file content",
			iconPath:              fakeIconPath,
			fileReaderFuncContent: []byte("2710"),
			want:                  pointer.String("MjcxMA=="),
			wantErr:               require.NoError,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			fileReaderFunc := func(path string) ([]byte, error) {
				return tt.fileReaderFuncContent, tt.fileReaderFuncErr
			}

			got, err := getIcon(tt.iconPath, fileReaderFunc)
			tt.wantErr(t, err)

			require.Equal(t, got, tt.want)
		})
	}
}
