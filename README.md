# GoLang / Go + MongoDb  Micro service with OAuth2

![GoLang / Go + MongoDb  Microservice with OAuth2](https://cdn.pbrd.co/images/HLimzk1.png)

This is a Ready to deploy GoLang / Go Micro service with OAuth2 authentication/security.You can use this if you want to quick start developing your own custom Micro service by skipping 95% of your scratch works.
Hopefully this will save lot of your time as this API includes all the basic stuffs you need to get started.

This API also includes a developer dashboard with the API documentation which is developed in Angularjs 6. This will be useful to manage your developer access to the API documentation.

[DEMO](http://developers.go.mongodb.nintriva.net)
-------------------
```
http://developers.go.mongodb.nintriva.net
Login: developer/developer
```
ENDPOINTS
-------------------
```
SERVER: http://api.go.mongodb.nintriva.net
```


### Download and Install 

Step1: cd [GOPATH]/src

Step2:
git clone -b master https://github.com/sirinibin/golang-mongodb-microservice-with-oauth2.git rest-api


Step3: cd rest-api


Step4: update db/db.go with mongoDb server details


Step5: Run the app

./rest-api

### Set Up Developer Dashboard
Step1: cd developers


       vim proxy.conf.json
        {
          "/v1/*": {
            "target": "<API_END_POINT>",
            "secure": false,
            "changeOrigin": true
          }
        }


Step2. Install App:
       npm install


Step3: Start Developer dashboard
       ng serve --port 8007  --proxy-config proxy.conf.json