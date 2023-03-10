### TESTING

1. FORM DATA

   - curl -X POST http://localhost:4000/user \
      -d 'name=Hanzoo' \
      -d 'email=hanzoop2@gmail.com'

2. JSON Payload

   - curl -X POST http://localhost:4000/user \
     -H 'Content-Type: application/json' \
     -d '{"name":"hayabuza", "email":"hayabuza12@gmail.com"}'

3. XML Payload

   - curl -X POST http://localhost:4000/user \
     -H 'Content-Type: application/xml' \
     -d '<?xml version="1.0"?>\
      <Data>\
      <Name>Gomgom</Name>\
      <Email>gomgom@gmail.com</Email>\
      </data>'

4. QueryParam
   - curl -X GET http://localhost:4000/user?name=rechard&email=rechard212@gmail.com
