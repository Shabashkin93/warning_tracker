include .env
export
export COMPOSE_PROJECT_NAME := ${PROJECT_NAME}

serviceVersion :=0.0.1

repository :=develop.nagtech.ru/akazancev/warning_tracker/
dateFlag :=-X '$(repository)pkg/buildinfo.buildTime=$(shell date +%Y.%m.%d.%H%M%S)'
commitFlags :=-X '$(repository)pkg/buildinfo.commitHash=$(shell git rev-parse --short HEAD)'

psqlBackupName :='${PROJECT_NAME}-$(shell hostname)-$(shell date +%Y.%m.%d.%H%M%S)'

POSTGRES_NAME :='pg-${PROJECT_NAME}'
postgresIP :='$(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.Gateway}}{{end}}' ${POSTGRES_NAME})'
POSTGRES_HOST :=${postgresIP}
INTERNAL_POSTGRES_PORT :=5432

HIDE_OUT := '>/dev/null 2>&1'

RELEASE_NAME :='warning_tracker-$(shell date +%Y.%m.%d.%H%M%S)-$(shell git rev-parse --short HEAD)'

export POSTGRESQL_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_ADDRESS}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=${POSTGRES_SSL_MODE}

INTERNAL_CACHE_PORT :=6379
CACHE_NAME :='redis-${PROJECT_NAME}'

all: image

up:
	@docker-compose up -d --build

down:
	@docker-compose down

down-v:
	@docker-compose down -v

stop:
	@docker-compose stop

restart:
	stop
	up

clean:
	@docker-compose down -v
	@docker-compose rm -f -v

# DATABASE
db-run:
	@docker volume create ${POSTGRES_NAME}
	@docker run --rm -d --env-file .env \
		-v ${POSTGRES_NAME}:/var/lib/postgresql/data \
		--name=${POSTGRES_NAME} \
		-p ${POSTGRES_PORT}:${INTERNAL_POSTGRES_PORT} \
		postgres:latest

db-stop:
	@docker stop ${POSTGRES_NAME}

db-ip:
	@echo ${POSTGRES_HOST}

db-logs:
	@docker logs ${POSTGRES_NAME}

db-backup:
	@docker exec -i ${POSTGRES_NAME} /usr/bin/pg_dumpall -U ${POSTGRES_USER} | gzip -9 > ${PWD}/${psqlBackupName}.sql.gz
	@echo ${psqlBackupName}.sql.gz

db-restore:
	@gunzip -c ${PROJECT_NAME}*.sql.gz | docker exec -i ${POSTGRES_NAME} psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} ${HIDE_OUT}

migrate-up:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate-down:
	@migrate -database ${POSTGRESQL_URL} -path db/migrations down

migrate-install:
	git clone https://github.com/golang-migrate/migrate.git
	cd migrate
	make
	sudo cp ./migrate /usr/local/bin
	cd ..
	rm -rf migrate

db-enable-backups:
	@rm -f cron_bkp
	@if [ ! $(shell crontab -l >/dev/null 2>&1) ];then \
		touch cron_bkp; \
	else \
		crontab -l > cron_bkp; \
	fi
	@echo "0 0 * * * $(shell pwd)/backupDB.sh ${POSTGRES_NAME} ${POSTGRES_USER} ${PROJECT_NAME}" >> cron_bkp
	@crontab cron_bkp
	@rm -f cron_bkp

# \DATABASE

# backend
image: backend-delete
	@docker build -t warning_tracker:${serviceVersion} -f Dockerfile .
	@docker save warning_tracker:${serviceVersion} > warning_tracker-${serviceVersion}.tar
	@docker rmi warning_tracker:${serviceVersion}

backend-load:
	@docker load --input warning_tracker-*.tar

backend-run:
	@docker run --rm -d \
	--env POSTGRES_HOST=${POSTGRES_HOST} \
	--env INTERNAL_POSTGRES_PORT=${INTERNAL_POSTGRES_PORT} \
	--env-file .env \
	-v $(shell pwd)/certs:/etc/ca-certificates/warning_tracker \
	--name=${PROJECT_NAME}-${serviceVersion} \
	-p ${SERVER_PORT}:${SERVER_PORT} \
	warning_tracker:${serviceVersion}

backend-stop:
	@docker stop ${PROJECT_NAME}-${serviceVersion}

backend-logs:
	@docker logs ${PROJECT_NAME}-${serviceVersion}

backend-delete:
	@if [ "$(shell docker images -q warning_tracker:${serviceVersion} 2> /dev/null)" != "" ]; then \
		docker rmi warning_tracker:${serviceVersion}; \
		rm -f warning_tracker-*.tar;\
	fi
# \backend

devrun:
	image
	load
	run

release: image
	@mkdir -p ${RELEASE_NAME}/db/migrations
	@mkdir -p ${RELEASE_NAME}/certs
	@mv warning_tracker-${serviceVersion}.tar ${RELEASE_NAME}
	@cp .env.example ${RELEASE_NAME}/.env
	@cp db/migrations/* ${RELEASE_NAME}/db/migrations
	@cp MakefileRelease ${RELEASE_NAME}/Makefile
	@cp backupDB.sh ${RELEASE_NAME}/backupDB.sh
	@chmod +x ${RELEASE_NAME}/backupDB.sh
	@cp README ${RELEASE_NAME}/README
	@tar -zcvf "warning_tracker-$(shell date +%Y.%m.%d.%H%M%S)-$(shell git rev-parse --short HEAD).tar.gz" ${RELEASE_NAME}
	@rm -rf ${RELEASE_NAME}

cache-up:
	@docker volume create ${CACHE_NAME}
	@docker run --rm -d --env-file .env \
		-v ${CACHE_NAME}:/data \
		--name=${CACHE_NAME} \
		-p ${REDIS_PORT}:${INTERNAL_CACHE_PORT} \
		redis:alpine redis-server \
		--save ${REDIS_BACKUP_TIME} 1 \
		--loglevel warning \
		--requirepass ${REDIS_PASSWORD}

cache-down:
	@docker stop ${CACHE_NAME}

update-go-packages:
	@go mod tidy
	@go get -u ./...

generate:
	go generate ./internal/...
