package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

func main() {
	insertData()
}

func insertData() {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{"127.0.0.1:7011"},
			Auth: clickhouse.Auth{
				Database: "telemetryd",
			},
			//Debug:           true,
			DialTimeout:     time.Second,
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxLifetime: time.Hour,
		})
	)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	err = generateTestData(conn)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Successfully finished")
}

func generateTestData(conn driver.Conn) error {
	var orgIds []string
	for i := 0; i < 5_000; i++ {
		orgIds = append(orgIds, uuid.NewString())
	}
	postgresVersions := []string{
		"9.6.1",
		"9.6.0",
		"9.0.0",
		"8.0.0",
		"7.0.0",
		"6.0.0",
		"5.0.0",
		"4.0.0",
		"3.0.0",
		"2.0.0",
		"1.0.0",
	}

	pmmVersions := []string{
		"1.2.3",
		"1.2.4",
		"1.3",
		"1.4.3",
		"1.5.3",
		"2.2.3",
		"3.2.3",
		"4.2.3",
		"1.3.3",
		"1.4.3",
		"5.2.3",
	}

	cpuArchitectures := []string{
		"x86",
		"arm",
	}

	rand.Seed(time.Now().UnixNano())

	for _, orgId := range orgIds {
		batch, err := conn.PrepareBatch(context.Background(), `INSERT INTO telemetryd.pmm_metrics`)
		if err != nil {
			return err
		}

		batch2, err := conn.PrepareBatch(context.Background(), `INSERT INTO telemetryd.pmm_metrics_2`)
		if err != nil {
			return err
		}

		version := postgresVersions[rand.Intn(len(postgresVersions))]
		pmmVersion := pmmVersions[rand.Intn(len(pmmVersions))]
		randCount := rand.Intn(100)

		date := time.Now().Add(-30 * 24 * time.Hour)
		for i := 0; i < 30; i++ {
			err := batch.Append(
				uuid.NewString(),
				date,
				orgId,
				nil,
				"",
				int32(1),
				[][]string{
					{`pmm_server_backup_management_enabled`, `0`},
					{`pmm_server_alert_manager_enabled`, `0`},
					{`pmm_cpu_info_json`, fmt.Sprintf(`{"name": "cpu", "info": { "cpu_architecture": "%s","percent": %f } }`, cpuArchitectures[rand.Intn(len(cpuArchitectures))], rand.Float64()*100)},
					{`pmm_another_info`, fmt.Sprintf(`{"name": "pmm", "info": { "count": %d} }`, randCount)},
					{`pmm_server_info`, fmt.Sprintf(`{"name": "pmm","version": "%s"}`, pmmVersion)},
					{`pmm_agents_status_info`, fmt.Sprintf(`{"name": "postgres","version": "%s"}`, version)},
					{`pmm_agents_status_info_2`, fmt.Sprintf(`{"name": "mysql","version": "%s"}`, version)},
					{`pmm_agents_status_info_3`, fmt.Sprintf(`{"name": "proxysql","version": "%s"}`, version)},
					{`pmm_server_grafana_stat_active_users`,
						`0`},
					{`pmm_server_grafana_stat_totals_annotations`,
						`0`},
					{`pmm_server_data_retention_period`,
						`2592000000000000`},
					{`pmm_server_usage_nodes_count`,
						`0`},
					{`pmm_server_usage_services_count`,
						`0`},
					{`pmm_server_usage_environments_count`,
						`0`},
					{`pmm_server_usage_clusters_count`,
						`0`},
					{`pmm_server_pmm_agent_version`,
						`2.32.0`},
					{`pmm_server_node_type`,
						`generic`},
					{`dbaas_services_count`,
						`0`},
					{`dbaas_clusters_count`,
						`0`},
					{`k8s_clusters_count`,
						`0`},
					{`pmm_server_version`,
						`2.30.0`},
					{`pmm_server_uptime_seconds`,
						`2790`},
					{`pmm_server_deployment_method`,
						`DOCKER`},
					{`pmm_server_country`,
						``},
				},
				[]string{},
			)
			if err != nil {
				return err
			}

			err = batch2.Append(
				uuid.NewString(),
				date,
				orgId,
				nil,
				"",
				int32(1),
				[][]string{
					{`pmm_server_backup_management_enabled`, `0`},
					{`pmm_server_alert_manager_enabled`, `0`},
					{`pmm_cpu_info_json`, fmt.Sprintf(`{"name": "cpu", "info": { "cpu_architecture": "%s","percent": %f } }`, cpuArchitectures[rand.Intn(len(cpuArchitectures))], rand.Float64()*100)},
					{`pmm_another_info`, fmt.Sprintf(`{"name": "pmm", "info": { "count": %d} }`, randCount)},
					{`pmm_server_info`, fmt.Sprintf(`{"name": "pmm","version": "%s"}`, pmmVersion)},
					{`pmm_agents_status_info`, fmt.Sprintf(`{"name": "postgres","version": "%s"}`, version)},
					{`pmm_agents_status_info_2`, fmt.Sprintf(`{"name": "mysql","version": "%s"}`, version)},
					{`pmm_agents_status_info_3`, fmt.Sprintf(`{"name": "proxysql","version": "%s"}`, version)},

					{`pmm_cpu_info_text`, fmt.Sprintf(`cpu_architecture:%s,percent:%f`, cpuArchitectures[rand.Intn(len(cpuArchitectures))], rand.Float64()*100)},

					{`pmm_server_grafana_stat_active_users`,
						`0`},
					{`pmm_server_grafana_stat_totals_annotations`,
						`0`},
					{`pmm_server_data_retention_period`,
						`2592000000000000`},
					{`pmm_server_usage_nodes_count`,
						`0`},
					{`pmm_server_usage_services_count`,
						`0`},
					{`pmm_server_usage_environments_count`,
						`0`},
					{`pmm_server_usage_clusters_count`,
						`0`},
					{`pmm_server_pmm_agent_version`,
						`2.32.0`},
					{`pmm_server_node_type`,
						`generic`},
					{`dbaas_services_count`,
						`0`},
					{`dbaas_clusters_count`,
						`0`},
					{`k8s_clusters_count`,
						`0`},
					{`pmm_server_version`,
						`2.30.0`},
					{`pmm_server_uptime_seconds`,
						`2790`},
					{`pmm_server_deployment_method`,
						`DOCKER`},
					{`pmm_server_country`,
						``},
				},
				[]string{},
			)
			if err != nil {
				return err
			}
			date = date.Add(24 * time.Hour)
		}
		err = batch.Send()
		if err != nil {
			return err
		}
		err = batch2.Send()
		if err != nil {
			return err
		}
	}
	return nil
}

