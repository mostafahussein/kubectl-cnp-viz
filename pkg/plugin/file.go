package plugin

import (
	"fmt"
	"os"
	"path/filepath"
)

func writeYAML(cnp []byte, outputDir string) error {
	cnpPath := filepath.Join(outputDir, yamlFileName)
	file, err := os.Create(cnpPath)
	if err != nil {
		return fmt.Errorf("failed to create a temporary file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(cnp)
	fmt.Printf("INFO: CiliumNetworkPolicy YAML saved at %s\n", cnpPath)
	return err
}
