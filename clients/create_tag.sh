curl -i -H "content-type=application/json" http://localhost:8080/tags --data-binary @- << EOF
{
    "name": "boys"
}
EOF