//func generateTestData(insertData func(insert string) error) error {
//	var orgIds []string
//	for i := 0; i < 1000; i++ {
//		orgIds = append(orgIds, uuid.NewString())
//	}
//	postgresVersions := []string{
//		"9.6.1",
//		"9.6.0",
//		"9.0.0",
//		"8.0.0",
//		"7.0.0",
//		"6.0.0",
//		"5.0.0",
//		"4.0.0",
//		"3.0.0",
//		"2.0.0",
//		"1.0.0",
//	}
//
//	pmmVersions := []string{
//		"1.2.3",
//		"1.2.4",
//		"1.3",
//		"1.4.3",
//		"1.5.3",
//		"2.2.3",
//		"3.2.3",
//		"4.2.3",
//		"1.3.3",
//		"1.4.3",
//		"5.2.3",
//	}
//
//	cpuArchitectures := []string{
//		"x86",
//		"arm",
//	}
//
//	rand.Seed(time.Now().UnixNano())
//
//	date := time.Now().Add(-30 * 24 * time.Hour)
//
//	for i := 0; i < 30; i++ {
//		for _, orgId := range orgIds {
//			dateStr := date.Format("2006-01-02")
//			version := postgresVersions[rand.Intn(len(postgresVersions))]
//			pmmVersion := pmmVersions[rand.Intn(len(pmmVersions))]
//			cpuArchitecture := cpuArchitectures[rand.Intn(len(cpuArchitectures))]
//			cpuUsage := rand.Float64() * 100
//
//			randCount := rand.Intn(100)
//
//			err := insertData(fmt.Sprintf(`
//INSERT INTO telemetryd.pmm_metrics
//    (event_time, pmm_server_telemetry_id, portal_org_id, portal_org_name, percona_customer_tier, pmm_server_metrics, pmm_server_metrics_with_classifier)
//VALUES
//    ('%s', '%s', null, '', 1, [('pmm_server_stt_enabled', '1'), ('pmm_server_backup_management_enabled', '0'), ('pmm_server_alert_manager_enabled', '0'), ('pmm_cpu_info', '{"name": "cpu", "info": { "cpu_architecture": "%s", "percent": %f } }'), ('pmm_cpu_info', '{"name": "cpu", "info": { "cpu_architecture": "%s", "percent": %f } }'), ('pmm_cpu_info', '{"name": "cpu", "info": { "cpu_architecture": "%s", "percent": %f } }'), ('pmm_another_info', '{"name": "pmm", "info": { "count": %d} }'), ('pmm_server_info', '{"name": "pmm", "version": "%s"}'), ('pmm_agents_status_info', '{"name": "postgres", "version": "%s"}'), ('pmm_agents_status_info_2', '{"name": "mysql", "version": "%s"}'), ('pmm_agents_status_info_3', '{"name": "proxysql", "version": "%s"}'), ('pmm_server_grafana_stat_total_users', '1'), ('pmm_server_grafana_stat_active_users', '0'), ('pmm_server_grafana_stat_totals_annotations', '0'), ('pmm_server_data_retention_period', '2592000000000000'), ('pmm_server_usage_nodes_count', '0'), ('pmm_server_usage_services_count', '0'), ('pmm_server_usage_environments_count', '0'), ('pmm_server_usage_clusters_count', '0'), ('pmm_server_pmm_agent_version', '2.32.0'), ('pmm_server_node_type', 'generic'), ('dbaas_services_count', '0'), ('dbaas_clusters_count', '0'), ('k8s_clusters_count', '0'), ('pmm_server_version', '2.30.0'), ('pmm_server_uptime_seconds', '2790'), ('pmm_server_deployment_method', 'DOCKER'), ('pmm_server_country', '')], []);`,
//				dateStr, orgId, cpuArchitecture, cpuUsage, cpuArchitecture, cpuUsage, cpuArchitecture, cpuUsage, randCount, pmmVersion, version, version, version))
//			if err != nil {
//				return err
//			}
//		}
//		date = date.Add(24 * time.Hour)
//	}
//	return nil
//}
