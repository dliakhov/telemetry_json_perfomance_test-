package main

import (
	"context"
	"testing"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

//key:value,key2:value2

const (
	getCPUArchitectureQueryStringParsing = `SELECT d,
       version        AS bucket,
       COUNT(bucket)
FROM (
         SELECT event_date                                                               as d,
                arrayElement(splitByChar(':', arrayElement(splitByChar(',', arrayJoin(arrayFilter(x -> x.1 = 'pmm_cpu_info_text', pmm_server_metrics)) .2), 1)), 2) as version
         FROM telemetryd.pmm_metrics_2
         GROUP BY d, pmm_server_telemetry_id, version
         )
GROUP BY d, bucket
order by d desc`

	getCPUArchitectureQueryJsonParsing = `SELECT d,
       version        AS bucket,
       COUNT(bucket)
FROM (
         SELECT event_date                                                               as d,
                JSON_VALUE(arrayJoin(arrayFilter(x -> x.1 = 'pmm_cpu_info_json', pmm_server_metrics)) .2, '$.info.cpu_architecture') as version
         FROM telemetryd.pmm_metrics_view
         GROUP BY d, pmm_server_telemetry_id, version
         )
GROUP BY d, bucket
order by d desc`

	getSimpleMetric = `SELECT d,
       version        AS bucket,
       COUNT(bucket)
FROM (
         SELECT event_date                                                               as d,
                arrayJoin(arrayFilter(x -> x.1 = 'pmm_server_grafana_stat_active_users', pmm_server_metrics)) .2 as version
         FROM telemetryd.pmm_metrics_view
         GROUP BY d, pmm_server_telemetry_id, version
         )
GROUP BY d, bucket
order by d desc`
)

// test

func BenchmarkSelectJsonQuery(b *testing.B) {
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
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		rows, err := conn.Query(ctx, getCPUArchitectureQueryJsonParsing)
		if err != nil {
			b.Fatal(err)
		}
		rows.Close()
		if rows.Err() != nil {
			b.Fatal(rows.Err())
		}
	}

	err = conn.Close()
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectStringQuery(b *testing.B) {
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
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		rows, err := conn.Query(ctx, getCPUArchitectureQueryStringParsing)
		if err != nil {
			b.Fatal(err)
		}
		rows.Close()
		if rows.Err() != nil {
			b.Fatal(rows.Err())
		}
	}

	err = conn.Close()
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectSimpleMetric(b *testing.B) {
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
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		rows, err := conn.Query(ctx, getSimpleMetric)
		if err != nil {
			b.Fatal(err)
		}
		rows.Close()
		if rows.Err() != nil {
			b.Fatal(rows.Err())
		}
	}

	err = conn.Close()
	if err != nil {
		b.Fatal(err)
	}
}
