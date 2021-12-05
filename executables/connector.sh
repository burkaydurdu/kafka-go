#!/bin/bash

# Create Connector
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d "@connectors/source-connector.json"

# Update Connector
curl -i -X PUT -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/source-connector/config -d "@connectors/source-connector-edit.json"
