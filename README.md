### Try it out

Notice \<id\> must be given.
  
```
# POST
curl -X POST http://ec2-18-184-29-167.eu-central-1.compute.amazonaws.com:8081/users -H "Content-Type: application/json" -d '{"id":"9","email":"9@9","name":"nine"}'

# GET
curl -X GET http://ec2-18-184-29-167.eu-central-1.compute.amazonaws.com:8081/users/<id>

# PATCH
curl -X PATCH http://ec2-18-184-29-167.eu-central-1.compute.amazonaws.com:8081/users/<id> -H "Content-Type: application/json" -d '{"email":"8@8","name":"eight"}'

# DELETE
curl -X DELETE http://ec2-18-184-29-167.eu-central-1.compute.amazonaws.com:8081/users/<id>
```

### Docker 
#### Build

Notice \<your-mongodb-atlas-uri\> must be given.
```
docker image build https://github.com/LimKianAn/go-openapi-mongodb.git --build-arg build_arg_mongodb_atlas_uri=<your-mongodb-atlas-uri>
```

#### Run
Noice \<image-ID\> must be given.
```
docker container run -d -p 8081:8080 <image-ID>
```
