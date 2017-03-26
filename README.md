# gps-sort

A coordinates sorting tool.

### Execution Flags

|Flag|Default value|Description|
|:----|:---|:---|
|input-file|geoData.csv|Input CSV File|
|mysql-host|127.0.0.1|MySQL host|
|mysql-port|3306|MySQL port|
|mysql-user|root|MySQL user|
|mysql-pass|my-secret-pw|MySQL pass|
|mysql-database|locations|MySQL database|
|input-mode|file|Data location (file or db)|
|top-amount|5|A number of records to show in TOPs|
|comparison-point-lat|51.925146|The latitude of the point the distance must be calculated to|
|comparison-point-lng|4.478617|The longitude of the point the distance must be calculated to|

### Godep required

The application uses MySQL driver. Godep allows to use the one provided in the same repository.

```
go get github.com/tools/godep
```

### Reading data from file

A quick example using default flags.

1. Navigate to the project's directory and make sure the *geoData.csv* file is there.
2. Run `godep go run main.go -input-mode file`
3. An output similar to the following will appear:
```
reading from the local file geoData.csv

calculating relative distances of 100 records to 51.925146,4.478617...
time taken: 25 microseconds

sorting records by distance...
time taken: 11 microseconds

== TOP 5 closest coordinates to 51.925146,4.478617 ==
1 --> ID: 442406, Latitude: 51.927167, Longitude: 4.482217, Distance: 333 meters
2 --> ID: 285782, Latitude: 51.925356, Longitude: 4.486310, Distance: 528 meters
3 --> ID: 429151, Latitude: 51.925630, Longitude: 4.488034, Distance: 648 meters
4 --> ID: 512818, Latitude: 51.926815, Longitude: 4.489072, Distance: 740 meters
5 --> ID: 25182, Latitude: 51.924912, Longitude: 4.490593, Distance: 821 meters

== TOP 5 furthest coordinates to 51.925146,4.478617 ==
1 --> ID: 7818, Latitude: 37.866754, Longitude: -122.259099, Distance: 8776646 meters
2 --> ID: 382013, Latitude: 37.399450, Longitude: -5.971514, Distance: 1810117 meters
3 --> ID: 381823, Latitude: 37.168004, Longitude: -3.602987, Distance: 1758848 meters
4 --> ID: 382582, Latitude: 37.176867, Longitude: -3.608897, Distance: 1758080 meters
5 --> ID: 382693, Latitude: 40.970240, Longitude: -5.661052, Distance: 1441719 meters
```

### Reading data from MySQL database

A quick example using default flags.

1. Navigate to the project's directory.
2. Run the MySQL instance executing the following: `docker run --name gps-sort-mysql -e MYSQL_DATABASE=locations -e MYSQL_ROOT_PASSWORD=my-secret-pw -v $PWD/sql:/docker-entrypoint-initdb.d -p 3306:3306 -d mysql:5.7`
3. Run `godep go run main.go -input-mode db`
4. An output similar to the following will appear:
```
reading from the database root:my-secret-pw@tcp(127.0.0.1:3306)/locations

calculating relative distances of 100 records to 51.925146,4.478617...
time taken: 35 microseconds

sorting records by distance...
time taken: 14 microseconds

== TOP 5 closest coordinates to 51.925146,4.478617 ==
1 --> ID: 442406, Latitude: 51.927167, Longitude: 4.482217, Distance: 333 meters
2 --> ID: 285782, Latitude: 51.925356, Longitude: 4.486310, Distance: 528 meters
3 --> ID: 429151, Latitude: 51.925630, Longitude: 4.488034, Distance: 648 meters
4 --> ID: 512818, Latitude: 51.926815, Longitude: 4.489072, Distance: 740 meters
5 --> ID: 25182, Latitude: 51.924912, Longitude: 4.490593, Distance: 821 meters

== TOP 5 furthest coordinates to 51.925146,4.478617 ==
1 --> ID: 7818, Latitude: 37.866754, Longitude: -122.259099, Distance: 8776646 meters
2 --> ID: 382013, Latitude: 37.399450, Longitude: -5.971514, Distance: 1810117 meters
3 --> ID: 381823, Latitude: 37.168004, Longitude: -3.602987, Distance: 1758848 meters
4 --> ID: 382582, Latitude: 37.176867, Longitude: -3.608897, Distance: 1758080 meters
5 --> ID: 382693, Latitude: 40.970240, Longitude: -5.661052, Distance: 1441719 meters
```
