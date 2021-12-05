curl  -v -X POST http://localhost:8091/node/controller/setupServices -d 'services=kv%2Cn1ql%2Cindex'

curl  -v -X POST http://localhost:8091/pools/default -d 'memoryQuota=1256' -d 'indexMemoryQuota=1256'

curl  -u Administrator:password -v -X POST http://localhost:8091/settings/web -d 'password=password&username=Administrator&port=SAME'

curl -v -X POST http://localhost:8091/pools/default/buckets \
-u Administrator:password \
-d name=rabbit \
-d bucketType=couchbase \
-d ramQuotaMB=1024 \
-d storageMode=plasma