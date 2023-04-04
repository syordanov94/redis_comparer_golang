env-up:
	docker-compose up -d

test:
	docker exec rediscomparer-go-app go test ./... -v -count=1

bench-redis-sequence:
	docker exec rediscomparer-go-app go test ./... -v -bench='(Redigo|Goredis)' -benchmem -benchtime=10000x -count=10 | tee bench-redis-1000000x-sequence.txt

bench: bench-redis-sequence
	benchstat bench-redis-1000000x-sequence.txt

env-down:
	docker-compose down --remove-orphans -v
