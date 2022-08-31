# Bigquery Cli

A little command line application to execute biquery from terminal.

## Example

```bash
bq-cli --query 'select created_at, id from bookings limit 3'
+----------------------+--------------------------------------+
|      CREATED AT      |                  ID                  |
+----------------------+--------------------------------------+
| 2017-12-15T08:21:15Z | fb1f456b-e1bb-4723-9075-ec67bc74433b |
| 2017-12-20T05:06:08Z | 274670ca-bd6a-40fd-8e3b-9d9d1441c32b |
| 2020-07-24T08:18:55Z | 03c13e25-8fc1-49ee-a8e7-2c0c1fcc94e8 |
+----------------------+--------------------------------------+
```
