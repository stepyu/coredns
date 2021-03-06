package object

import (
	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// Service is a stripped down api.Service with only the items we need for CoreDNS.
type Service struct {
	Version      string
	Name         string
	Namespace    string
	Index        string
	ClusterIP    string
	Type         api.ServiceType
	ExternalName string
	Ports        []api.ServicePort

	*Empty
}

// ServiceKey return a string using for the index.
func ServiceKey(namespace, name string) string { return namespace + "/" + name }

// ToService converts an api.Service to a *Service.
func ToService(obj interface{}) interface{} {
	svc, ok := obj.(*api.Service)
	if !ok {
		return nil
	}

	s := &Service{
		Version:      svc.GetResourceVersion(),
		Name:         svc.GetName(),
		Namespace:    svc.GetNamespace(),
		Index:        ServiceKey(svc.GetName(), svc.GetNamespace()),
		ClusterIP:    svc.Spec.ClusterIP,
		Type:         svc.Spec.Type,
		ExternalName: svc.Spec.ExternalName,
	}

	if len(svc.Spec.Ports) == 0 {
		// Add sentinal if there are no ports.
		s.Ports = []api.ServicePort{{Port: -1}}
	} else {
		s.Ports = make([]api.ServicePort, len(svc.Spec.Ports))
		copy(s.Ports, svc.Spec.Ports)
	}

	*svc = api.Service{}

	return s
}

var _ runtime.Object = &Service{}

// DeepCopyObject implements the ObjectKind interface.
func (s *Service) DeepCopyObject() runtime.Object {
	s1 := &Service{
		Version:      s.Version,
		Name:         s.Name,
		Namespace:    s.Namespace,
		Index:        s.Index,
		ClusterIP:    s.ClusterIP,
		Type:         s.Type,
		ExternalName: s.ExternalName,
		Ports:        make([]api.ServicePort, len(s.Ports)),
	}
	copy(s1.Ports, s.Ports)
	return s1
}

// GetNamespace implements the metav1.Object interface.
func (s *Service) GetNamespace() string { return s.Namespace }

// SetNamespace implements the metav1.Object interface.
func (s *Service) SetNamespace(namespace string) {}

// GetName implements the metav1.Object interface.
func (s *Service) GetName() string { return s.Name }

// SetName implements the metav1.Object interface.
func (s *Service) SetName(name string) {}

// GetResourceVersion implements the metav1.Object interface.
func (s *Service) GetResourceVersion() string { return s.Version }

// SetResourceVersion implements the metav1.Object interface.
func (s *Service) SetResourceVersion(version string) {}
