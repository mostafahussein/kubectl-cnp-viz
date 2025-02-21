package plugin

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
)

var (
	namespace         = "default"
	fileName          = "ciliumnetworkpolicy"
	timestamp         = time.Now().Unix()
	yamlFileName      = fmt.Sprintf("%s_%d.yaml", fileName, timestamp)
	diagramFileName   = fmt.Sprintf("%s_%d.png", fileName, timestamp)
	networkpolicySite = "https://editor.networkpolicy.io/"
)

type DiagramCaptureOptions struct {
	Scale      float64
	TranslateX float64
	TranslateY float64
	OutputDir  string
}

func RunPlugin(cnpName string, configFlags *genericclioptions.ConfigFlags, opts DiagramCaptureOptions) error {
	if configFlags.Namespace != nil {
		ns, _, err := configFlags.ToRawKubeConfigLoader().Namespace()
		if err == nil {
			namespace = ns
		}
	}

	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return fmt.Errorf("failed to read the kubeconfig file: %w", err)
	}

	clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create the Kubernetes clientset: %w", err)
	}

	cnpGVR := schema.GroupVersionResource{
		Group:    "cilium.io",
		Version:  "v2",
		Resource: "ciliumnetworkpolicies",
	}

	cnp, err := clientset.Resource(cnpGVR).Namespace(namespace).Get(context.TODO(), cnpName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to fetch the CiliumNetworkPolicy: %w", err)
	}

	yamlFileName = fmt.Sprintf("%s_%s", cnpName, yamlFileName)
	diagramFileName = fmt.Sprintf("%s_%s", cnpName, diagramFileName)

	yamlBytes, err := yaml.Marshal(cnp.UnstructuredContent())
	if err != nil {
		return fmt.Errorf("failed to marshal the YAML data: %w", err)
	}

	if err := writeYAML(yamlBytes, opts.OutputDir); err != nil {
		return fmt.Errorf("failed to write the YAML file: %w", err)
	}

	if err := captureDiagram(opts.TranslateX, opts.TranslateY, opts.Scale, opts.OutputDir); err != nil {
		return fmt.Errorf("failed to capture the diagram: %w", err)
	}

	return nil
}
