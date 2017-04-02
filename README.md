Gopher
=

Gopher, written in Go (of course) munches on local log files, parses them using the provided Grok filter, and delivers them to AWS Firehose for processing into ElasticSearch/S3 etc.

**Configuration**

```
FirehoseDeliveryStreamName = "foo"

[[Files]]
location = "/var/log/system.log"
grok = "%{GREEDYDATA:message}"
```

**Usage**

`AWS_PROFILE=gs AWS_REGION=eu-west-1 Gopher --configfile=gopher.cfg`


