#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh

#Levantamos mongo
sudo service mongod start

#Regresamos a la raiz para crear la carpeta bin
cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el APIRest2
cd ./APIRest2/src
go build -o "../../bin/APIRest2" ./APIRest.go

#corremos el APIRest2
read -p "Quieres levantar el APIRest2?(+si)" par
if test "$par" = "si"
then
     ./../../bin/APIRest2
else
     echo "Compilado en Bin APIRest2, pero no ejecutado"	
fi

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
