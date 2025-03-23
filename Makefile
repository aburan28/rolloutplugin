build: 
	go build -o bin/statefulset .
	aws s3 cp bin/statefulset s3://helmchartsglobal/statefulset 
clean:
	rm -rf bin/ 