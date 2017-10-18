#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#Levantamos mongo
#sudo service mongod start

#Regresamos a la raiz para crear la carpeta bin
#cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el APIRest1
cd ./APIRest1/src
go build -o "../../bin/APIRest1" ./APIRest.go

#corremos el APIRest1
./../../bin/APIRest1

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
