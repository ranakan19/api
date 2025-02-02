package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// These are valid conditions of a Notification

	// NotificationDeletionError indicates that the notification failed to be deleted
	NotificationDeletionError ConditionType = deletionError

	// NotificationSent reflects whether the notification has been sent to the user
	NotificationSent ConditionType = "Sent"

	// Status condition reasons
	NotificationSentReason          = "Sent"
	NotificationDeletionErrorReason = "UnableToDeleteNotification"
	NotificationContextErrorReason  = "NotificationContextError"
	NotificationDeliveryErrorReason = "DeliveryError"

	// NotificationUserNameLabelKey is used to identify the user that the notification belongs to
	NotificationUserNameLabelKey = LabelKeyPrefix + "username"

	// NotificationTypeLabelKey is used to identify the notification type, for example: deactivated
	NotificationTypeLabelKey = LabelKeyPrefix + "type"

	// Notification Types which describe the type of notification being sent
	NotificationTypeDeactivating = "deactivating"
	NotificationTypeDeactivated  = "deactivated"
	NotificationTypeProvisioned  = "provisioned"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NotificationSpec defines the desired state of Notification
// +k8s:openapi-gen=true
type NotificationSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// UserID is the user ID from RHD Identity Provider token (“sub” claim).  The UserID is used by
	// the notification service (i.e. the NotificationController) to lookup the UserSignup resource for the user,
	// and extract from it the values required to generate the notification content and to deliver the notification
	UserID string `json:"userID,omitempty"`

	// Recipient may be used as an alternative to UserID to specify an email address where the notification will be delivered.
	Recipient string `json:"recipient,omitempty"`

	// Template is the name of the NotificationTemplate resource that will be used to generate the notification
	Template string `json:"template,omitempty"`

	// Subject is used when no template value is specified, in cases where the complete notification subject is
	// specified at notification creation time
	Subject string `json:"subject,omitempty"`

	// Content is used when no template value is specified, in cases where the complete notification content is
	// specified at notification creation time
	Content string `json:"content,omitempty"`
}

// NotificationStatus defines the observed state of Notification
// +k8s:openapi-gen=true
type NotificationStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Conditions is an array of current Notification conditions
	// Supported condition types:
	// Sent
	// +optional
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Notification registers a notification in the CodeReady Toolchain
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
// +kubebuilder:printcolumn:name="User ID",type="string",JSONPath=`.spec.userID`,priority=1
// +kubebuilder:printcolumn:name="Sent",type="string",JSONPath=`.status.conditions[?(@.type=="Sent")].status`
// +kubebuilder:validation:XPreserveUnknownFields
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Notification"
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NotificationSpec   `json:"spec,omitempty"`
	Status NotificationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotificationList contains a list of Notification
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
