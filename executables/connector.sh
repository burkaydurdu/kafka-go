#!/bin/bash

# Create Source Connector
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d "@connectors/source-connector.json"

# Create Sink Connector
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/ -d "@connectors/sink-connector.json"

# Update Source Connector
curl -i -X PUT -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/source-connector/config -d "@connectors/source-connector-edit.json"

# Update Source Connector
curl -i -X PUT -H "Accept:application/json" -H "Content-Type:application/json" localhost:8083/connectors/sink-connector/config -d "@connectors/sink-connector-edit.json"