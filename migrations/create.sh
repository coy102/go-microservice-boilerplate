#!/bin/bash
log=created.log
date >> $log
migrate create -ext sql -dir source -format "unix" $1 |& tee -a $log
echo '' >> $log
