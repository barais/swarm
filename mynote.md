to start it


Stop you local docker
start it

> sudo docker daemon -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock

$GOPATH/bin/swarm help
$GOPATH/bin/swarm create
$GOPATH/bin/swarm manage token://<cluster_id> -H 127.0.0.1:2378
$GOPATH/bin/swarm join --addr=<node_ip:2375> token://<cluster_id>



$GOPATH/bin/swarm manage token://4fb70dfdecd4801436aef2a9aa1403d4 -H 127.0.0.1:2378 --strategy=myownstrategy -c task-experimental


$GOPATH/bin/swarm join --addr=127.0.0.1:2375 token://4fb70dfdecd4801436aef2a9aa1403d4


docker -H 127.0.0.1:2378 run -d -e name:toto -e deps:titi ubuntu /bin/sleep 10
docker -H 127.0.0.1:2378 run -d -e name:titi ubuntu /bin/sleep 10
