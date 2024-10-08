package conditions

import (
	machinev1 "github.com/openshift/api/machine/v1beta1"
)

type MachineWrapper struct {
	*machinev1.Machine
}

func (m *MachineWrapper) GetConditions() []machinev1.Condition {
	return m.Status.Conditions
}

func (m *MachineWrapper) SetConditions(conditions []machinev1.Condition) {
	m.Status.Conditions = conditions
}

type MachineHealthCheckWrapper struct {
	*machinev1.MachineHealthCheck
}

func (m *MachineHealthCheckWrapper) GetConditions() []machinev1.Condition {
	return m.Status.Conditions
}

func (m *MachineHealthCheckWrapper) SetConditions(conditions []machinev1.Condition) {
	m.Status.Conditions = conditions
}

type MachineSetWrapper struct {
	*machinev1.MachineSet
}

func (m *MachineSetWrapper) GetConditions() []machinev1.Condition {
	return m.Status.Conditions
}

func (m *MachineSetWrapper) SetConditions(conditions []machinev1.Condition) {
	m.Status.Conditions = conditions
}
