package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	//+kubebuilder:scaffold:imports
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	codebaseApi "github.com/epam/edp-codebase-operator/v2/api/v1"
	buildInfo "github.com/epam/edp-common/pkg/config"
	edpCompApi "github.com/epam/edp-component-operator/api/v1"
	perfApiV1 "github.com/epam/edp-perf-operator/v2/api/v1"
	perfApiV1Alpha "github.com/epam/edp-perf-operator/v2/api/v1alpha1"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcegitlab"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcejenkins"
	"github.com/epam/edp-perf-operator/v2/controllers/perfdatasourcesonar"
	"github.com/epam/edp-perf-operator/v2/controllers/perfserver"
	"github.com/epam/edp-perf-operator/v2/pkg/util/cluster"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

const (
	perfOperatorLock               = "edp-perf-operator-lock"
	errMsgUnableToCreateController = "unable to create controller"
	keyAndValueController          = "controller"
	managerPort                    = 9443
)

func main() {
	var (
		metricsAddr          string
		enableLeaderElection bool
		probeAddr            string
	)

	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", cluster.RunningInCluster(),
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")

	mode, err := cluster.GetDebugMode()
	if err != nil {
		setupLog.Error(err, "unable to get debug mode value")
		os.Exit(1)
	}

	opts := zap.Options{
		Development: mode,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(perfApiV1Alpha.AddToScheme(scheme))
	utilruntime.Must(perfApiV1.AddToScheme(scheme))
	utilruntime.Must(codebaseApi.AddToScheme(scheme))
	utilruntime.Must(edpCompApi.AddToScheme(scheme))

	v := buildInfo.Get()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	setupLog.Info("Starting the Perf Operator",
		"version", v.Version,
		"git-commit", v.GitCommit,
		"git-tag", v.GitTag,
		"build-date", v.BuildDate,
		"go-version", v.Go,
		"go-client", v.KubectlVersion,
		"platform", v.Platform,
	)

	ns, err := cluster.GetWatchNamespace()
	if err != nil {
		setupLog.Error(err, "unable to get watch namespace")
		os.Exit(1)
	}

	cfg := ctrl.GetConfigOrDie()

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		HealthProbeBindAddress: probeAddr,
		Port:                   managerPort,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       perfOperatorLock,
		MapperProvider: func(c *rest.Config) (meta.RESTMapper, error) {
			return apiutil.NewDynamicRESTMapper(cfg)
		},
		Namespace: ns,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	ctrlLog := ctrl.Log.WithName("controllers")

	pdsgCtrl := perfdatasourcegitlab.NewReconcilePerfDataSourceGitLab(mgr.GetClient(), mgr.GetScheme(), ctrlLog)
	if err := pdsgCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, errMsgUnableToCreateController, keyAndValueController, "perf-data-source-gitlab")
		os.Exit(1)
	}

	pdsjCtrl := perfdatasourcejenkins.NewReconcilePerfDataSourceJenkins(mgr.GetClient(), mgr.GetScheme(), ctrlLog)
	if err := pdsjCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, errMsgUnableToCreateController, keyAndValueController, "perf-data-source-jenkins")
		os.Exit(1)
	}

	pdssCtrl := perfdatasourcesonar.NewReconcilePerfDataSourceSonar(mgr.GetClient(), mgr.GetScheme(), ctrlLog)
	if err := pdssCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, errMsgUnableToCreateController, keyAndValueController, "perf-data-source-sonar")
		os.Exit(1)
	}

	psCtrl := perfserver.NewReconcilePerfServer(mgr.GetClient(), mgr.GetScheme(), ctrlLog)
	if err := psCtrl.SetupWithManager(mgr); err != nil {
		setupLog.Error(err, errMsgUnableToCreateController, keyAndValueController, "perf-server")
		os.Exit(1)
	}

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}

	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	setupLog.Info("starting manager")

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
