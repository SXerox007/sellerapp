
# CURL to POST the order data
curl -H "Content-type:application/json" -X POST 'http://:50051/sellerapp/v1/order' -d '{"orderData":{"sourceOrderId":"1234512345","items":[{"sku":"Business Cards","sourceItemId":"1234512346","components":[{"code":"Content","fetch":true,"path":"http://www.w2psite.com/businessCard.pdf"}]}],"shipments":[{"shipTo":{"name":"John Doe","companyName":"Acme","address1":"1234 Main St.","town":"Capitol","postcode":"12345","isoCountry":"US"},"carrier":{"code":"fedex","service":"ground"}}]}}' -k -v
# CURL test Get 
curl -H "Content-type:application/json" -X GET 'http://:50051/sellerapp/v1/order'
# initialize Go modules
GO111MODULE=on go mod init sellerapp

# go build 
go build -o sellerapp Assignment-1/myapplication.go Assignment-1/sellerapp.go

# build docker file
# docker build -t sellerapp  -f Dockerfile .
docker build -t sellerapp .

# list of available image
docker image ls
#
docker ps
# docker ps -a

# run docker
docker run -d -p 50051:50051 sellerapp
# docker run --rm -p 50051:8000 sellerapp
# docker run -p 50051:8000 -m http.server --bind 0.0.0.0
# docker run -p 50051:8000 -t sellerapp .
# docker run -t -P sellerapp
# docker run --publish 50051:50051 -t sellerapp
# docker run -itd --name=sellerapp-test sellerapp

# MongoDB docker container
# docker network create container-net
# docker run -itd --name mongodb --network container-net -p 27017:27017 mongo 
#
docker-compose up
docker-compose build && docker-compose up


#
docker run sellerapp

# stop container
docker container stop <container ID>

# docker run
sudo docker run <container ID>

#debug
docker attach <container ID>

# delete unused images
docker system prune
# docker system prune -a

# check port
docker port <name>

#
docker-machine ip default

# network inspect
docker network inspect bridge

# docker exec
docker exec <name> ls

# start a Bash session.
docker exec -d <name> touch /temp/execWorks
docker exec -it <name> bash

#ssl
openssl s_client -connect localhost:50051


# don't have /usr/local/bin
PATH=/usr/local/bin:$PATH
hash -r

