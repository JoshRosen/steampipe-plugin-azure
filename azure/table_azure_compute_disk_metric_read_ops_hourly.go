package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureComputeDiskMetricReadOpsHourly(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_compute_disk_metric_read_ops_hourly",
		Description: "Azure Compute Disk Metrics - Read Ops (Hourly)",
		List: &plugin.ListConfig{
			ParentHydrate: listAzureComputeDisks,
			Hydrate:       listComputeDiskMetricReadOpsHourly,
		},
		Columns: monitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the disk.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}

//// LIST FUNCTION

func listComputeDiskMetricReadOpsHourly(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	diskInfo := h.Item.(compute.Disk)

	return listAzureMonitorMetricStatistics(ctx, d, "HOURLY", "Microsoft.Compute/disks", "Composite Disk Read Operations/sec", *diskInfo.ID)
}
