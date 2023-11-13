/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	ddbctlv1alpha1 "github.com/jittakal/ddbctl-dtp-operator/api/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	jobOwnerKey = ".metadata.controller"
	apiGVStr    = ddbctlv1alpha1.GroupVersion.String()
)

// DeleteTablePartitionDataJobReconciler reconciles a DeleteTablePartitionDataJob object
type DeleteTablePartitionDataJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=batch,resources=jobs,verbs=create;get;list;watch;delete
//+kubebuilder:rbac:groups=ddbctl.operators.jittakal.io,resources=deletetablepartitiondatajobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=ddbctl.operators.jittakal.io,resources=deletetablepartitiondatajobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=ddbctl.operators.jittakal.io,resources=deletetablepartitiondatajobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeleteTablePartitionDataJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.0/pkg/reconcile
func (r *DeleteTablePartitionDataJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("DynamoDB Delete Table Patition Job - Reconcilation", "status", "started")

	// Fetch the DeleteTablePartitionDataJob instance
	var ddbCtlDtpJob ddbctlv1alpha1.DeleteTablePartitionDataJob
	if err := r.Get(ctx, req.NamespacedName, &ddbCtlDtpJob); err != nil {
		log.Error(err, "unable to fetch DeleteTablePartitionDataJob")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Define the Pod Template
	podSpec := corev1.PodSpec{
		Containers: []corev1.Container{
			{
				Name:    "ddbctl-dtp-container",
				Image:   "jittakal/go-dynamodb-partition-delete:latest",
				Command: []string{"/ddbctl", "delete-partition"},
				Args: []string{
					"-t",
					ddbCtlDtpJob.Spec.TableName,
					"-p",
					ddbCtlDtpJob.Spec.PartitionValue,
					"-e",
					ddbCtlDtpJob.Spec.EndpointURL,
					"-r",
					ddbCtlDtpJob.Spec.AWSRegion,
					"-s",
				},
			},
		},
		RestartPolicy: corev1.RestartPolicyNever,
	}

	// Define DeleteTablePartitionDataJob

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ddbCtlDtpJob.Name + "-job",
			Namespace: req.Namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": "ddbctl-dtp"},
				},
				Spec: podSpec,
			},
		},
	}

	// Set DeleteTablePartitionDataJob instance as a the owner of the Job
	if err := ctrl.SetControllerReference(&ddbCtlDtpJob, job, r.Scheme); err != nil {
		log.Error(err, "unable to set controller reference for the Job")
		return ctrl.Result{}, err
	}

	// Create or Update the Job
	if err := r.Create(ctx, job); err != nil {
		log.Error(err, "unable to create Job for DeleteTablePartitionDataJob")
		return ctrl.Result{}, err
	}

	log.Info("DynamoDB Delete Table Patition Job - Reconcilation", "status", "completed")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeleteTablePartitionDataJobReconciler) SetupWithManager(mgr ctrl.Manager) error {

	if err := mgr.GetFieldIndexer().IndexField(context.Background(), &batchv1.Job{}, jobOwnerKey, func(rawObj client.Object) []string {
		// grab the job object, extract the owner...
		job := rawObj.(*batchv1.Job)
		owner := metav1.GetControllerOf(job)
		if owner == nil {
			return nil
		}
		// ...make sure it's a CronJob...
		if owner.APIVersion != apiGVStr || owner.Kind != "DeleteTablePartitionDataJob" {
			return nil
		}

		// ...and if so, return it
		return []string{owner.Name}
	}); err != nil {
		return err
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&ddbctlv1alpha1.DeleteTablePartitionDataJob{}).
		Owns(&batchv1.Job{}).
		Complete(r)
}
