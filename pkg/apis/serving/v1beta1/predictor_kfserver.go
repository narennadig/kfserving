package v1beta1

import (
	"fmt"
	"github.com/kubeflow/kfserving/pkg/constants"
	v1 "k8s.io/api/core/v1"
)

// KFServerSpec defines arguments for configuring KFServer model serving.
type KFServerSpec struct {
	// Contains fields shared across all predictors
	PredictorExtensionSpec `json:",inline"`
}

// Validate returns an error if invalid
func (k *KFServerSpec) Validate() error {
	return nil
}

// Default sets defaults on the resource
func (k *KFServerSpec) Default() {}

// GetContainers transforms the resource into a container spec
func (k *KFServerSpec) GetContainer(modelName string, config *InferenceServicesConfig) *v1.Container {
	arguments := []string{
		fmt.Sprintf("%s=%s", constants.ArgumentModelName, modelName),
		fmt.Sprintf("%s=%s", constants.ArgumentModelDir, constants.DefaultModelLocalMountPath),
		fmt.Sprintf("%s=%s", constants.ArgumentHttpPort, constants.InferenceServiceDefaultHttpPort),
	}
	/*if parallelism != 0 {
		arguments = append(arguments, fmt.Sprintf("%s=%s", constants.ArgumentWorkers, strconv.Itoa(parallelism)))
	}*/
	k.Image = config.Predictors.SKlearn.ContainerImage + ":" + k.RuntimeVersion
	k.Name = constants.InferenceServiceContainerName
	k.Args = arguments
	return &v1.Container{
		Name:  k.Name,
		Image: k.Image,
		Args:  k.Args,
	}
}
