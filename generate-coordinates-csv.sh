#!/bin/bash

AMOUNT=$1
OUTPUT=geoBig2.csv

if [ "$AMOUNT" == "" ]; then
  AMOUNT=1000000
fi

# randomly select either negative or positive value
function neg {
  rand=`shuf -i 0-1 -n 1`
  if [ $rand -eq 0 ]; then
    echo -
  else
    echo
  fi
}

echo '"id","lat","lng"' > $OUTPUT

for i in `seq 1 $AMOUNT`
do
  id=$i
  lat=`neg``shuf -i 0-90 -n 1`.`shuf -i 0-999999 -n 1`
  lng=`neg``shuf -i 0-180 -n 1`.`shuf -i 0-999999 -n 1`
  echo $id,$lat,$lng >> $OUTPUT
done
