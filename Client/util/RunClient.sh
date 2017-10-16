#!/bin/bash

## para agregar permisos
# Agregar permisos al script chmod +x ./run.sh


#Regresamos a la raiz para crear la carpeta bin
cd ./../../

#creamos la estructura

if [ ! -d bin ]; then
  echo mkdir bin
fi

#compilamos el cliente
cd ./Client/src
go build -o "../../bin/Client" ./Client.go

#corremos el cliente
read -p "Quieres levantar el Cliente?(+si)" par
if test "$par" = "si"
then
     ./../../bin/Client
else
     echo "Compilado en Bin Client, pero no ejecutado"	
fi

##NO CORRER DESDE LA MISMA RUTA LOS BINARIOS, X EL ERROR MAGICO!
