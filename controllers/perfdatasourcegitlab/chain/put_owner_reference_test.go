package chain

import (
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	perfApi "github.com/epam/edp-perf-operator/v2/api/v1"
)

const (
	fakeName      = "fake-name"
	fakeNamespace = "fake-namespace"
)

func TestPutOwnerReference_ServeRequest(t *testing.T) {
	t.Parallel()

	scheme := runtime.NewScheme()

	require.NoError(t, perfApi.AddToScheme(scheme))
	require.NoError(t, codebaseApi.AddToScheme(scheme))

	tests := []struct {
		name           string
		perfDataSource *perfApi.PerfDataSourceGitLab
		objects        []runtime.Object
		wantErr        require.ErrorAssertionFunc
	}{
		{
			name: "should not set owner reference, because perf data source contains perf server owner reference",
			perfDataSource: &perfApi.PerfDataSourceGitLab{
				ObjectMeta: metav1.ObjectMeta{
					OwnerReferences: []metav1.OwnerReference{
						{
							Kind: "Codebase",
						},
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "should set owner reference",
			perfDataSource: &perfApi.PerfDataSourceGitLab{
				ObjectMeta: metav1.ObjectMeta{
					Name:      fakeName,
					Namespace: fakeNamespace,
				},
				Spec: perfApi.PerfDataSourceGitLabSpec{
					CodebaseName:   fakeName,
					PerfServerName: fakeName,
				},
			},
			objects: []runtime.Object{
				&codebaseApi.Codebase{
					ObjectMeta: metav1.ObjectMeta{
						Name:      fakeName,
						Namespace: fakeNamespace,
					},
				},
			},
			wantErr: require.NoError,
		},
		{
			name: "should not find perf server",
			perfDataSource: &perfApi.PerfDataSourceGitLab{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: fakeNamespace,
				},
				Spec: perfApi.PerfDataSourceGitLabSpec{
					PerfServerName: fakeName,
				},
			},
			objects: []runtime.Object{
				&perfApi.PerfServer{},
			},
			wantErr: require.Error,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.objects = append(tt.objects, tt.perfDataSource)

			putOwnerReference := PutOwnerReference{
				scheme: scheme,
				client: fake.NewClientBuilder().WithScheme(scheme).WithRuntimeObjects(tt.objects...).Build(),
			}

			tt.wantErr(t, putOwnerReference.ServeRequest(tt.perfDataSource))
		})
	}
}
