package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2020-01-01/mysql"
	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

//// TABLE DEFINITION

func tableAzureMySQLServerMetricActiveConnectionsDaily(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "azure_mysql_server_metric_active_connections_daily",
		Description: "Azure MySQL Server Metrics - Active Connections (Daily)",
		List: &plugin.ListConfig{
			ParentHydrate: listMySQLServers,
			Hydrate:       listAzureMySQLServerMetricActiveConnectionsDaily,
		},
		Columns: monitoringMetricColumns([]*plugin.Column{
			{
				Name:        "name",
				Description: "Contains ID to identify a server uniquely.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("DimensionValue").Transform(lastPathElement),
			},
		}),
	}
}

//// LIST FUNCTION

func listAzureMySQLServerMetricActiveConnectionsDaily(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	serverInfo := h.Item.(mysql.Server)

	return listAzureMonitorMetricStatistics(ctx, d, "DAILY", "Microsoft.DBforMySQL/servers", "active_connections", *serverInfo.ID)
}
