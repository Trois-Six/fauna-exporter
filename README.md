# Fauna Prometheus Exporter
fauna-exporter permits to get billing and metric information from the Fauna DB dashboard. The billing API is not used anymore by the dashboard but it is always available, that is how I get these information.

## How do I build it?
```sh
$ make build
```

## Usage
```sh
$ ./fauna-exporter
NAME:
   fauna-exporter CLI - Run fauna-exporter

USAGE:
   fauna-exporter [global options] command [command options] [arguments...]

COMMANDS:
   exporter  Export Fauna metrics
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Exporter usage
```sh
$ ./fauna-exporter exporter --help
NAME:
   fauna-exporter exporter - Export Fauna metrics

USAGE:
   fauna-exporter exporter [command options] [arguments...]

DESCRIPTION:
   Export Fauna metrics from the Fauna dashboard.

OPTIONS:
   --host value            Host to listen on. [$HOST]
   --fauna-email value     Email to connect to the Fauna dashboard. [$FAUNA_EMAIL]
   --fauna-password value  Password to connect to the Fauna dashboard. [$FAUNA_PASSWORD]
   --fauna-days value      Number of days of the metrics. (default: 7) [$FAUNA_DAYS]
   --metrics-path value    Path of the metrics. (default: "/metrics") [$METRICS_PATH]
   --help, -h              show help (default: false)
