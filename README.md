# Nagios JSON API

Intended to provide a JSON API like features for Nagios installations. Its written in golang and therefore systems wanting to run it only need to run a single binary for their distribution.

Installation:
==
>> Follow [golang installation](https://golang.org/doc/install) to install golang.
>> >> Clone this repo.
>> >> >> Run: go build . from inside repo directory.
>> >> >> >> You can now use the standalone binary to run the service.

Run:
==
```
$ ./go-nagios-api --addr=:9090 --cachefile=/opt/nagios/object.cache --statusfile=/opt/nagios/status.dat --commandfile=/opt/nagios/nagios.cmd

Or you can provide a configuration file with these parameter in json format (configuration file overwrites cli flags)

$ ./go-nagios-api --config=nagios-api.json
```
It will start the api service on port 8080. If you wish to change the port simply pass --addr=:80 to make it run on port 80. For running in production see init scripts.

API Calls
==

#### Hosts and Services
```
 GET /hosts : get all configured hosts
 GET /host/<hostname> : get this host
 GET /hoststatus : get all hoststatus
 GET /hoststatus/<hostname> : get hoststatus for this host
 GET /hostgroups : get all configured hostgroups
 GET /services : get all configured services
 GET /servicestatus : get all servicestatus
 GET /servicestatus/<servicename> : get service status for this service
```

#### External Commands
```
POST /disable_notifications 
POST /enable_notifications
POST /disable_host_check  
POST /enable_host_check   
POST /disable_host_notifications
POST /enable_host_notifications
POST /acknowledge_host_problem
POST /acknowledge_service_problem
POST /add_host_comment
POST /add_svc_comment
POST /del_all_host_comment
POST /del_all_svc_comment
POST /del_host_comment
POST /del_svc_comment
POST /disable_all_notification_beyond_host
POST /enable_all_notification_beyond_host
POST /disable_hostgroup_host_checks
POST /enable_hostgroup_host_checks
POST /disable_hostgroup_host_notifications
POST /enable_hostgroup_host_notifications
POST /disable_hostgroup_svc_checks
POST /enable_hostgroup_svc_checks
POST /disable_hostgroup_svc_notifications
POST /enable_hostgroup_svc_notifications
POST /disable_host_and_child_notifications
POST /enable_host_and_child_notifications
POST /schedule_host_downtime
POST /force_service_checks
POST /force_host_checks
```

#### Examples
```
To disable host check for host host1.example.net
curl -i -XPOST http://127.0.0.1:9090/disable_host_check -d '{"hostname": "host1.example.net"}'

To disable notification for all hosts
curl -i -XPOST http://127.0.0.1:9090/disable_notifications

To get all configured hostgroups
curl -i http://127.0.0.1:9090/hostgroups

To disable host check for a particular hostgroup
curl -i -XPOST http://127.0.0.1:9090/disable_hostgroup_host_checks -d '{"hostgroup":"AwesomeHostGroup"}'

To get details for a given host host1.example.net
curl -i http://127.0.0.1:9090/host/host1.example.net
```

