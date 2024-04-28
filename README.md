# go with test container 

This is a example for test go with connect to mysql with test container

## Examples
### Example 1 

use Test Container to test External Service for example Mysql 

[test/integration/dbclient_test.go](test/integration/dbclient_test.go)

### Example 2

use Test Container to test Current Service as container

[test/integration/httpclient_test.go](test/integration/httpclient_test.go)

## Problem I face

### mysql:8.0.36 has connection problem due to user/password setup failed

```
2024/04/28 22:15:28 🚧 Waiting for container id deae46b0d495 image: mysql:8.0.36. Waiting for: &{timeout:<nil> Log:port: 3306  MySQL Community Server IsRegexp:false Occurrence:1 PollInterval:100ms}
2024/04/28 22:16:28 container logs (context deadline exceeded):
2024-04-28 14:15:28+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.36-1.el8 started.
2024-04-28 14:15:28+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
2024-04-28 14:15:28+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.36-1.el8 started.
2024-04-28 14:15:28+00:00 [Warn] [Entrypoint]: MYSQL_PASSWORD specified, but missing MYSQL_USER; MYSQL_PASSWORD will be ignored
2024-04-28 14:15:28+00:00 [Note] [Entrypoint]: Initializing database files
2024-04-28T14:15:28.783248Z 0 [Warning] [MY-011068] [Server] The syntax '--skip-host-cache' is deprecated and will be removed in a future release. Please use SET GLOBAL host_cache_size=0 instead.
2024-04-28T14:15:28.783328Z 0 [System] [MY-013169] [Server] /usr/sbin/mysqld (mysqld 8.0.36) initializing of server in progress as process 81
2024-04-28T14:15:28.887489Z 1 [System] [MY-013576] [InnoDB] InnoDB initialization has started.
2024-04-28T14:15:39.217310Z 1 [System] [MY-013577] [InnoDB] InnoDB initialization has ended.
2024-04-28T14:16:05.477287Z 6 [Warning] [MY-010453] [Server] root@localhost is created with an empty password ! Please consider switching off the --initialize-insecure option.
```

### mysql:5.7 is available

