#TestCDIDataImpact

## Developpement

We use a `docker-compose.yaml` file to launch app dependency dockers like :
```
docker-compose up --build
```

You will have to create an `.env` file from the` .env.example` template and put your parameters there for the dev.

Run the following command to launch dependencies for local dev:

```
make deploy
```

This one you tearn down the dependencies:

```
make teardown
```


## Endpoints

### Health check

- `GET /status`

  Example response:

  ```json
  {
    "status": "ok"
  }
  ```
- ``POST /api/users `` 

To add users from ``DataSet.json`` to the mongodb database and haching the password by bcrypt
Example response:

  ```
  [
      {
          "ID": string,
          "Email": string,
          "Password": string,
          "IsActive": boolean,
          "Balance": string,
          "Age": int,
          "Name": string,
          "Gender": string,
          "Company": string,
          "Phone": string,
          "Address": string,
          "About": string,
          "Registered": string,
          "Latitude": int,
          "Longitude": int,
          "Friends": []friend,
          "Tags": []string,
          "Data": string
      }
  ]
  ```

- ``POST /api/auth ``
Example Body:

```json
 {
   "email": "nikkifarley@anivet.com",
   "password":"CGUsfQkS06lo7LM2guSV" 
 }
```
Example response:
```json
 {
     "code": 200,
     "expire": "2022-03-21T22:34:12+01:00",
     "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im5pa2tpZmFybGV5QGFuaXZldC5jb20iLCJleHAiOjE2NDc4OTg0NTIsImlkIjoiIiwiaWRlbnRpdHkiOiIiLCJvcmlnX2lhdCI6MTY0Nzg5NDg1Mn0.BXOvzTaPs-Qg-9U8gOb-CpIEgd_vCZx1neNFcMNQdmw"
 }
```
This Endpoint generate a BearerToken which will be needed for accessing all the other endpoints and expires after 60 minutes

- ``Get /api/users/list ``
To get all users in DB

- ``Get /api/users/:id ``
To get one user by id

- ``Patch /api/users/:id ``
Update user by id

## Tests

Unit tests:

```sh
go test ./...
```
