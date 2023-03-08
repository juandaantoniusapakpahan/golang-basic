### TESTING

1. run application: **go run \*.go**
2. test service with **crul**:
   - curl -X GET --user batman:secret http://localhost:9000/student
   - curl -X GET --user batman:secret http://localhost:9000/student?id=s001
