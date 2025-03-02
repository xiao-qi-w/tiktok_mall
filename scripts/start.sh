#!/bin/bash

make run svc=cart &
make run svc=checkout &
make run svc=email &
make run svc=frontend &
make run svc=order &
make run svc=payment &
make run svc=product &
make run svc=user &
wait
