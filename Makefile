build: 
	GODEBUG=asyncpreemptoff=1 go build -o bin/statefulset .
	chmod 0777 bin/statefulset
	aws s3 cp bin/statefulset s3://helmchartsglobal/statefulset 
clean:
	rm -rf bin/