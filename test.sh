#!/bin/bash

BASE_URL="http://localhost:8080"
GREEN='\033[0;32m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m'

separator() {
    echo -e "${CYAN}─────────────────────────────────────────${NC}"
}

header() {
    echo -e "\n${YELLOW}$1${NC}"
    separator
}

request() {
    local description=$1
    local response=$2
    echo -e "${GREEN}> $description${NC}"
    echo -e "  $response\n"
}

header "STACK (LIFO)"
request "Push 'hola'" "$(curl -s -X POST $BASE_URL/stack/mystack/push -H "Content-Type: application/json" -d '{"value": "hola"}')"
request "Push 'mundo'" "$(curl -s -X POST $BASE_URL/stack/mystack/push -H "Content-Type: application/json" -d '{"value": "mundo"}')"
request "Estado actual" "$(curl -s $BASE_URL/stack/mystack)"
request "Pop" "$(curl -s -X POST $BASE_URL/stack/mystack/pop)"
request "Estado despues del pop" "$(curl -s $BASE_URL/stack/mystack)"

header "QUEUE (FIFO)"
request "Enqueue 'primero'" "$(curl -s -X POST $BASE_URL/queue/myqueue/enqueue -H "Content-Type: application/json" -d '{"value": "primero"}')"
request "Enqueue 'segundo'" "$(curl -s -X POST $BASE_URL/queue/myqueue/enqueue -H "Content-Type: application/json" -d '{"value": "segundo"}')"
request "Estado actual" "$(curl -s $BASE_URL/queue/myqueue)"
request "Dequeue" "$(curl -s -X POST $BASE_URL/queue/myqueue/dequeue)"
request "Estado despues del dequeue" "$(curl -s $BASE_URL/queue/myqueue)"

header "LINKED LIST"
request "Append 'nodo1'" "$(curl -s -X POST $BASE_URL/list/mylist/append -H "Content-Type: application/json" -d '{"value": "nodo1"}')"
request "Append 'nodo2'" "$(curl -s -X POST $BASE_URL/list/mylist/append -H "Content-Type: application/json" -d '{"value": "nodo2"}')"
request "Append 'nodo3'" "$(curl -s -X POST $BASE_URL/list/mylist/append -H "Content-Type: application/json" -d '{"value": "nodo3"}')"
request "Estado actual" "$(curl -s $BASE_URL/list/mylist)"
request "Remove" "$(curl -s -X POST $BASE_URL/list/mylist/remove)"
request "Estado despues del remove" "$(curl -s $BASE_URL/list/mylist)"

separator
echo -e "\n${GREEN}Tests completados${NC}\n"
