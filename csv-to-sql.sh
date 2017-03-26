#!/bin/bash

INPUT=geoData.csv
DATABASE=locations
TABLE=points
OUTPUT=sql/2-geoData.sql

echo "USE $DATABASE;" > $OUTPUT
cat $INPUT | sed 1d | while read line
  do
    id=`echo $line | cut -d ',' -f 1`
    lat=`echo $line | cut -d ',' -f 2`
    lng=`echo $line | cut -d ',' -f 3`

    echo "INSERT INTO $TABLE (ID, Latitude, Longitude) VALUES($id, $lat, $lng);" >> $OUTPUT
  done
