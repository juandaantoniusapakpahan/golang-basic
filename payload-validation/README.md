### TESTING

1. Valid payload

   - curl -X POST http://localhost:8080/users \
     -H 'Content-Type: application/json' \
     -d '{"name": "Grishing", "email": "grishing900@gmail.com", "age": 34}'

   - result: **true**

2. Bad payload

   - curl -X POST http://localhost:8080/users \
      -H 'Content-Type: application/json' \
      -d '{"name": "Grishing", "email": "grishing900@gmail.com", "age": 123}

   - result: {
     "message": "Internal Server Error"
     }
