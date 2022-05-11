# NAPP Integrations
NAPP integrations is an application wich has lots of microservices to menage the NAPP e-commerce ecosystem

# Setup project

1. To set up project you will nedd have installed docker in your local machine;
2. Make the clone of the project to your local machine and than open on the napp directory;
3. Open a new terminal e than run the following command: make start
4. To stop the docker containers you can run: make stop
5. To run testes you can run: make test

 - The expose por available on docker is 8080
 - So when docker starts your applocation you can acess by http://localhost:8080/api/v1/
 - The Pré route is aways /api/v1/
 - The health check route is  http://localhost:8080/api/v1/health

# Documentation

To have a good using of this project you will need follow this two documentations bellow:

1. API Documentation
    Swagger Link: https://github.com/alvessergio/napp/blob/main/documentation/napp-api-documentation-swagger.yaml
    To see and edit you can open on Editor: https://editor.swagger.io/
    
2. Postman Collectoin 
    API Calls Link: https://github.com/alvessergio/napp/blob/main/documentation/NAPP.postman_collection.json
    To see and edit you can open on Editor: https://www.postman.com/

# Used Technologies in this project
1. Golang 1.17
2. Postgress database
3. Gorm ORM
4. Logrus
5. Docker
6. Makefile
7. Testfy

# Observations
The project is an initial step so some tests are not implemented yet
