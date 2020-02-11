
# CURL to POST the order data
curl -H "Content-type:application/json" -X POST 'http://171.48.101.121/sellerapp/v1/order' -d '{"orderData":{"sourceOrderId":"1234512345","items":[{"sku":"Business Cards","sourceItemId":"1234512346","components":[{"code":"Content","fetch":true,"path":"http://www.w2psite.com/businessCard.pdf"}]}],"shipments":[{"shipTo":{"name":"John Doe","companyName":"Acme","address1":"1234 Main St.","town":"Capitol","postcode":"12345","isoCountry":"US"},"carrier":{"code":"fedex","service":"ground"}}]}}'

# initialize Go modules
GO111MODULE=on go mod init sellerapp

#go build 
go build -o sellerapp Assignment-1/myapplication.go Assignment-1/sellerapp.go

#build docker file
docker build -t sellerapp  -f Dockerfile .

#list of available image
docker image ls
#
docker ps

#run docker
docker run -d -p 50051:8000 sellerapp
# docker run --rm -p 50051:8000 sellerapp
# docker run -p 50051:8000 -m http.server --bind 0.0.0.0

#
docker run sellerapp

# stop container
docker container stop <container ID>

#docker run
sudo docker run <container ID>

#debug
docker attach <container ID>