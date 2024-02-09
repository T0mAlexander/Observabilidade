TIMES=1
for i in $(eval echo "{1..$TIMES}")
do
  siege -c 1 -r 10 http://localhost:8080/
  siege -c 3 -r 5 http://localhost:8080/io_task
  siege -c 2 -r 5 http://localhost:8080/cpu_task
  siege -c 5 -r 3 http://localhost:8080/random_sleep
  siege -c 2 -r 10 http://localhost:8080/random_status
  siege -c 2 -r 3 http://localhost:8080/chain
  siege -c 1 -r 1 http://localhost:8080/error_test
  sleep 3

  siege -c 1 -r 10 http://localhost:8081/
  siege -c 3 -r 5 http://localhost:8081/io_task
  siege -c 2 -r 5 http://localhost:8081/cpu_task
  siege -c 5 -r 3 http://localhost:8081/random_sleep
  siege -c 2 -r 10 http://localhost:8081/random_status
  siege -c 2 -r 3 http://localhost:8081/chain
  siege -c 1 -r 1 http://localhost:8081/error_test
  sleep 3

  siege -c 1 -r 10 http://localhost:8082/
  siege -c 3 -r 5 http://localhost:8082/io_task
  siege -c 2 -r 5 http://localhost:8082/cpu_task
  siege -c 5 -r 3 http://localhost:8082/random_sleep
  siege -c 2 -r 10 http://localhost:8082/random_status
  siege -c 2 -r 3 http://localhost:8082/chain
  siege -c 1 -r 1 http://localhost:8082/error_test
  sleep 3
done
