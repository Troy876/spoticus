package commands

import (
	"fmt"
	"log"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const launchUsage = "" +
	"📦 *Launch Command — Detailed Usage*\n\n" +
	"This command provisions a new cluster using a specified platform and resource tier.\n\n" +
	"🔧 *Syntax*:\n" +
	"```\n" +
	"launch <cluster_type> <size>\n" +
	"```\n\n" +
	"🧪 *Examples*:\n" +
	"```\n" +
	"launch k8s large\n" +
	"launch openshift medium\n" +
	"```\n\n" +
	"🧱 *Supported Cluster Types*:\n" +
	"• `k8s` — Standard upstream Kubernetes cluster\n" +
	"• `openshift` — Red Hat OpenShift Container Platform\n" +
	"_Only these values are accepted. Input is case-insensitive._\n\n" +
	"📐 *Supported Sizes*:\n" +
	"• `medium` — 8 CPUs / 32 GB RAM\n" +
	"• `large` — 16 CPUs / 64 GB RAM\n" +
	"• `xlarge` — 32 CPUs / 128 GB RAM\n\n" +
	"💰 *⚡ Spot Instances (Cost Optimization)*:\n" +
	"All clusters are provisioned using **cloud spot instances** for maximum cost-efficiency.\n"

// supportedClusterTypes defines the valid cluster types that can be launched.
// Currently supports Kubernetes and OpenShift.
// The keys are the cluster type names, and the values are empty structs
// to allow for efficient O(1) existence checks.
// TODO!: Check ROSA and Karpenter support in the future.
var supportedClusterTypes = map[string]struct{}{
	"k8s":       {},
	"openshift": {},
}

// SizeSpec defines the resource specifications for a given cluster size.
// This includes the number of CPUs and the amount of RAM.
type SizeSpec struct {
	CPU string
	RAM string
}

// supportedSizes defines the available cluster sizes and their specifications.
// Each entry includes the size label (e.g., "large") and its corresponding resources.
var supportedSizes = map[string]SizeSpec{
	"medium": {
		CPU: "8 CPUs",
		RAM: "32 GB RAM",
	},
	"large": {
		CPU: "16 CPUs",
		RAM: "64 GB RAM",
	},
	"xlarge": {
		CPU: "32 CPUs",
		RAM: "128 GB RAM",
	},
}

// HandleLaunch is the main entry point for the "launch" Slack command.
//
// It expects exactly two arguments:
//  1. cluster type — currently one of: "k8s", "openshift"
//  2. cluster size — currently one of: "medium", "large", "xlarge"
//
// If the command is malformed, the user will receive contextual error feedback.
// Otherwise, a confirmation message is sent to the channel describing the requested launch.
//
// The function logs the action for auditing/debugging and ensures the user
// receives structured output with specs.
func HandleLaunch(api *slack.Client, event *slackevents.MessageEvent, args []string) {
	if len(args) < 2 {
		respondError(api, event.Channel, "❌ Missing arguments.\n\n"+launchUsage)
		return
	}

	clusterType := strings.ToLower(args[0])
	size := strings.ToLower(args[1])

	// Validate cluster type
	if !isSupportedClusterType(clusterType) {
		respondError(api, event.Channel,
			fmt.Sprintf("❌ Unsupported cluster type: *%s*\nSupported types: `k8s`, `openshift`", clusterType))
		return
	}

	spec, ok := supportedSizes[size]
	if !ok {
		respondError(api, event.Channel,
			fmt.Sprintf("❌ Invalid size: *%s*\nValid sizes:\n%s", size, formatSupportedSizes()))
		return
	}

	log.Printf("Launching cluster: user=%s type=%s size=%s", event.User, clusterType, size)

	// Compose confirmation message with detailed spec
	message := fmt.Sprintf(
		"🚀 Launching a *%s* cluster of size *%s* for <@%s>\n• CPU: %s\n• Memory: %s",
		clusterType, size, event.User, spec.CPU, spec.RAM)

	// Post the result back to Slack
	if _, _, err := api.PostMessage(event.Channel, slack.MsgOptionText(message, false)); err != nil {
		log.Printf("Error posting launch message: %v", err)
	}
}

// isSupportedClusterType checks if the provided cluster type is one of the supported ones.
// It performs a case-insensitive lookup in the predefined supportedClusterTypes set.
func isSupportedClusterType(t string) bool {
	_, ok := supportedClusterTypes[t]
	return ok
}

// formatSupportedSizes constructs a Slack-friendly bullet list of valid cluster sizes and their specs.
// This is used in error messages to inform the user of acceptable input values.
func formatSupportedSizes() string {
	var b strings.Builder
	for name, spec := range supportedSizes {
		b.WriteString(fmt.Sprintf("• `%s`: %s, %s\n", name, spec.CPU, spec.RAM))
	}
	return b.String()
}

// respondError sends a standardized error message to the given Slack channel.
//
// This is used to provide consistent and visible feedback to the user
// when the input is invalid, missing, or unsupported.
// It logs any failures during Slack message delivery.
func respondError(api *slack.Client, channel, text string) {
	if _, _, err := api.PostMessage(channel, slack.MsgOptionText(text, false)); err != nil {
		log.Printf("Slack error response failed: %v", err)
	}
}
