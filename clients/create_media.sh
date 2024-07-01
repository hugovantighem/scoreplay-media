curl -i -H "content-type=multipart/form-data" http://localhost:8080/media \
    -F name=world \
    -F tags=5d8a8dd6-407c-4c1c-b9a8-d759d70102ae \
    -F file=@myfile.txt