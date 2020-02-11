# sellerapp
OnBoarding Process Assignments


## How to run make file
Goto Assignment-1

```
make debug
```

## How to run Docker file
```
//build docker file
docker build -t sellerapp  -f Dockerfile .

//list of available image
docker image ls

//run docker
docker run -d -p 50051:8000 sellerapp
```

## How to run Binary File
```
//go build 
go build -o sellerapp Assignment-1/myapplication.go Assignment-1/sellerapp.go

// run binary fike
./sellerapp

```

## Go Server:
```
server run on localhost:50051
```

## Mongodb:
```
MongoDB run on mongodb://localhost:27017
```


### For Testing
Run Curl
```
curl -H "Content-type:application/json" -X POST 'http://localhost:50051/sellerapp/v1/order' -d '{"orderData":{"sourceOrderId":"1234512345","items":[{"sku":"Business Cards","sourceItemId":"1234512346","components":[{"code":"Content","fetch":true,"path":"http://www.w2psite.com/businessCard.pdf"}]}],"shipments":[{"shipTo":{"name":"John Doe","companyName":"Acme","address1":"1234 Main St.","town":"Capitol","postcode":"12345","isoCountry":"US"},"carrier":{"code":"fedex","service":"ground"}}]}}'
```