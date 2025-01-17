package nais_io_v1

import (
	"time"

	"github.com/nais/liberator/pkg/events"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Status contains different NAIS status properties
type Status struct {
	SynchronizationTime     int64               `json:"synchronizationTime,omitempty"`
	RolloutCompleteTime     int64               `json:"rolloutCompleteTime,omitempty"`
	CorrelationID           string              `json:"correlationID,omitempty"`
	DeploymentRolloutStatus string              `json:"deploymentRolloutStatus,omitempty"`
	SynchronizationState    string              `json:"synchronizationState,omitempty"`
	SynchronizationHash     string              `json:"synchronizationHash,omitempty"`
	Conditions              *[]metav1.Condition `json:"conditions,omitempty"`
}

func (in *Status) SetStatusConditions() {

	if in.Conditions == nil {
		in.Conditions = &[]metav1.Condition{}
	}

	reconcilingConditionStatus := metav1.ConditionFalse
	if in.SynchronizationState != events.RolloutComplete && in.SynchronizationState != events.FailedSynchronization {
		reconcilingConditionStatus = metav1.ConditionTrue
	}

	readyConditionStatus := metav1.ConditionFalse
	if in.SynchronizationState == events.RolloutComplete {
		readyConditionStatus = metav1.ConditionTrue
	}

	stalledConditionStatus := metav1.ConditionFalse
	if in.SynchronizationState == events.FailedSynchronization {
		stalledConditionStatus = metav1.ConditionTrue
	}

	in.addStatusCondition("Ready", readyConditionStatus)
	in.addStatusCondition("Stalled", stalledConditionStatus)
	in.addStatusCondition("Reconciling", reconcilingConditionStatus)
}

func (in *Status) addStatusCondition(name string, conditionStatus metav1.ConditionStatus) {
	condition := meta.FindStatusCondition(*in.Conditions, name)
	if condition == nil {
		condition = &metav1.Condition{
			Type:               name,
			Status:             conditionStatus,
			LastTransitionTime: metav1.NewTime(time.Now()),
			Reason:             in.SynchronizationState,
			Message:            in.DeploymentRolloutStatus,
		}
	} else {
		condition.Status = conditionStatus
		condition.Reason = in.SynchronizationState
		condition.Message = in.DeploymentRolloutStatus
	}

	meta.SetStatusCondition(in.Conditions, *condition)
}
