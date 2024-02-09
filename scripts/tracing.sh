curl --request GET \
  --url http://localhost:8080/ \
  --header 'traceparent: 00-df853039b602c93e641526aaa7d67b8c-339f2b7a83c7d606-01'

curl --request POST \
  --url http://localhost:8080/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Husky Siberiano", "description": "Um cachorro do ártico!" }'

curl --request POST \
  --url http://localhost:8080/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Gralha-Azul", "description": "Um passáro agradável" }'

curl --request POST \
  --url http://localhost:8080/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Richard Rasmussen","description": "Biólogo"}'

curl --request GET \
  --url http://localhost:8080/peanuts/1

curl --request GET \
  --url http://localhost:8080/peanuts/1

curl --request GET \
  --url http://localhost:8080/peanuts/2

curl --request GET \
  --url http://localhost:8080/peanuts/2

curl --request GET \
  --url http://localhost:8080/peanuts/3

curl --request GET \
  --url http://localhost:8080/peanuts/3

####

curl --request GET \
  --url http://localhost:8081/ \
  --header 'traceparent: 00-df853039b602c93e641526aaa7d67b8c-339f2b7a83c7d606-01'

curl --request POST \
  --url http://localhost:8081/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Richard Rasmussen","description": "Biólogo"}'

curl --request POST \
  --url http://localhost:8081/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Gralha-Azul", "description": "Um passáro agradável" }'

curl --request POST \
  --url http://localhost:8081/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Richard Rasmussen","description": "Biólogo"}'

curl --request GET \
  --url http://localhost:8081/peanuts/1

curl --request GET \
  --url http://localhost:8081/peanuts/1

curl --request GET \
  --url http://localhost:8081/peanuts/2

curl --request GET \
  --url http://localhost:8081/peanuts/2

curl --request GET \
  --url http://localhost:8081/peanuts/3

curl --request GET \
  --url http://localhost:8081/peanuts/3

####

curl --request GET \
  --url http://localhost:8082/ \
  --header 'traceparent: 00-df853039b602c93e641526aaa7d67b8c-339f2b7a83c7d606-01'

curl --request POST \
  --url http://localhost:8082/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Richard Rasmussen","description": "Biólogo"}'

curl --request POST \
  --url http://localhost:8082/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Gralha-Azul", "description": "Um passáro agradável" }'

curl --request POST \
  --url http://localhost:8082/peanuts \
  --header 'content-type: application/json' \
  --data '{"name": "Richard Rasmussen","description": "Biólogo"}'

curl --request GET \
  --url http://localhost:8082/peanuts/1

curl --request GET \
  --url http://localhost:8082/peanuts/1

curl --request GET \
  --url http://localhost:8082/peanuts/2

curl --request GET \
  --url http://localhost:8082/peanuts/2

curl --request GET \
  --url http://localhost:8082/peanuts/3

curl --request GET \
  --url http://localhost:8082/peanuts/3