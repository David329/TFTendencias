#!/bin/bash

#Primero se para el servicio mongod -> service mongod stop

#Creamos la estructura de carpetas para las bases de datos
mkdir -p rs1 rs2 rs3

#Iniciamos una BD en el puerto 27017
sudo mongod --dbpath /root/Escritorio/dev/upc/Tendencias/TF/Mongo/rs1 --replSet lushfly --port 27017

#Iniciamos una BD en el puerto 27018
sudo mongod --dbpath /root/Escritorio/dev/upc/Tendencias/TF/Mongo/rs2 --replSet lushfly --port 27018

#Iniciamos una BD en el puerto 27019
sudo mongod --dbpath /root/Escritorio/dev/upc/Tendencias/TF/Mongo/rs3 --replSet lushfly --port 27019


##Para ver los procesos de las BDs: ps -ef | grep mongod
##Para Matar un Proceso, en este caso la BD primary para q un esclavo tome su lugar(2do parametro): kill (NumOfProcess) 