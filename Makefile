docker:
	sh ./executables/docker.sh
couchbase:
	sh ./executables/couchbase.sh
connector:
	sh ./executables/connector.sh
open_browser:
	sh ./executables/open.sh
run:
	CB_USERNAME=Administrator CB_PASSWORD=password CB_BUCKET=rabbit CB_HOST=localhost KF_TOPIC=dbserver1.inventory.customers go run .