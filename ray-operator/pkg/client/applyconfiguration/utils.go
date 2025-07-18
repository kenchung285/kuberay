// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	internal "github.com/ray-project/kuberay/ray-operator/pkg/client/applyconfiguration/internal"
	rayv1 "github.com/ray-project/kuberay/ray-operator/pkg/client/applyconfiguration/ray/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=ray.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("AppStatus"):
		return &rayv1.AppStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("AutoscalerOptions"):
		return &rayv1.AutoscalerOptionsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeletionPolicy"):
		return &rayv1.DeletionPolicyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("DeletionStrategy"):
		return &rayv1.DeletionStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("GcsFaultToleranceOptions"):
		return &rayv1.GcsFaultToleranceOptionsApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HeadGroupSpec"):
		return &rayv1.HeadGroupSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("HeadInfo"):
		return &rayv1.HeadInfoApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayCluster"):
		return &rayv1.RayClusterApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayClusterSpec"):
		return &rayv1.RayClusterSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayClusterStatus"):
		return &rayv1.RayClusterStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayJob"):
		return &rayv1.RayJobApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayJobSpec"):
		return &rayv1.RayJobSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayJobStatus"):
		return &rayv1.RayJobStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayJobStatusInfo"):
		return &rayv1.RayJobStatusInfoApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayService"):
		return &rayv1.RayServiceApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayServiceSpec"):
		return &rayv1.RayServiceSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayServiceStatus"):
		return &rayv1.RayServiceStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayServiceStatuses"):
		return &rayv1.RayServiceStatusesApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RayServiceUpgradeStrategy"):
		return &rayv1.RayServiceUpgradeStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("RedisCredential"):
		return &rayv1.RedisCredentialApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ScaleStrategy"):
		return &rayv1.ScaleStrategyApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ServeDeploymentStatus"):
		return &rayv1.ServeDeploymentStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubmitterConfig"):
		return &rayv1.SubmitterConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("WorkerGroupSpec"):
		return &rayv1.WorkerGroupSpecApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
