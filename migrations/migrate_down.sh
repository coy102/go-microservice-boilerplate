#!/bin/bash
log=migrated.log
db=`cat dbparam.txt`
date >> $log
migrate -path source -database ${db} -verbose down $1 |& tee -a $log
echo '' >> $log
