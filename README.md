# sat-cli

## To Run in local clone the package and run a sql server 

## Create a DB sat_results_db AND Create a table in this db sat_results

```
name varchar(255) PK 
address varchar(255
city varchar(255)
country varchar(255)
pincode varchar(10) 
sat_score int(11) 
result tinyint(1)
```

## Following commands are supported 

```
  go run main.go help to get the menu

  go run main.go [command]

  Available Commands:
  delete      Delete a record by name
  getrank     Get the rank for a specific candidate
  help        Help about any command
  insert      Create a sat entry
  updatescore Update SAT score for a candidate
  view        View all SAT result data

```
