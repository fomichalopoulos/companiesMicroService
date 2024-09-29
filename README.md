# Companies Backend 

## Table Of Contents
- [Description](#description)
- [Mongo](#mongo)
    * [Schema](#schema)
    * [Initialization](#initialization)
    * [Mongo Init Image](#mongo-init-image)
    * [MongoDB](*mongoDB)
- [KAFKA](#kafka)
- [API](#api)
    * [Image](#image)
    * [Configuration](#configuration)
    * [Deployment](#deployment)
    * [Testing](#testing)



### Description
The backend of a companies microservice supporting the operations mentioned below: 

```
1. Creation of a company by an authenticated user 
2. Get of one company by any user
3. Deletion of a company by an authenticated user
4. Patching of a company by an authenticated user
```

### Mongo

#### Schema
The schema of the database is pretty simple as each company consists of the following:

1. _id --> unique, required
2. name --> string, name cannot exceed 15 characters, required
3. description --> string, cannot exceed 3000 characters, required
3. amount --> int, required
4. registered  --> string, one of {Corporations, NonProfit, Cooperative, Sole Proprietorship}

#### Initialization
The initialization of the MongoDB is done via the `mongoInit` service as declared in the [docker-compose.yml](/apiContainerized/docker-compose.yml). This service is responsible for adding mockup data into the MongoDB.

#### Mongo Init Image
The mongo initialization image is declared in the [Dockerfile](/mongoInitialization/Dockerfile).

#### MongoDB 
The `mongoDB` service is declared in the single [docker-compose.yml](/apiContainerized/docker-compose.yml).  

### KAFKA
A KAFKA component is deployed as a docker service as well and declared in the [docker-compose.yml](/apiContainerized/docker-compose.yml).

### API

#### Image
In order to build image locally the following [Dockerfile](/Dockerfile) will be executed from a single `docker-compose.yml` when the service will be built.

#### Configuration
The services variables are configured before their deployment via the [.env](/apiContainerized/.env) file.

#### Deployment
All the components are declared in a single [docker-compose.yml](/apiContainerized/docker-compose.yml). The following command needs to be executed: 

```
docker-compose up --build -d
```

The following services are deployed:

1. `mongoDB`
2. `api`
3. `mongoInit`
4. `kafka`

in the same docker network namely `company-network`. The api along with the mongoInit service wait until the mongoDB service is healthy before being deployed.

#### Testing 
The tests are provided via the [endoints_test.go](./endpoints_test.go). The test command needs to be executed after the components deployment. The command is the following:

```
go test -v
```

from the root directory.










