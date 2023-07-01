to run open a terminal in this directory(nginx) and run the below command

docker run --network=host  --name docker-nginx2 -p 443:443 -p 80:80  -v ./web:/etc/nginx/conf.d  nginx
FOR DOCKERS



docker run --add-host host.docker.internal:host-gateway  --name docker-nginx2 -p 443:443 -p 80:80  -v ./web:/etc/nginx/conf.d  nginx
FOR ME IN LOCALHOST


--network="host" is for access to send to localhost 6433 for gatway that is not in docker

