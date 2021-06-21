/*


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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GmtSpec defines the desired state of Gmt
type GmtSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Gmt. Edit Gmt_types.go to remove/update
	//Foo string `json:"foo,omitempty"`

	UpdateInterval int64 `json:"updateInterval,omitempty"`
}

// GmtStatus defines the observed state of Gmt
type GmtStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CardList       CardList     `json:"cardList,omitempty"`
	CardNumber     uint         `json:"cardNumber,omitempty"`
	UpdateTime     *metav1.Time `json:"updateTime,omitempty"`
	TotalMemorySum uint64       `json:"totalMemorySum,omitempty"`
	FreeMemorySum  uint64       `json:"freeMemorySum,omitempty"`
}

// CardList implements sort.Interface for []Card based on
// the ID field.
type CardList []Card

func (a CardList) Len() int           { return len(a) }
func (a CardList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a CardList) Less(i, j int) bool { return a[i].ID < a[j].ID }


type Card struct {
	ID          uint   `json:"id"`
	Health      string `json:"health,omitempty"`
	Model       string `json:"model,omitempty"`
	Power       uint   `json:"power,omitempty"`
	Core        uint   `json:"core,omitempty"`
	Clock       uint   `json:"clock,omitempty"`
	TotalMemory uint64 `json:"totalMemory,omitempty"`
	FreeMemory  uint64 `json:"freeMemory,omitempty"`
	GPUUtil     uint   `json:"gpuUtil,omitempty"`
	Bandwidth   uint   `json:"bandwidth,omitempty"`
	Temperature uint   `json:"temperature,omitempty"`
}

// +kubebuilder:object:root=true

// Add:
// +kubebuilder:resource:scope=

// Gmt is the Schema for the gmts API
type Gmt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GmtSpec   `json:"spec,omitempty"`
	Status GmtStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GmtList contains a list of Gmt
type GmtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gmt `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gmt{}, &GmtList{})
}
