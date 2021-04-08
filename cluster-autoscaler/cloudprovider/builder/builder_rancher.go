// +build rancher

package builder

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/rancher"
	"k8s.io/autoscaler/cluster-autoscaler/config"
)

// AvailableCloudProviders supported by the rancher provider builder.
var AvailableCloudProviders = []string{
	cloudprovider.RancherProviderName,
}

// DefaultCloudProvider for do-only build is rancher.
const DefaultCloudProvider = cloudprovider.RancherProviderName

func buildCloudProvider(opts config.AutoscalingOptions, do cloudprovider.NodeGroupDiscoveryOptions, rl *cloudprovider.ResourceLimiter) cloudprovider.CloudProvider {
	switch opts.CloudProviderName {
	case cloudprovider.RancherProviderName:
		return rancher.BuildRancher(opts, do, rl)
	}

	return nil
}
