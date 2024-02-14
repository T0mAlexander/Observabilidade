#!/bin/sh

nohup sh /usr/sbin/grafana-server --homepath=/usr/share/grafana &

sleep 3

curl -sX PUT \
  -H "Content-Type: application/json" \
  -d '{
        "email":"admin-a@localhost",
        "name":"Admin A",
        "login":"admin-a",
        "theme":"dark"
      }' \
  http://admin:admin@localhost:3000/api/users/1

curl -sX PUT \
  -H "Content-Type: application/json" \
  -d '{
        "oldPassword": "admin",
        "newPassword": "grafana"
      }' \
  http://admin-a:admin@localhost:3000/api/user/password

curl -sX PUT \
  -H "Content-Type: application/json" \
  -d '{"name":"Organização A"}' \
  http://admin-a:grafana@localhost:3000/api/orgs/1

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{
        "name":"Time A",
        "email": "time-a@teste.com.br",
        "orgId": 1
      }' \
  http://admin-a:grafana@localhost:3000/api/teams

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{
        "name":"Usuario A",
        "email":"user-a@localhost",
        "login":"user-a",
        "password":"123456",
        "orgId": 1
      }' \
  http://admin-a:grafana@localhost:3000/api/admin/users

curl -sX PATCH \
  -H "Content-Type: application/json" \
  -d '{"role": "Viewer"}' \
  http://admin-a:grafana@localhost:3000/api/orgs/1/users/2

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{"userId": 2}' \
  http://admin-a:grafana@localhost:3000/api/teams/1/members

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{"name":"Organização B"}' \
  http://admin-a:grafana@localhost:3000/api/orgs/

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{
        "name":"Admin B",
        "email":"admin-b@localhost",
        "login":"admin-b",
        "password":"admin",
        "orgId": 2
      }' \
  http://admin-a:grafana@localhost:3000/api/admin/users

curl -sX PUT \
  -H "Content-Type: application/json" \
  -d '{
        "oldPassword": "admin",
        "newPassword": "grafana-b"
      }' \
  http://admin-b:admin@localhost:3000/api/user/password

curl -sX PUT \
  -H "Content-Type: application/json" \
  -d '{"isGrafanaAdmin": true}' \
  http://admin-a:grafana@localhost:3000/api/admin/users/3/permissions

curl -sX PATCH \
  -H "Content-Type: application/json" \
  -d '{"role": "Admin"}' \
  http://admin-a:grafana@localhost:3000/api/orgs/2/users/3

curl -sX DELETE -H "Content-Type: application/json" \
  http://admin-b:grafana-b@localhost:3000/api/orgs/2/users/1

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{"name":"Time B", "email": "time-b@teste.com.br", "orgId": 2}' \
  http://admin-b:grafana-b@localhost:3000/api/teams

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{
        "name":"Usuario B",
        "email":"user-b@localhost",
        "login":"user-b",
        "password":"123456",
        "orgId": 2
      }' \
  http://admin-b:grafana-b@localhost:3000/api/admin/users

curl -sX PATCH \
  -H "Content-Type: application/json" \
  -d '{"role": "Viewer"}' \
  http://admin-b:grafana-b@localhost:3000/api/orgs/2/users/4

curl -sX POST \
  -H "Content-Type: application/json" \
  -d '{"userId": 4}' \
  http://admin-b:grafana-b@localhost:3000/api/teams/2/members

sleep infinity