// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type nodeJsBuildOptions struct {
	Install            bool     `json:"install,omitempty"`
	RunScripts         []string `json:"runScripts,omitempty"`
	DefaultNpmRegistry string   `json:"defaultNpmRegistry,omitempty"`
	SapNpmRegistry     string   `json:"sapNpmRegistry,omitempty"`
}

// NodeJsBuildCommand Build a JavaScript backend (node js) project
func NodeJsBuildCommand() *cobra.Command {
	metadata := nodeJsBuildMetadata()
	var stepConfig nodeJsBuildOptions
	var startTime time.Time

	var createNodeJsBuildCmd = &cobra.Command{
		Use:   "nodeJsBuild",
		Short: "Build a JavaScript backend (node js) project",
		Long:  `todo`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName("nodeJsBuild")
			log.SetVerbose(GeneralConfig.Verbose)
			return PrepareConfig(cmd, &metadata, "nodeJsBuild", &stepConfig, config.OpenPiperFile)
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, "nodeJsBuild")
			nodeJsBuild(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
		},
	}

	addNodeJsBuildFlags(createNodeJsBuildCmd, &stepConfig)
	return createNodeJsBuildCmd
}

func addNodeJsBuildFlags(cmd *cobra.Command, stepConfig *nodeJsBuildOptions) {
	cmd.Flags().BoolVar(&stepConfig.Install, "install", false, "Run npm install or similar commands depending on the project structure.")
	cmd.Flags().StringSliceVar(&stepConfig.RunScripts, "runScripts", []string{}, "List of additinal run scripts to execute from package.json.")
	cmd.Flags().StringVar(&stepConfig.DefaultNpmRegistry, "defaultNpmRegistry", os.Getenv("PIPER_defaultNpmRegistry"), "URL of the npm registry to use. Defaults to https://registry.npmjs.org/")
	cmd.Flags().StringVar(&stepConfig.SapNpmRegistry, "sapNpmRegistry", "https://npm.sap.com", "The default npm registry url to be used as the remote mirror for the SAP npm packages.")

}

// retrieve step metadata
func nodeJsBuildMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "nodeJsBuild",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "install",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "runScripts",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "[]string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "defaultNpmRegistry",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "sapNpmRegistry",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