```

## Example of output
```json
# HELP fauna_billing_data_storage_cost Billing information - Data Storage cost.
# TYPE fauna_billing_data_storage_cost gauge
fauna_billing_data_storage_cost{type="dollars"} 0.22077963314950466
# HELP fauna_billing_end_period Billing information - End Period.
# TYPE fauna_billing_end_period gauge
fauna_billing_end_period{date="2021-02-23T00:00:00Z"} 1
# HELP fauna_billing_metric_amount_byte_read_ops Billing information - Metric amount byte read operations.
# TYPE fauna_billing_metric_amount_byte_read_ops gauge
fauna_billing_metric_amount_byte_read_ops{type="byte_read_ops"} 45543
# HELP fauna_billing_metric_amount_byte_write_ops Billing information - Metric amount byte write operations.
# TYPE fauna_billing_metric_amount_byte_write_ops gauge
fauna_billing_metric_amount_byte_write_ops{type="byte_write_ops"} 5923
# HELP fauna_billing_metric_amount_compute_ops Billing information - Metric amount compute operations.
# TYPE fauna_billing_metric_amount_compute_ops gauge
fauna_billing_metric_amount_compute_ops{type="compute_ops"} 41743
# HELP fauna_billing_metric_amount_storage Billing information - Metric amount storage.
# TYPE fauna_billing_metric_amount_storage gauge
fauna_billing_metric_amount_storage{type="storage"} 23
# HELP fauna_billing_metric_usage_byte_read_ops Billing information - Metric usage byte read operations.
# TYPE fauna_billing_metric_usage_byte_read_ops gauge
fauna_billing_metric_usage_byte_read_ops{type="byte_read_ops"} 9.10846002e+08
# HELP fauna_billing_metric_usage_byte_write_ops Billing information - Metric usage byte write operations.
# TYPE fauna_billing_metric_usage_byte_write_ops gauge
fauna_billing_metric_usage_byte_write_ops{type="byte_write_ops"} 2.3688078e+07
# HELP fauna_billing_metric_usage_compute_ops Billing information - Metric usage compute operations.
# TYPE fauna_billing_metric_usage_compute_ops gauge
fauna_billing_metric_usage_compute_ops{type="compute_ops"} 1.85754844e+08
# HELP fauna_billing_metric_usage_storage Billing information - Metric usage storage.
# TYPE fauna_billing_metric_usage_storage gauge
fauna_billing_metric_usage_storage{type="storage"} 9.48241304e+08
# HELP fauna_billing_start_period Billing information - Start Period.
# TYPE fauna_billing_start_period gauge
fauna_billing_start_period{date="2021-01-25T00:00:00Z"} 1
# HELP fauna_billing_total_amount Billing information - Total amount.
# TYPE fauna_billing_total_amount gauge
fauna_billing_total_amount{type="dollars"} 932.32
# HELP fauna_billing_transactional_compute_ops_cost Billing information - Transactional Compute Ops cost.
# TYPE fauna_billing_transactional_compute_ops_cost gauge
fauna_billing_transactional_compute_ops_cost{type="dollars"} 417.948399
# HELP fauna_billing_transactional_read_ops_cost Billing information - Transactional Read Ops cost.
# TYPE fauna_billing_transactional_read_ops_cost gauge
fauna_billing_transactional_read_ops_cost{type="dollars"} 455.423001
# HELP fauna_billing_transactional_write_ops_cost Billing information - Transactional Write Ops cost.
# TYPE fauna_billing_transactional_write_ops_cost gauge
fauna_billing_transactional_write_ops_cost{type="dollars"} 59.220195000000004
# HELP fauna_exporter_build_info A metric with a constant '1' value labeled by version, revision, branch, and goversion from which fauna_exporter was built.
# TYPE fauna_exporter_build_info gauge
fauna_exporter_build_info{branch="",goversion="go1.15.8",revision="",version=""} 1
# HELP fauna_usage_byte_read_ops Usage information - byte read operations.
# TYPE fauna_usage_byte_read_ops gauge
fauna_usage_byte_read_ops{collection="a"} 1.3363e+06
fauna_usage_byte_read_ops{collection="all"} 1.118353852e+09
fauna_usage_byte_read_ops{collection="b"} 978
fauna_usage_byte_read_ops{collection="c"} 9.80934755e+08
fauna_usage_byte_read_ops{collection="d"} 709062
fauna_usage_byte_read_ops{collection="e"} 78977
fauna_usage_byte_read_ops{collection="f"} 1.0004098e+07
fauna_usage_byte_read_ops{collection="g"} 228006
fauna_usage_byte_read_ops{collection="h"} 3.1753621e+07
fauna_usage_byte_read_ops{collection="i"} 9.3283476e+07
fauna_usage_byte_read_ops{collection="j"} 1961
# HELP fauna_usage_byte_write_ops Usage information - byte write operations.
# TYPE fauna_usage_byte_write_ops gauge
fauna_usage_byte_write_ops{collection="a"} 5
fauna_usage_byte_write_ops{collection="all"} 2.9110121e+07
fauna_usage_byte_write_ops{collection="b"} 0
fauna_usage_byte_write_ops{collection="c"} 2.9076252e+07
fauna_usage_byte_write_ops{collection="d"} 5412
fauna_usage_byte_write_ops{collection="e"} 3042
fauna_usage_byte_write_ops{collection="f"} 357
fauna_usage_byte_write_ops{collection="g"} 0
fauna_usage_byte_write_ops{collection="h"} 2128
fauna_usage_byte_write_ops{collection="i"} 22898
fauna_usage_byte_write_ops{collection="j"} 27
# HELP fauna_usage_compute_ops Usage information - compute operations.
# TYPE fauna_usage_compute_ops gauge
fauna_usage_compute_ops{collection="a"} 1.046286e+06
fauna_usage_compute_ops{collection="all"} 2.25167222e+08
fauna_usage_compute_ops{collection="b"} 796
fauna_usage_compute_ops{collection="c"} 1.70627359e+08
fauna_usage_compute_ops{collection="d"} 349465
fauna_usage_compute_ops{collection="e"} 46634
fauna_usage_compute_ops{collection="f"} 616737
fauna_usage_compute_ops{collection="g"} 22804
fauna_usage_compute_ops{collection="h"} 1.44492e+06
fauna_usage_compute_ops{collection="i"} 5.1002634e+07
fauna_usage_compute_ops{collection="j"} 1614
# HELP fauna_usage_indexes Usage information - indexes.
# TYPE fauna_usage_indexes gauge
fauna_usage_indexes{collection="a"} 424
fauna_usage_indexes{collection="all"} 3.827482e+06
fauna_usage_indexes{collection="b"} 35
fauna_usage_indexes{collection="c"} 1.727713e+06
fauna_usage_indexes{collection="d"} 17401
fauna_usage_indexes{collection="e"} 589110
fauna_usage_indexes{collection="f"} 8366
fauna_usage_indexes{collection="g"} 131
fauna_usage_indexes{collection="h"} 117463
fauna_usage_indexes{collection="i"} 1.364849e+06
fauna_usage_indexes{collection="j"} 2380
# HELP fauna_usage_storage Usage information - storage.
# TYPE fauna_usage_storage gauge
fauna_usage_storage{collection="a"} 4010
fauna_usage_storage{collection="all"} 9.87241902e+08
fauna_usage_storage{collection="b"} 358
fauna_usage_storage{collection="c"} 7.01011977e+08
fauna_usage_storage{collection="d"} 2.2603285e+07
fauna_usage_storage{collection="e"} 765399
fauna_usage_storage{collection="f"} 79876
fauna_usage_storage{collection="g"} 430
fauna_usage_storage{collection="h"} 249049
fauna_usage_storage{collection="i"} 2.62590428e+08
fauna_usage_storage{collection="j"} 4296
# HELP fauna_usage_versions Usage information - versions.
# TYPE fauna_usage_versions gauge
fauna_usage_versions{collection="a"} 3586
fauna_usage_versions{collection="all"} 9.8341442e+08
fauna_usage_versions{collection="b"} 323
fauna_usage_versions{collection="c"} 6.99284264e+08
fauna_usage_versions{collection="d"} 2.2585884e+07
fauna_usage_versions{collection="e"} 176289
fauna_usage_versions{collection="f"} 71510
fauna_usage_versions{collection="g"} 299
fauna_usage_versions{collection="h"} 131586
fauna_usage_versions{collection="i"} 2.61225579e+08
fauna_usage_versions{collection="j"} 1916
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 0
go_gc_duration_seconds{quantile="0.25"} 0
go_gc_duration_seconds{quantile="0.5"} 0
go_gc_duration_seconds{quantile="0.75"} 0
go_gc_duration_seconds{quantile="1"} 0
go_gc_duration_seconds_sum 0
go_gc_duration_seconds_count 0
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 7
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.15.8"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 814848
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 814848
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 4414
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 159
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 4.005312e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 814848
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 6.5175552e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 1.572864e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 3555
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 6.5134592e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.6748416e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 0
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 3714
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 13888
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 16384
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 37536
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 49152
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.473924e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 728330
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 360448
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 360448
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 7.1912456e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 5
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 8192
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 81
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 9.916416e+06
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.61409547924e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 7.29362432e+08
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```