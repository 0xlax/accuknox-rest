## Accuknox REST-API for Note-taking Application


Entering web service container 
`docker-copose run --service-ports web bash`


Building the docker container
`docker-compose build`

Run the service
`docker-compose up`


Database :Postgres
ORM Library : GORM postgres driver from gorm.io/driver/postgres


Prepare Database:
`sudo apt install postgresql`

Check status : `systemctl status postgresql`
`sudo -u postgres psql`
create database : `CREATE DATABASE notetakingapp;`



Challenges Faced:
1) Connecting Container's localhost instead of (postgres container)[https://stackoverflow.com/a/63977978/19775731]
2) Connecting POstgres Container to another container 
3) Enviroment Variables Clashed



Incase of ERROR_SQLSTATE 22023, wipe and run docker again: 
docker-compose down && docker volume rm accuknox-rest_postgres-db && docker-compose up
