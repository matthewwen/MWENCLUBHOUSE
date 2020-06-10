# Register MWENCLUBHOUSE as a Service on Linux Computer

The Software that MWENCLUBHOUSE runs on is Ubuntu Server. In the past, I would build the program, use htop to kill the old program, and then run the program in the background again. 

example:
```
make build
sudo make install
go clean
go build
sudo htop (find program, then kill)
sudo ./server dep (dep meaning to release the deployed version)
```

This is tedious and can cost time. As a result, I register my executable as a service so I can use systemctl to check the status of the program.

```
make build
sudo make install (library to use with CGO)
go clean
go build

sudo systemctl restart server
```

It also tracks the printf and print statement from GoLang. As a result, if the server crash, a print statement can inform me what happened and when it happened.

```
● server.service - Server Running Backend for Matthew Wen's mwenclubhouse
   Loaded: loaded (/etc/systemd/system/server.service; static; vendor pre
   Active: active (running) since Wed 2020-06-10 04:36:28 UTC; 8min ago
 Main PID: 1558 (server)
    Tasks: 7 (limit: 4915)
   CGroup: /system.slice/server.service
           └─1558 /home/mwen/server/server dep

Jun 10 04:36:28 mwenclubhouse systemd[1]: Started Server Running Backend 
Jun 10 04:36:28 mwenclubhouse server[1558]: Deploy
```
