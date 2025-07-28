# GROW LOG
Grow Log is a plant care tracker for the command line written in Go and PostgreSQL.  

## DEPENDENCIES
In order to use Grow Log, you'll need to have PostgreSQL and Go installed on your machine.  
For information on how to install PostgreSQL, visit: [PostgreSQL](https://www.postgresql.org/download/)  
For information on how to install Go, visit: [Go](https://go.dev/doc/install)  

## BUILDING AND INSTALLING GROW LOG
Assuming that Git and Go are installed; cloning, running, building, and installing the code can be done using the following commands:  
    `$ git clone https://github.com/haley-marie/growlog.git
     $ cd growlog
     $ go build
     $ go install`

Grow Log expects that you have a `postgres` instance running on the provided port with a `growlog` database created. If issues are encountered while installing Grow Log, please ensure all dependencies are satisfied and that the `growlog` database exists.  
Additionally, you will need to create a .env file in the root directory that specifies a DB_URL value. The format is as follows:  
    `DB_URL="postgres://username:password@localhost:port/growlog?sslmode=disable"`  

Where username and password are the values you assigned when setting up PostgreSQL. For more information on setting up PostgreSQL, visit: [Microsoft](https://learn.microsoft.com/en-us/windows/wsl/tutorials/wsl-database#install-postgresql)  

## USING GROW LOG
Once grow log is installed on your machine, you can run the application with `growlog` combined with one of the following commands:  
- `addPlant <plantType> [timePlanted]`: adds the specified plant to grow log.  
- `resetPlants`: removes all plants from grow log.  
- `removePlant <plantID>`: remove the specified plant from grow log.  
- `addEvent <eventName>`: creates the specified custom care event in grow log. Note that grow log comes with 5 preexisting care events, see listEvents for details on how to view care events.  
- `addLog <plantID> <eventID> [timeOfCare]`: creates the specified care log entry in grow log.  
- `listPlants`: lists all plants stored in grow log.  
- `listEvents`: lists all events stored in grow log. Grow log comes with 5 preexisiting care events, please use this command to get their IDs if needed.  
