#!/bin/bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  xdg-open http://localhost:8080
  xdg-open http://localhost:8081
  xdg-open http://localhost:8091
elif [[ "$OSTYPE" == "darwin"* ]]; then
  open http://localhost:8080
  open http://localhost:8081
  open http://localhost:8091
elif [[ "$OSTYPE" == "cygwin" ]]; then
  start http://localhost:8080
  start http://localhost:8081
  start http://localhost:8091
fi