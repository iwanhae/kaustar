package main

import (
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/rs/zerolog/log"

	authenticationv1beta1 "github.com/iwanhae/kaustar/api/v1beta1"
	"github.com/iwanhae/kaustar/internal/controller"
	//+kubebuilder:scaffold:imports
)

func init() {

	utilruntime.Must(authenticationv1beta1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	Execute()
}

var (
	scheme = runtime.NewScheme()
)

func runController() {
	logger := log.Logger.With().Str("phase", "setup").Logger()
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	if err != nil {
		logger.Err(err).Msg("unable to start manager")
		os.Exit(1)
	}

	if err = (&controller.TokenReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		logger.Err(err).Msg("unable to create Token controller")
		os.Exit(1)
	}
	if err = (&controller.UserReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		logger.Err(err).Msg("unable to create User controller")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	logger.Info().Msg("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		logger.Err(err).Msg("problem running manager")
		os.Exit(1)
	}
}
