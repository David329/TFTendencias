#!/bin/bash

#xterm turn on mongo
# xterm -e sudo service mongod start &

#xterm APIRest1
"C:/Program Files/Git/bin/bash.exe" -e ./APIRest1/util/RunAPIRest.sh & 

#xterm APIRest2
"C:/Program Files/Git/bin/bash.exe" -e ./APIRest2/util/RunAPIRest.sh &

#xterm RunAPIRestBalancer
"C:/Program Files/Git/bin/bash.exe" -e ./LoadBalancer/APIRestBalancer/RunAPIRestBalancer.sh #& 

#xterm RunWebServer1
# xterm -e ./WebServer1/util/RunWebServer.sh & 

#xterm RunWebServer2
# xterm -e ./WebServer2/util/RunWebServer.sh & 

#xterm RunWebServerBalancer
# xterm -e ./LoadBalancer/WebServerBalancer/RunWebServerBalancer.sh 